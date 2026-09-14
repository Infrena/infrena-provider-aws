package ccprov

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata/pkg/resource"
	"github.com/infrata/infrata/pkg/value"
)

// Create sends the desired state, waits for the request, and reads the resource back. Once AWS may have created
// something it never returns an error for a resource that exists: the host drops an errored create's result, and the
// resource would be real and recorded nowhere.
func (p *Provider) Create(ctx context.Context, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	t, err := p.lookup(desired.Type)
	if err != nil {
		return nil, err
	}
	region, err := regionOf(t, desired.Attrs)
	if err != nil {
		return nil, err
	}
	body, err := desiredJSON(t, desired.Attrs)
	if err != nil {
		return nil, err
	}
	doc, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	cl := p.clients.get(region)
	out, err := cl.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName: aws.String(t.CFN), DesiredState: aws.String(string(doc)), ClientToken: aws.String(p.token()),
	})
	if err != nil {
		return nil, failure(p.instance, "CreateResource", t, region, err)
	}

	ctx = context.WithoutCancel(ctx) // from here AWS may have acted
	ev, waitErr := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "create"))
	id := aws.ToString(ev.Identifier)
	if id == "" {
		id = aws.ToString(out.ProgressEvent.Identifier)
	}
	if id == "" {
		if waitErr == nil {
			waitErr = fmt.Errorf("request %s succeeded without an identifier", aws.ToString(out.ProgressEvent.RequestToken))
		}
		return nil, failure(p.instance, "create", t, region, waitErr)
	}
	where := FormatID(t, region, id)

	st, readErr := p.read(ctx, t, region, id, desired.Attrs, p.patience)
	switch {
	case readErr == nil && st != nil:
		if waitErr != nil {
			fmt.Fprintf(p.log, "aws instance %q: creating %s %s: %v\nthe resource exists, so it is recorded rather than orphaned: check it before relying on it\n",
				p.instance, t.Name, where, waitErr)
		}
		return st, nil
	case waitErr != nil:
		return nil, failure(p.instance, "create", t, where, waitErr)
	}
	why := "it is not visible yet"
	if readErr != nil {
		why = readErr.Error()
	}
	fmt.Fprintf(p.log, "aws instance %q: created %s %s, but it could not be read back (%s); recording the configured values until the next refresh\n",
		p.instance, t.Name, where, why)
	return assumedState(t, region, id, desired.Attrs), nil
}

func (p *Provider) Read(ctx context.Context, current *resource.ResourceState) (*resource.ResourceState, error) {
	t, err := p.lookup(current.Type)
	if err != nil {
		return nil, err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return nil, err
	}
	return p.read(ctx, t, region, id, current.Attributes, p.patience)
}

// read is GetResource with patience for propagation: (nil, nil) when the resource is gone.
func (p *Provider) read(ctx context.Context, t *catalog.Type, region, id string, reference map[string]value.Value, pt patience) (*resource.ResourceState, error) {
	cl := p.clients.get(region)
	var desc *types.ResourceDescription
	found, err := pt.wait(ctx, func() (bool, error) {
		out, err := cl.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String(t.CFN), Identifier: aws.String(id)})
		if isNotFound(err) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		desc = out.ResourceDescription
		return true, nil
	})
	if err != nil {
		return nil, failure(p.instance, "GetResource", t, FormatID(t, region, id), err)
	}
	if !found {
		return nil, nil
	}
	props, err := decodeProperties(aws.ToString(desc.Properties))
	if err != nil {
		return nil, fmt.Errorf("aws instance %q: %s %s: AWS returned properties that are not a JSON object: %w",
			p.instance, t.Name, FormatID(t, region, id), err)
	}
	identifier := aws.ToString(desc.Identifier)
	if identifier == "" {
		identifier = id
	}
	return stateFrom(t, region, identifier, props, reference)
}

// Delete succeeds for a resource that is already gone, whenever that is discovered.
func (p *Provider) Delete(ctx context.Context, current *resource.ResourceState) error {
	t, err := p.lookup(current.Type)
	if err != nil {
		return err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return err
	}
	cl := p.clients.get(region)
	out, err := cl.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{
		TypeName: aws.String(t.CFN), Identifier: aws.String(id), ClientToken: aws.String(p.token()),
	})
	if isNotFound(err) {
		return nil
	}
	if err != nil {
		return failure(p.instance, "DeleteResource", t, current.ProviderID, err)
	}
	if _, err := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "delete")); err != nil && !isNotFound(err) {
		return failure(p.instance, "delete", t, current.ProviderID, err)
	}
	return nil
}

// Import reads once: the ID came from discovery a moment ago. Nested keys come back in snake_case, having no reference.
func (p *Provider) Import(ctx context.Context, resourceType, id string) (*resource.ResourceState, error) {
	t, err := p.lookup(resourceType)
	if err != nil {
		return nil, err
	}
	region, identifier, err := ParseID(t, id)
	if err != nil {
		return nil, err
	}
	st, err := p.read(ctx, t, region, identifier, nil, once)
	if err != nil {
		return nil, err
	}
	if st == nil {
		return nil, fmt.Errorf("aws instance %q: there is no %s %s (%s in %s)", p.instance, t.Name, id, t.CFN, region)
	}
	return st, nil
}
