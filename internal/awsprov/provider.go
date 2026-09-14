package awsprov

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/schema"
)

// Provider is one configured instance: one AWS account's credentials, any number of regions.
type Provider struct {
	instance string
	config   instanceConfig
	clients  *clients
	patience patience
}

var _ provider.Provider = (*Provider)(nil)

func newProvider(instance string, ic instanceConfig, cfg aws.Config) *Provider {
	return &Provider{instance: instance, config: ic, clients: newClients(cfg), patience: notFoundPatience}
}

func (p *Provider) Name() string                              { return PluginName }
func (p *Provider) Definitions() []*schema.ResourceDefinition { return definitions() }

func unknownType(t string) error { return fmt.Errorf("the aws plugin does not serve %q", t) }

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	switch current.Type {
	case typeVPC:
		return p.readVPC(ctx, current.ProviderID, p.patience)
	case typeSubnet:
		return p.readSubnet(ctx, current.ProviderID, p.patience)
	}
	return nil, unknownType(current.Type)
}

func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	switch desired.Type {
	case typeVPC:
		return p.createVPC(ctx, desired.Attrs)
	case typeSubnet:
		return p.createSubnet(ctx, desired.Attrs)
	}
	return nil, unknownType(desired.Type)
}

func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	switch current.Type {
	case typeVPC:
		return p.updateVPC(ctx, current, desired)
	case typeSubnet:
		return p.updateSubnet(ctx, current, desired)
	}
	return nil, unknownType(current.Type)
}

func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	switch current.Type {
	case typeVPC:
		return p.deleteVPC(ctx, current)
	case typeSubnet:
		return p.deleteSubnet(ctx, current)
	}
	return unknownType(current.Type)
}

// Discover and Import arrive with discovery; until then each says so rather than guessing.

func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

// ClassifyError delegates to classify, a pure function of the error.
func (p *Provider) ClassifyError(err error) provider.Retryability { return classify(err) }
