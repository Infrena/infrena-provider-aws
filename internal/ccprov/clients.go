package ccprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// clients holds one Cloud Control client per region, built on first use. The SDK's standard retryer stays on: every
// mutation carries a ClientToken, so a resend cannot make a second resource.
//
// EC2 clients live here too, though nothing but discovery uses them: every resource is created, read, updated and
// deleted through Cloud Control, and EC2 is asked one read-only question Cloud Control cannot answer (defaults.go).
type clients struct {
	cfg         aws.Config
	mu          sync.Mutex
	byRegion    map[string]*cloudcontrol.Client
	ec2ByRegion map[string]*ec2.Client
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
	cl := cloudcontrol.NewFromConfig(c.cfg, func(o *cloudcontrol.Options) { o.Region = region })
	c.byRegion[region] = cl
	return cl
}

func (c *clients) ec2(region string) *ec2.Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cl, ok := c.ec2ByRegion[region]; ok {
		return cl
	}
	cl := ec2.NewFromConfig(c.cfg, func(o *ec2.Options) { o.Region = region })
	c.ec2ByRegion[region] = cl
	return cl
}
