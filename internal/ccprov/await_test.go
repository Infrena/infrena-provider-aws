package ccprov

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrena/infrena-provider-aws/internal/ccfake"
	"github.com/infrena/infrena/pkg/provider"
)

// scripted answers GetResourceRequestStatus from a list, repeating the last, and records whether each call's context
// had been cancelled.
type scripted struct {
	events  []types.ProgressEvent
	calls   int
	ctxErrs []error
}

func (s *scripted) GetResourceRequestStatus(ctx context.Context, _ *cloudcontrol.GetResourceRequestStatusInput, _ ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceRequestStatusOutput, error) {
	s.ctxErrs = append(s.ctxErrs, ctx.Err())
	ev := s.events[min(s.calls, len(s.events)-1)]
	s.calls++
	return &cloudcontrol.GetResourceRequestStatusOutput{ProgressEvent: &ev}, nil
}

// clock is a fake time source whose sleeps advance it.
type clock struct {
	now   time.Time
	slept []time.Duration
}

func (c *clock) pacing() pacing {
	return pacing{
		first: time.Second, max: 4 * time.Second,
		now: func() time.Time { return c.now },
		sleep: func(_ context.Context, d time.Duration) error {
			c.slept = append(c.slept, d)
			c.now = c.now.Add(d)
			return nil
		},
	}
}

func event(status types.OperationStatus) types.ProgressEvent {
	return types.ProgressEvent{OperationStatus: status, RequestToken: aws.String("req-1"), Operation: types.OperationCreate}
}

func TestAwaitPollsWithGrowingDelaysUntilSuccess(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress), event(types.OperationStatusInProgress), event(types.OperationStatusSuccess)}}
	start := event(types.OperationStatusPending)
	ev, err := c.pacing().await(context.Background(), api, &start, time.Hour)
	if err != nil || ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("await = %+v, %v", ev, err)
	}
	if got := c.slept; len(got) != 3 || got[0] != time.Second || got[1] != 2*time.Second || got[2] != 4*time.Second {
		t.Errorf("slept %v, want [1s 2s 4s]", got)
	}
}

func TestAwaitHonoursRetryAfter(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	start := event(types.OperationStatusInProgress)
	start.RetryAfter = aws.Time(c.now.Add(7 * time.Second))
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusSuccess)}}
	if _, err := c.pacing().await(context.Background(), api, &start, time.Hour); err != nil {
		t.Fatal(err)
	}
	if len(c.slept) != 1 || c.slept[0] != 7*time.Second {
		t.Errorf("slept %v, want [7s]", c.slept)
	}
}

// TestCancellationNeverAbandonsARequestAlreadySent. The host waits for the real answer: an abandoned create is an
// orphan.
func TestCancellationNeverAbandonsARequestAlreadySent(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress), event(types.OperationStatusSuccess)}}
	start := event(types.OperationStatusInProgress)
	if _, err := c.pacing().await(ctx, api, &start, time.Hour); err != nil {
		t.Fatal(err)
	}
	for i, e := range api.ctxErrs {
		if e != nil {
			t.Errorf("status call %d ran with a cancelled context: %v", i+1, e)
		}
	}
}

func TestAFailedRequestIsAHandlerError(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	failed := event(types.OperationStatusFailed)
	failed.ErrorCode, failed.StatusMessage, failed.Identifier = types.HandlerErrorCodeNotStabilized, aws.String("still attaching"), aws.String("vpc-9")
	api := &scripted{events: []types.ProgressEvent{failed}}
	start := event(types.OperationStatusInProgress)
	ev, err := c.pacing().await(context.Background(), api, &start, time.Hour)
	var he *HandlerError
	if !errors.As(err, &he) || he.Code != "NotStabilized" || he.Identifier != "vpc-9" || he.Token != "req-1" || he.Message != "still attaching" {
		t.Fatalf("err = %#v", err)
	}
	if ev == nil || aws.ToString(ev.Identifier) != "vpc-9" {
		t.Errorf("the final event was not returned with the error: %+v", ev)
	}
}

func TestAwaitGivesUpAtTheHandlerTimeoutAndSaysTheRequestMayFinish(t *testing.T) {
	c := &clock{now: time.Unix(1_700_000_000, 0)}
	api := &scripted{events: []types.ProgressEvent{event(types.OperationStatusInProgress)}}
	start := event(types.OperationStatusInProgress)
	_, err := c.pacing().await(context.Background(), api, &start, 10*time.Second)
	if err == nil || !strings.Contains(err.Error(), "req-1") || !strings.Contains(err.Error(), "refresh") {
		t.Fatalf("err = %v", err)
	}
	if classify(err) != provider.ConditionallyRetryable {
		t.Errorf("classify = %v, want ConditionallyRetryable", classify(err))
	}
	var total time.Duration
	for _, d := range c.slept {
		total += d
	}
	if total > 10*time.Second {
		t.Errorf("waited %s, beyond the 10s timeout", total)
	}
}

func TestTimeoutsComeFromTheSchemaOrCloudFormationsDefault(t *testing.T) {
	db := mustType(t, "aws.dbinstance")
	if got := timeoutFor(db, "create"); got != 2160*time.Minute {
		t.Errorf("create = %s", got)
	}
	if got := timeoutFor(db, "delete"); got != 120*time.Minute {
		t.Errorf("delete = %s, want the 120 minute default", got)
	}
}

// TestAwaitAgainstTheFake: the real SDK's GetResourceRequestStatus, and the per-region client cache.
func TestAwaitAgainstTheFake(t *testing.T) {
	fake := ccfake.New()
	defer fake.Close()
	fake.Register(ccfake.TypeConfig{TypeName: "AWS::EC2::VPC", Identifier: "VpcId", IDPrefix: "vpc-"})
	fake.PollsToComplete = 3
	cl := newClients(testConfig(fake.URL, 1))
	if cl.get("eu-west-1") != cl.get("eu-west-1") || cl.get("eu-west-1") == cl.get("us-east-1") {
		t.Fatal("clients are not cached per region")
	}
	if got := cl.get("eu-west-1").Options().Region; got != "eu-west-1" {
		t.Fatalf("client region = %q", got)
	}
	out, err := cl.get("eu-west-1").CreateResource(context.Background(), &cloudcontrol.CreateResourceInput{
		TypeName: aws.String("AWS::EC2::VPC"), DesiredState: aws.String(`{}`), ClientToken: aws.String("t"),
	})
	if err != nil {
		t.Fatal(err)
	}
	ev, err := instantPacing.await(context.Background(), cl.get("eu-west-1"), out.ProgressEvent, time.Minute)
	if err != nil || ev.OperationStatus != types.OperationStatusSuccess {
		t.Fatalf("await = %+v, %v", ev, err)
	}
	if n := fake.Calls("GetResourceRequestStatus"); n != 3 {
		t.Errorf("status calls = %d, want 3", n)
	}
}
