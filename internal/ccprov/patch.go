package ccprov

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

// patchOps is the RFC 6902 patch from current to desired, one operation per top-level property whose value differs
// under AWS's names. Sorted by path, so a request is stable.
func patchOps(t *catalog.Type, current, desired map[string]value.Value) ([]map[string]any, error) {
	attrs := append([]*catalog.Attribute(nil), t.Attributes...)
	sort.Slice(attrs, func(i, j int) bool { return attrs[i].Name < attrs[j].Name })
	var ops []map[string]any
	for _, a := range attrs {
		v, set := desired[a.Name]
		if !set || (a.Computed && !a.Optional) {
			continue
		}
		want, err := encodeAttr(t, a, v)
		if err != nil {
			return nil, err
		}
		op := "add"
		if cur, has := current[a.Name]; has && cur.Known {
			if have, err := encodeAttr(t, a, cur); err == nil {
				if sameJSON(want, have) {
					continue
				}
				op = "replace"
			}
		}
		ops = append(ops, map[string]any{"op": op, "path": "/" + a.Name, "value": want})
	}
	return ops, nil
}

func sameJSON(a, b any) bool {
	ja, errA := json.Marshal(a)
	jb, errB := json.Marshal(b)
	return errA == nil && errB == nil && bytes.Equal(ja, jb)
}

// Update patches what changed, waits for the request, and reads the resource back with configuration as the
// reference. When nothing differs under AWS's names (a spelling change) it only reads back, so state takes the new
// spelling.
//
// current, as infrena hands it to Update, is the refreshed observation from immediately before planning (as of
// infrena v0.7.1, commit 1399f20; before that it was the last state persisted to disk, which could be older than
// what AWS held right now, and this method worked around it with an extra read — found running the e2e suite,
// commit c7ff98e). The patch is computed straight from current.Attributes. current.Attributes still supplies the
// reference a write-only property's last known value is carried forward from (using desired.Attrs there would make
// a just-configured write-only value look unchanged, since nothing else ever reports it back).
func (p *Provider) Update(ctx context.Context, current *resource.ResourceState, desired *resource.DesiredResource) (*resource.ResourceState, error) {
	t, err := p.lookup(current.Type)
	if err != nil {
		return nil, err
	}
	region, id, err := ParseID(t, current.ProviderID)
	if err != nil {
		return nil, err
	}
	if !t.HasUpdate {
		return nil, fmt.Errorf("%s (%s) has no update handler, so every change needs a replacement; infrena was sent an update, which is a defect in the catalog's force-new flags",
			t.Name, t.CFN)
	}
	ops, err := patchOps(t, current.Attributes, desired.Attrs)
	if err != nil {
		return nil, err
	}
	if len(ops) > 0 {
		doc, err := json.Marshal(ops)
		if err != nil {
			return nil, err
		}
		cl := p.clients.get(region)
		out, err := cl.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
			TypeName: aws.String(t.CFN), Identifier: aws.String(id), PatchDocument: aws.String(string(doc)), ClientToken: aws.String(p.token()),
		})
		if err != nil {
			return nil, updateFailure(p.instance, "UpdateResource", t, current.ProviderID, ops, err)
		}
		ctx = context.WithoutCancel(ctx) // from here AWS may have acted
		if _, err := p.pacing.await(ctx, cl, out.ProgressEvent, timeoutFor(t, "update")); err != nil {
			return nil, updateFailure(p.instance, "update", t, current.ProviderID, ops, err)
		}
	}
	st, err := p.read(ctx, t, region, id, desired.Attrs, p.patience)
	if err != nil {
		return nil, err
	}
	if st == nil {
		return nil, fmt.Errorf("aws instance %q: %s %s no longer exists after the update", p.instance, t.Name, current.ProviderID)
	}
	return st, nil
}

// updateFailure names the patched paths when AWS refuses to change them in place, which is the schema being wrong
// about what is create-only. Only paths: the values may be secret.
func updateFailure(instance, action string, t *catalog.Type, where string, ops []map[string]any, err error) error {
	wrapped := failure(instance, action, t, where, err)
	switch errorCode(err) {
	case "NotUpdatable", "NotUpdatableException":
		paths := make([]string, len(ops))
		for i, op := range ops {
			paths[i] = op["path"].(string)
		}
		return fmt.Errorf("%w\nAWS will not change %s in place, although the schema does not mark it create-only; replace the resource, and report the schema",
			wrapped, strings.Join(paths, ", "))
	}
	return wrapped
}
