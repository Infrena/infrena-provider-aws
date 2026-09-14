package awsprov

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/infrata/infrata/pkg/value"
)

// reservedPrefix is AWS's own: tags under it cannot be edited or deleted, so they are neither reported
// (every plan would propose removing them) nor accepted.
const reservedPrefix = "aws:"

func tagsFrom(attrs map[string]value.Value) (map[string]string, error) {
	v, ok := attrs["tags"]
	if !ok {
		return map[string]string{}, nil
	}
	items, isMap := v.Raw.(map[string]value.Value)
	if v.Kind != value.KindMap || !isMap {
		return nil, fmt.Errorf("tags must be a map of strings, got %s", v.Kind)
	}
	out := make(map[string]string, len(items))
	for k, item := range items {
		if strings.HasPrefix(k, reservedPrefix) {
			return nil, fmt.Errorf("tag %q: keys starting %q are reserved by AWS and cannot be set", k, reservedPrefix)
		}
		text, isString := item.AsString()
		if !isString {
			return nil, fmt.Errorf("tag %q must be a string, got %s", k, item.Kind)
		}
		out[k] = text
	}
	return out, nil
}

func fromAWSTags(tags []types.Tag) map[string]string {
	out := map[string]string{}
	for _, t := range tags {
		if k := aws.ToString(t.Key); !strings.HasPrefix(k, reservedPrefix) {
			out[k] = aws.ToString(t.Value)
		}
	}
	return out
}

// putTags sets `tags`, or removes it when there are none: an empty map where configuration has no
// `tags:` would plan "removed from configuration" on every run.
func putTags(attrs map[string]value.Value, tags map[string]string) {
	if len(tags) == 0 {
		delete(attrs, "tags")
		return
	}
	items := make(map[string]value.Value, len(tags))
	for k, v := range tags {
		items[k] = str(v)
	}
	attrs["tags"] = value.Map(items, value.SourceProvider)
}

func toAWSTags(tags map[string]string) []types.Tag {
	keys := make([]string, 0, len(tags))
	for k := range tags {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := make([]types.Tag, 0, len(keys))
	for _, k := range keys {
		out = append(out, types.Tag{Key: aws.String(k), Value: aws.String(tags[k])})
	}
	return out
}

// tagSpecs tags a resource in its create call, which AWS applies atomically: a resource is created with
// its tags or not at all. No follow-up CreateTags, so no window where it exists untagged.
func tagSpecs(rt types.ResourceType, tags map[string]string) []types.TagSpecification {
	if len(tags) == 0 {
		return nil
	}
	return []types.TagSpecification{{ResourceType: rt, Tags: toAWSTags(tags)}}
}

// syncTags makes AWS's tags equal desired's: sets what is new or changed, deletes what is gone.
func (p *Provider) syncTags(ctx context.Context, region, awsID string, current, desired map[string]value.Value) error {
	have, err := tagsFrom(current)
	if err != nil {
		return err
	}
	want, err := tagsFrom(desired)
	if err != nil {
		return err
	}
	set := map[string]string{}
	for k, v := range want {
		if old, ok := have[k]; !ok || old != v {
			set[k] = v
		}
	}
	var remove []types.Tag
	for k := range have {
		if _, ok := want[k]; !ok {
			remove = append(remove, types.Tag{Key: aws.String(k)})
		}
	}
	sort.Slice(remove, func(i, j int) bool { return *remove[i].Key < *remove[j].Key })
	client := p.clients.ec2(region)
	if len(set) > 0 {
		if _, err := client.CreateTags(ctx, &ec2.CreateTagsInput{Resources: []string{awsID}, Tags: toAWSTags(set)}); err != nil {
			return p.failed("CreateTags", region, awsID, err)
		}
	}
	if len(remove) > 0 {
		if _, err := client.DeleteTags(ctx, &ec2.DeleteTagsInput{Resources: []string{awsID}, Tags: remove}); err != nil {
			return p.failed("DeleteTags", region, awsID, err)
		}
	}
	return nil
}
