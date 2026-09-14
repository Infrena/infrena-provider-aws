package ccprov

import (
	"context"

	"github.com/infrata/infrata/pkg/provider"
)

// Discover is not built yet. It says so rather than guessing.

func (p *Provider) Discover(context.Context, provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	return nil, provider.ErrNotImplemented
}
