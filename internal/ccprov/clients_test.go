package ccprov

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
)

// ADAPTIVE, NOT STANDARD, and the difference is the whole reason this test
// exists. The standard retryer backs off one request at a time, which does
// nothing to lower the rate a FAN-OUT hits the API: a refresh runs eight
// reads at once, each retrying on its own schedule, and on 2026-09-18 that
// was enough to keep a real account throttled through the engine's own five
// attempts. Adaptive mode adds a client-wide token bucket that measures
// throttle responses and paces every request the client makes.
//
// Asserted on the constructed client rather than on behaviour: what would
// regress here is somebody dropping the option, and the timing tests that
// would "prove" pacing are the kind this project has already been burned by
// twice.
func TestCloudControlClientsUseAdaptiveRetry(t *testing.T) {
	c := newClients(aws.Config{Region: "us-east-1"})

	got := c.get("us-east-1").Options().Retryer
	if _, ok := got.(*retry.AdaptiveMode); !ok {
		t.Errorf("Cloud Control retryer is %T, want *retry.AdaptiveMode — a per-request backoff does not slow a fan-out", got)
	}
}

// EC2 is asked one read-only question (defaults.go), but it is asked once
// per VPC or subnet during a discovery, so it fans out the same way.
func TestEC2ClientsUseAdaptiveRetry(t *testing.T) {
	c := newClients(aws.Config{Region: "us-east-1"})

	got := c.ec2("us-east-1").Options().Retryer
	if _, ok := got.(*retry.AdaptiveMode); !ok {
		t.Errorf("EC2 retryer is %T, want *retry.AdaptiveMode", got)
	}
}
