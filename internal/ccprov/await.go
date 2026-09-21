package ccprov

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

type statusAPI interface {
	GetResourceRequestStatus(context.Context, *cloudcontrol.GetResourceRequestStatusInput, ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error)
}

// pacing is how often a request's status is checked.
type pacing struct {
	first, max time.Duration
	sleep      func(context.Context, time.Duration) error
	now        func() time.Time
}

var defaultPacing = pacing{first: time.Second, max: 10 * time.Second, sleep: sleepCtx, now: time.Now}

// defaultTimeoutMinutes is CloudFormation's handler timeout when a schema names none.
const defaultTimeoutMinutes = 120

func timeoutFor(t *catalog.Type, handler string) time.Duration {
	if m := t.Timeouts[handler]; m > 0 {
		return time.Duration(m) * time.Minute
	}
	return defaultTimeoutMinutes * time.Minute
}

// waitTimeout unwraps to context.DeadlineExceeded, so it classifies as ConditionallyRetryable.
type waitTimeout struct {
	token  string
	status types.OperationStatus
	after  time.Duration
}

func (e *waitTimeout) Error() string {
	return fmt.Sprintf("request %s was still %s after %s; it may yet finish, so run `infrena refresh` before trying again",
		e.token, e.status, e.after)
}
func (e *waitTimeout) Unwrap() error { return context.DeadlineExceeded }

// await follows a request until it settles. Once a request is sent, cancellation never abandons it: the host waits for
// the real answer, because an abandoned create is a resource nothing records. The handler's timeout bounds the wait.
// The last event is returned with any error, so a caller can still see an identifier AWS assigned.
func (pc pacing) await(ctx context.Context, api statusAPI, ev *types.ProgressEvent, timeout time.Duration) (*types.ProgressEvent, error) {
	ctx = context.WithoutCancel(ctx)
	token := aws.ToString(ev.RequestToken)
	deadline := pc.now().Add(timeout)
	delay := pc.first
	for {
		switch ev.OperationStatus {
		case types.OperationStatusSuccess:
			return ev, nil
		case types.OperationStatusFailed, types.OperationStatusCancelComplete:
			code := string(ev.ErrorCode)
			if code == "" {
				code = "Cancelled"
			}
			return ev, &HandlerError{Operation: string(ev.Operation), Code: code, Message: aws.ToString(ev.StatusMessage),
				Identifier: aws.ToString(ev.Identifier), Token: token}
		}
		now := pc.now()
		if !now.Before(deadline) {
			return ev, &waitTimeout{token: token, status: ev.OperationStatus, after: timeout}
		}
		wait := delay
		if ev.RetryAfter != nil {
			if d := ev.RetryAfter.Sub(now); d > wait {
				wait = d
			}
		}
		wait = min(wait, deadline.Sub(now))
		_ = pc.sleep(ctx, wait) // ctx is never cancelled here
		delay = min(delay*2, pc.max)
		out, err := api.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: aws.String(token)})
		if err != nil {
			return ev, fmt.Errorf("checking request %s: %w", token, err)
		}
		ev = out.ProgressEvent
	}
}
