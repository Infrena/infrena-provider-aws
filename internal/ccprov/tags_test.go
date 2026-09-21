package ccprov

import (
	"reflect"
	"strings"
	"testing"

	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

func TestTagsAreAMapInConfigurationAndAListInAWS(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	got, err := encodeAttr(vpc, a, obj("team", sv("platform"), "app", sv("web")))
	if err != nil {
		t.Fatal(err)
	}
	want := []any{map[string]any{"Key": "app", "Value": "web"}, map[string]any{"Key": "team", "Value": "platform"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("encoded = %#v", got)
	}
	if _, err := encodeAttr(vpc, a, list(sv("team"))); err == nil || !strings.Contains(err.Error(), "map of tag names") {
		t.Errorf("tags as a list: err = %v", err)
	}
}

// TestSystemTagsAreNeverReported. `aws:` tags are AWS's; configuration cannot set them, so reporting them is a diff
// nobody can resolve.
func TestSystemTagsAreNeverReported(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	got, _, err := decodeAttr(vpc, a, jsonDatum(t, `[{"Key":"team","Value":"platform"},{"Key":"aws:cloudformation:stack-name","Value":"x"}]`), nil)
	if err != nil || !got.Equal(obj("team", sv("platform"))) {
		t.Errorf("decoded = %v, %v", got, err)
	}
}

func TestATagValueWrittenAsANumberKeepsItsForm(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("Tags")
	ref := obj("cost_centre", iv(1234))
	sent, _ := encodeAttr(vpc, a, ref)
	if sent.([]any)[0].(map[string]any)["Value"] != "1234" {
		t.Errorf("sent %v, want the value as text", sent)
	}
	got, _, _ := decodeAttr(vpc, a, jsonDatum(t, `[{"Key":"cost_centre","Value":"1234"}]`), &ref)
	if !got.Equal(ref) {
		t.Errorf("decoded = %v, want %v", got, ref)
	}
}

func TestTagsRoundTripThroughTheProviderAndDriftShows(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	tags := obj("team", sv("platform"))
	st, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "Tags": tags}))
	if err != nil || !st.Attributes["Tags"].Equal(tags) {
		t.Fatalf("after create Tags = %v, %v", st.Attributes["Tags"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	stored["Tags"] = []any{map[string]any{"Key": "team", "Value": "someone-else"}}
	fake.Put("us-east-1", "AWS::EC2::VPC", id, stored)
	got, err := p.Read(ctx, st)
	if err != nil || got.Attributes["Tags"].Equal(tags) {
		t.Fatalf("after drift Tags = %v, %v; want the changed value", got.Attributes["Tags"], err)
	}
}

func TestAnEmptyTagMapStaysEmpty(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	none := value.Map(map[string]value.Value{}, value.SourceExplicit)
	st, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "Tags": none}))
	if err != nil || !st.Attributes["Tags"].Equal(none) {
		t.Fatalf("after create Tags = %v, %v", st.Attributes["Tags"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	delete(stored, "Tags")
	fake.Put("us-east-1", "AWS::EC2::VPC", id, stored)
	got, _ := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: st.ProviderID, Attributes: st.Attributes})
	if !got.Attributes["Tags"].Equal(none) {
		t.Errorf("AWS omitting an empty tag list: Tags = %v, want {}", got.Attributes["Tags"])
	}
}
