package ccprov

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

// obj builds a map value from key/value pairs.
func obj(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}

func list(items ...value.Value) value.Value { return value.List(items, value.SourceExplicit) }

func rule(protocol string, from, to int64, spelling map[string]string) value.Value {
	name := func(k string) string {
		if s, ok := spelling[k]; ok {
			return s
		}
		return k
	}
	return obj(name("IpProtocol"), sv(protocol), name("FromPort"), iv(from), name("ToPort"), iv(to), name("CidrIp"), sv("0.0.0.0/0"))
}

var snake = map[string]string{"IpProtocol": "ip_protocol", "FromPort": "from_port", "ToPort": "to_port", "CidrIp": "cidr_ip"}

func TestNestedKeysAreSentUnderAWSNamesWhateverTheSpelling(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	v := list(obj("ip_protocol", sv("tcp"), "FROMPORT", iv(443), "to_port", iv(443), "CidrIp", sv("0.0.0.0/0")))
	got, err := encodeAttr(sg, a, v)
	if err != nil {
		t.Fatal(err)
	}
	want := []any{map[string]any{"IpProtocol": "tcp", "FromPort": int64(443), "ToPort": int64(443), "CidrIp": "0.0.0.0/0"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("encoded = %#v\nwant      %#v", got, want)
	}
}

func TestAnUnknownNestedKeyIsRefusedNamingWhatIsAccepted(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	_, err := encodeAttr(sg, a, list(obj("ip_protocol", sv("tcp"), "port", iv(443))))
	for _, want := range []string{"aws.securitygroup.SecurityGroupIngress[0]", `"port"`, "FromPort (from_port)"} {
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Errorf("err = %v, want it to mention %s", err, want)
		}
	}
	_, err = encodeAttr(sg, a, list(obj("cidr_ip", sv("a"), "CidrIp", sv("b"))))
	if err == nil || !strings.Contains(err.Error(), "twice") {
		t.Errorf("two spellings of one key: err = %v", err)
	}
}

// TestOpaqueValuesAreSentAndReturnedExactly (P5). A policy document's keys are data, not schema names.
func TestOpaqueValuesAreSentAndReturnedExactly(t *testing.T) {
	role := mustType(t, "aws.role")
	a, _ := role.Attribute("AssumeRolePolicyDocument")
	doc := obj("Version", sv("2012-10-17"), "Statement", list(obj("Effect", sv("Allow"), "principal", obj("Service", sv("ec2.amazonaws.com")))))
	got, err := encodeAttr(role, a, doc)
	if err != nil {
		t.Fatal(err)
	}
	stmt := got.(map[string]any)["Statement"].([]any)[0].(map[string]any)
	if _, kept := stmt["principal"]; !kept {
		t.Errorf("an opaque key was translated: %v", stmt)
	}
	back, ok, err := decodeAttr(role, a, jsonDatum(t, `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","principal":{"Service":"ec2.amazonaws.com"}}]}`), nil)
	if err != nil || !ok || !back.Equal(doc) {
		t.Errorf("decoded without a reference = %v, want the keys exactly as AWS returned them", back)
	}
}

func TestReturnedValuesUseTheReferenceSpellingAndDropKeysAWSAdded(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0","Description":""}]`)
	got, ok, err := decodeAttr(sg, a, datum, &ref)
	if err != nil || !ok || !got.Equal(ref) {
		t.Fatalf("decoded = %v, want %v", got, ref)
	}
}

func TestUnorderedListsComeBackInTheReferenceOrder(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 80, 80, nil), rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[
		{"IpProtocol":"udp","FromPort":53,"ToPort":53,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":80,"ToPort":80,"CidrIp":"0.0.0.0/0"}]`)
	got, _, err := decodeAttr(sg, a, datum, &ref)
	if err != nil {
		t.Fatal(err)
	}
	items := got.Raw.([]value.Value)
	refItems := ref.Raw.([]value.Value)
	if len(items) != 3 || !items[0].Equal(refItems[0]) || !items[1].Equal(refItems[1]) {
		t.Fatalf("decoded = %v", got)
	}
	if extra := items[2].Raw.(map[string]value.Value); extra["from_port"].Raw != int64(53) {
		t.Errorf("the item AWS added = %v, want it appended in snake_case", items[2])
	}
}

// TestAChangedItemKeepsItsSpellingAndStillShowsTheChange. Drift inside a list is a real change and must stay visible.
func TestAChangedItemKeepsItsSpellingAndStillShowsTheChange(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(rule("tcp", 80, 80, nil), rule("tcp", 443, 443, snake))
	datum := jsonDatum(t, `[
		{"IpProtocol":"tcp","FromPort":443,"ToPort":444,"CidrIp":"0.0.0.0/0"},
		{"IpProtocol":"tcp","FromPort":80,"ToPort":80,"CidrIp":"0.0.0.0/0"}]`)
	got, _, _ := decodeAttr(sg, a, datum, &ref)
	items, refItems := got.Raw.([]value.Value), ref.Raw.([]value.Value)
	if len(items) != 2 || !items[0].Equal(refItems[0]) {
		t.Fatalf("decoded = %v", got)
	}
	changed := items[1].Raw.(map[string]value.Value)
	if changed["to_port"].Raw != int64(444) || items[1].Equal(refItems[1]) {
		t.Errorf("the changed item = %v, want snake_case keys and to_port 444", items[1])
	}
}

func TestOrderedListsInsideObjectsPairByPosition(t *testing.T) {
	bucket := mustType(t, "aws.bucket")
	a, _ := bucket.Attribute("LifecycleConfiguration")
	ref := obj("rules", list(obj("id", sv("archive"), "status", sv("Enabled"),
		"transitions", list(obj("storage_class", sv("GLACIER"), "transition_in_days", iv(30)), obj("StorageClass", sv("DEEP_ARCHIVE"), "TransitionInDays", iv(180))))))
	datum := jsonDatum(t, `{"Rules":[{"Id":"archive","Status":"Enabled","Prefix":"",
		"Transitions":[{"StorageClass":"GLACIER","TransitionInDays":30},{"StorageClass":"DEEP_ARCHIVE","TransitionInDays":180}]}]}`)
	got, _, err := decodeAttr(bucket, a, datum, &ref)
	if err != nil || !got.Equal(ref) {
		t.Fatalf("decoded = %v, %v\nwant %v", got, err, ref)
	}
}

func TestWithNoReferenceNestedKeysAreSnakeCase(t *testing.T) {
	bucket := mustType(t, "aws.bucket")
	a, _ := bucket.Attribute("VersioningConfiguration")
	got, _, _ := decodeAttr(bucket, a, jsonDatum(t, `{"Status":"Enabled"}`), nil)
	if !got.Equal(obj("status", sv("Enabled"))) {
		t.Errorf("decoded = %v, want {status: Enabled}", got)
	}
}

// TestAScalarAWSReturnsInAnotherFormKeepsTheWrittenForm. `from_port: "443"` and AWS's 443 are one value.
func TestAScalarAWSReturnsInAnotherFormKeepsTheWrittenForm(t *testing.T) {
	sg := mustType(t, "aws.securitygroup")
	a, _ := sg.Attribute("SecurityGroupIngress")
	ref := list(obj("ip_protocol", sv("tcp"), "from_port", sv("443"), "to_port", sv("443"), "cidr_ip", sv("0.0.0.0/0")))
	got, _, _ := decodeAttr(sg, a, jsonDatum(t, `[{"IpProtocol":"tcp","FromPort":443,"ToPort":443,"CidrIp":"0.0.0.0/0"}]`), &ref)
	if !got.Equal(ref) {
		t.Errorf("decoded = %v, want %v", got, ref)
	}
}

// TestNestedValuesConvergeThroughTheProvider: the user's spelling survives create and every later read, although AWS
// adds keys and reorders the list.
func TestNestedValuesConvergeThroughTheProvider(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	cfg := awstest.FakeType(mustType(t, "aws.securitygroup"), "sg-", nil)
	cfg.OnRead = func(props map[string]any) {
		rules, _ := props["SecurityGroupIngress"].([]any)
		for _, r := range rules {
			r.(map[string]any)["Description"] = ""
		}
		slices.Reverse(rules)
	}
	fake.Register(cfg)
	ingress := list(rule("tcp", 80, 80, snake), rule("tcp", 443, 443, map[string]string{"FromPort": "fromPort"}))

	st, err := p.Create(ctx, desired("aws.securitygroup", map[string]value.Value{
		"region": sv("us-east-1"), "GroupDescription": sv("web"), "SecurityGroupIngress": ingress,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if got := st.Attributes["SecurityGroupIngress"]; !got.Equal(ingress) {
		t.Fatalf("after create = %v\nwant %v: a plan would never converge", got, ingress)
	}
	again, err := p.Read(ctx, &resource.ResourceState{Type: st.Type, ProviderID: st.ProviderID, Attributes: st.Attributes})
	if err != nil || !again.Attributes["SecurityGroupIngress"].Equal(ingress) {
		t.Fatalf("after read = %v, %v", again.Attributes["SecurityGroupIngress"], err)
	}
	id := strings.TrimPrefix(st.ProviderID, "us-east-1/")
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::SecurityGroup", id)
	first := stored["SecurityGroupIngress"].([]any)[0].(map[string]any)
	if _, ok := first["FromPort"]; !ok {
		t.Errorf("AWS was sent %v, want AWS's names", first)
	}
}
