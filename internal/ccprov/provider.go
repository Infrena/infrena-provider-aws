package ccprov

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/schema"
	"github.com/infrena/infrena/pkg/value"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = "aws"

// Options is what an instance's configuration tells the provider beyond credentials.
type Options struct {
	DiscoverRegions []string // regions Discover scans for regional types
	DiscoverTypes   []string // what Discover lists when asked about everything; empty means the catalog's DiscoverDefault (P4)
}

// Provider is one configured instance: one account's credentials, every catalog type, any number of regions.
type Provider struct {
	instance string
	cat      *catalog.Catalog
	opts     Options
	clients  *clients
	ec2      ec2Defaults // what AWS itself owns in a region; asked by Discover only
	patience patience
	pacing   pacing
	log      io.Writer // stderr: stdout is the plugin protocol
	token    func() string

	defsOnce sync.Once
	defs     []*schema.ResourceDefinition
}

var _ provider.Provider = (*Provider)(nil)

// New builds an instance. It makes no network call.
func New(instance string, cat *catalog.Catalog, cfg aws.Config, opts Options) *Provider {
	cl := newClients(cfg)
	return &Provider{
		instance: instance, cat: cat, opts: opts, clients: cl, ec2: &ec2API{clients: cl},
		patience: notFoundPatience, pacing: defaultPacing, log: os.Stderr, token: newToken,
	}
}

func (p *Provider) Name() string { return PluginName }

func (p *Provider) Definitions() []*schema.ResourceDefinition {
	p.defsOnce.Do(func() { p.defs = p.cat.Definitions() })
	return p.defs
}

// ClassifyError delegates to classify, a pure function of the error.
func (p *Provider) ClassifyError(err error) provider.Retryability { return classify(err) }

func (p *Provider) lookup(name string) (*catalog.Type, error) {
	t, ok := p.cat.Lookup(name)
	if !ok {
		return nil, fmt.Errorf("the aws plugin does not serve %q; run `infrena explain <type>` to check a name", name)
	}
	return t, nil
}

// regionOf is where a resource lives: its region attribute, or GlobalRegion for a global type.
func regionOf(t *catalog.Type, attrs map[string]value.Value) (string, error) {
	if t.Global() {
		return GlobalRegion, nil
	}
	region, ok := attrs[t.RegionAttr].AsString()
	if !ok || region == "" {
		return "", fmt.Errorf("%s needs %s: set it on the resource, or once for every resource with the provider's defaults: {%s: ...}",
			t.Name, t.RegionAttr, t.RegionAttr)
	}
	return region, nil
}

// assumedState records what configuration asked for, when AWS created a resource the plugin cannot read back yet.
func assumedState(t *catalog.Type, region, identifier string, attrs map[string]value.Value) *resource.ResourceState {
	out := map[string]value.Value{}
	for _, a := range t.Attributes {
		if v, ok := attrs[a.Name]; ok && !(a.Computed && !a.Optional) {
			out[a.Name] = v
		}
	}
	if !t.Global() {
		out[t.RegionAttr] = value.String(region, value.SourceProvider)
	}
	return &resource.ResourceState{Type: t.Name, ProviderID: FormatID(t, region, identifier), Attributes: out}
}

// newToken is a ClientToken: Cloud Control treats a repeat within 36 hours as the same request.
func newToken() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
