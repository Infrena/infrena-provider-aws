package awsprov

import (
	"context"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/schema"
)

// unconfigured stands in for the Cloud Control provider until Task 9 of the plan builds it: a configured instance that
// says so rather than guessing.
type unconfigured struct{ cat *catalog.Catalog }

var _ provider.Provider = unconfigured{}

func (u unconfigured) Name() string                              { return PluginName }
func (u unconfigured) Definitions() []*schema.ResourceDefinition { return u.cat.Definitions() }
func (u unconfigured) Read(context.Context, *resource.ResourceState) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Create(context.Context, *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Update(context.Context, *resource.ResourceState, *resource.DesiredResource) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Delete(context.Context, *resource.ResourceState) error {
	return provider.ErrNotImplemented
}
func (u unconfigured) Discover(context.Context, provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) Import(context.Context, string, string) (*resource.ResourceState, error) {
	return nil, provider.ErrNotImplemented
}
func (u unconfigured) ClassifyError(error) provider.Retryability { return provider.NotSafeToRetry }
