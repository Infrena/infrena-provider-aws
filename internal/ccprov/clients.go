package ccprov

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// clients holds one Cloud Control client per region, built on first use. The SDK's standard retryer stays on: every
// mutation carries a ClientToken, so a resend cannot make a second resource.
type clients struct {
	cfg      aws.Config
	mu       sync.Mutex
	byRegion map[string]*cloudcontrol.Client
}

func newClients(cfg aws.Config) *clients {
	return &clients{cfg: cfg, byRegion: map[string]*cloudcontrol.Client{}}
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
