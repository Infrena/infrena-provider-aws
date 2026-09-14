package awsprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// clients holds one EC2 client per region, built on first use from the instance's configuration.
type clients struct {
	cfg      aws.Config
	mu       sync.Mutex
	byRegion map[string]*ec2.Client
}

func newClients(cfg aws.Config) *clients {
	return &clients{cfg: cfg, byRegion: map[string]*ec2.Client{}}
}

func (c *clients) ec2(region string) *ec2.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.byRegion[region]; ok {
		return cl
	}
	cl := ec2.NewFromConfig(c.cfg, func(o *ec2.Options) { o.Region = region })
	c.byRegion[region] = cl
	return cl
}

// createOnce stops the SDK resending a create. CreateVpc and CreateSubnet take no client token, so a
// resend after AWS acted makes a second resource nothing records; infrata's executor decides instead,
// from ClassifyError. The SDK ignores a per-call value equal to the client's own, which is harmless:
// equal means the client already makes one attempt.
func createOnce(o *ec2.Options) { o.RetryMaxAttempts = 1 }
