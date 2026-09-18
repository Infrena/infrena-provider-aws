package ccprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// clients holds one Cloud Control client per region, built on first use. Retries stay on: every mutation carries a
// ClientToken, so a resend cannot make a second resource.
//
// THE RETRYER IS ADAPTIVE, NOT STANDARD, and the caching above is what makes that mean anything. The standard
// retryer backs off one request at a time, which does not lower the rate a FAN-OUT puts on the API: a refresh reads
// eight resources at once, each retrying on its own schedule, and on 2026-09-18 that kept a real account throttled
// straight through infrena's own five attempts — the retries were part of the load. Adaptive mode adds a token
// bucket that measures throttled responses and paces every request the client sends, so the whole client slows down
// together instead of each request discovering the throttle separately.
//
// ONE BUCKET PER REGION PER SERVICE, which falls out of these being cached: AWS quotas are per account per region,
// so a throttle in us-east-1 should not slow eu-west-1, and Cloud Control and EC2 are metered apart. A client built
// fresh per call would reset its bucket every time and learn nothing, which is the failure mode this cache happens
// to prevent.
//
// EC2 clients live here too, though nothing but discovery uses them: every resource is created, read, updated and
// deleted through Cloud Control, and EC2 is asked one read-only question Cloud Control cannot answer (defaults.go).
type clients struct {
	cfg         aws.Config
	mu          sync.Mutex
	byRegion    map[string]*cloudcontrol.Client
	ec2ByRegion map[string]*ec2.Client
}

// retryer is the retryer every client here is built with.
//
// A configuration that already names one WINS, and that is deliberate: AWS_RETRY_MODE and the shared config file
// are how an operator says what their account needs, and a provider that overrode them would be answering a
// question the user had already answered. Only the default — nil, meaning "the SDK's standard three attempts" — is
// replaced.
func (c *clients) retryer() aws.Retryer {
	if c.cfg.Retryer != nil {
		return c.cfg.Retryer()
	}
	return retry.NewAdaptiveMode()
}

func newClients(cfg aws.Config) *clients {
	return &clients{cfg: cfg, byRegion: map[string]*cloudcontrol.Client{}, ec2ByRegion: map[string]*ec2.Client{}}
}

func (c *clients) get(region string) *cloudcontrol.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.byRegion[region]; ok {
		return cl
	}
	cl := cloudcontrol.NewFromConfig(c.cfg, func(o *cloudcontrol.Options) {
		o.Region = region
		o.Retryer = c.retryer()
	})
	c.byRegion[region] = cl
	return cl
}

func (c *clients) ec2(region string) *ec2.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.ec2ByRegion[region]; ok {
		return cl
	}
	cl := ec2.NewFromConfig(c.cfg, func(o *ec2.Options) {
		o.Region = region
		o.Retryer = c.retryer()
	})
	c.ec2ByRegion[region] = cl
	return cl
}
