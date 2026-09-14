package awsprov

import (
	"context"

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

// The operations arrive type by type in later changes. Until then each says so rather than guessing.

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	return provider.ErrNotImplemented
}

func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}

func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}

// ClassifyError answers NotSafeToRetry, the safe default, until classification is built.
func (p *Provider) ClassifyError(err error) provider.Retryability { return provider.NotSafeToRetry }
