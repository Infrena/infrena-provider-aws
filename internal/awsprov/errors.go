package awsprov

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws/retry"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/infrata/infrata/pkg/provider"
)

// classify answers infrata's question — how dangerous is another attempt — from the error alone. It
// must be a pure function: the plugin SDK asks any one configured instance, not the one that failed.
//
// Order matters. A throttle may carry a 5xx status, so the throttle code is checked before the status.
// A *url.Error is a net.Error, so the dial case is checked before the generic one. EC2's query-protocol
// errors carry no smithy fault, so the HTTP status is the only server-fault signal.
func classify(err error) provider.Retryability {
	if err == nil || errors.Is(err, context.Canceled) {
		return provider.NotSafeToRetry
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		if _, throttled := retry.DefaultThrottleErrorCodes[apiErr.ErrorCode()]; throttled {
			return provider.SafeToRetry // refused before acting
		}
	}
	var opErr *net.OpError
	if errors.As(err, &opErr) && opErr.Op == "dial" {
		return provider.SafeToRetry // no connection, so nothing was sent
	}
	var respErr *smithyhttp.ResponseError
	if errors.As(err, &respErr) && respErr.HTTPStatusCode() >= 500 {
		return provider.ConditionallyRetryable // may have acted: no client token tells us otherwise
	}
	var netErr net.Error
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || errors.As(err, &netErr) {
		return provider.ConditionallyRetryable // sent, answer lost
	}
	return provider.NotSafeToRetry
}

// hasCode reports whether err carries one of the given AWS error codes.
func hasCode(err error, codes ...string) bool {
	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		return false
	}
	for _, c := range codes {
		if apiErr.ErrorCode() == c {
			return true
		}
	}
	return false
}

// apiFailure is the message a user reads in a failed apply, with the SDK error kept underneath for
// classification.
type apiFailure struct {
	msg string
	err error
}

func (e *apiFailure) Error() string { return e.msg }
func (e *apiFailure) Unwrap() error { return e.err }

// failed says which instance, which call, where, what AWS said, and the request ID AWS support will ask
// for. Never anything from the request itself, which may hold a secret.
func (p *Provider) failed(op, region, awsID string, err error) error {
	where := region
	if awsID != "" {
		where = formatID(region, awsID)
	}
	msg := fmt.Sprintf("aws instance %q: EC2 %s for %s failed: ", p.instance, op, where)
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		msg += apiErr.ErrorCode() + ": " + apiErr.ErrorMessage()
	} else {
		msg += err.Error()
	}
	var re *awshttp.ResponseError
	if errors.As(err, &re) && re.ServiceRequestID() != "" {
		msg += " (request ID " + re.ServiceRequestID() + ")"
	}
	switch {
	case hasCode(err, "UnauthorizedOperation", "AuthFailure"):
		msg += fmt.Sprintf("\ncheck the IAM permissions of the credentials instance %q uses", p.instance)
	case hasCode(err, "DependencyViolation"):
		msg += "\nsomething infrata does not manage still depends on it (subnets, network interfaces, gateways): remove that first"
	}
	return &apiFailure{msg: msg, err: err}
}
