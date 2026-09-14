package ccprov

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/provider"
)

// Discover lists what exists for the requested types in the instance's regions, reading each resource once.
func (p *Provider) Discover(ctx context.Context, req provider.DiscoverRequest) ([]provider.DiscoveredResource, error) {
	var types []*catalog.Type
	var needsParent, unlistable []string
	for _, name := range p.discoverTypes(req.Types) {
		t, ok := p.cat.Lookup(name)
		switch {
		case !ok:
			continue
		case !t.HasList:
			unlistable = append(unlistable, name)
		case t.ListNeedsModel:
			needsParent = append(needsParent, name)
		default:
			types = append(types, t)
		}
	}
	if len(unlistable) > 0 {
		fmt.Fprintf(p.log, "aws instance %q: not discovering %s: AWS offers no way to list them\n", p.instance, strings.Join(unlistable, ", "))
	}
	if len(needsParent) > 0 {
		fmt.Fprintf(p.log, "aws instance %q: not discovering %s: listing them needs a parent resource, which discovery does not have\n",
			p.instance, strings.Join(needsParent, ", "))
	}
	if len(p.opts.DiscoverRegions) == 0 && slices.ContainsFunc(types, func(t *catalog.Type) bool { return !t.Global() }) {
		return nil, fmt.Errorf("aws instance %q has no `discover_regions`, so discovery does not know where to look: set it, e.g. discover_regions: [us-east-1]",
			p.instance)
	}

	var out []provider.DiscoveredResource
	var failed []error
	attempts := 0
	for _, t := range types {
		regions := p.opts.DiscoverRegions
		if t.Global() {
			regions = []string{GlobalRegion}
		}
		for _, region := range regions {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			attempts++
			found, err := p.discoverIn(ctx, t, region)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil, err
				}
				fmt.Fprintln(p.log, err)
				failed = append(failed, err)
				continue
			}
			out = append(out, found...)
		}
	}
	if attempts > 0 && len(failed) == attempts {
		return nil, errors.Join(failed...)
	}
	return out, nil
}

func (p *Provider) discoverIn(ctx context.Context, t *catalog.Type, region string) ([]provider.DiscoveredResource, error) {
	pages := cloudcontrol.NewListResourcesPaginator(p.clients.get(region), &cloudcontrol.ListResourcesInput{TypeName: aws.String(t.CFN)})
	var out []provider.DiscoveredResource
	for pages.HasMorePages() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, failure(p.instance, "ListResources", t, region, err)
		}
		for _, d := range page.ResourceDescriptions {
			st, err := p.read(ctx, t, region, aws.ToString(d.Identifier), nil, once)
			if err != nil {
				return nil, err
			}
			if st != nil { // gone between the list and the read
				out = append(out, provider.DiscoveredResource{Type: t.Name, ProviderID: st.ProviderID, Attributes: st.Attributes})
			}
		}
	}
	return out, nil
}

// discoverTypes is what a request covers (P4). A request for every type the catalog holds is infrata asking about
// everything, which here is ~1,584 ListResources calls per region: it means the instance's discover_types, or the
// catalog's default set. A request naming fewer was narrowed by the user.
func (p *Provider) discoverTypes(requested []string) []string {
	if len(requested) < len(p.cat.Types) {
		return requested
	}
	if len(p.opts.DiscoverTypes) > 0 {
		return p.opts.DiscoverTypes
	}
	return p.cat.DiscoverDefault
}
