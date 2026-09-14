package ccprov

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/infrena/infrena/pkg/value"
)

func jsonDatum(t *testing.T, doc string) any {
	t.Helper()
	dec := json.NewDecoder(strings.NewReader(doc))
	dec.UseNumber()
	var out any
	if err := dec.Decode(&out); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestValuesCrossJSONAndComeBackEqual(t *testing.T) {
	v := value.Map(map[string]value.Value{
		"name": sv("web"), "count": iv(3), "ratio": value.Float(0.5, value.SourceExplicit), "on": value.Bool(true, value.SourceExplicit),
		"items": value.List([]value.Value{sv("a"), iv(2)}, value.SourceExplicit),
	}, value.SourceExplicit)
	raw, err := json.Marshal(plain(v))
	if err != nil {
		t.Fatal(err)
	}
	back, ok := infer(jsonDatum(t, string(raw)))
	if !ok || !back.Equal(v) {
		t.Fatalf("round trip = %v, want %v", back, v)
	}
}

func TestNullsAreAbsent(t *testing.T) {
	got, ok := infer(jsonDatum(t, `{"a": null, "b": [1, null], "c": "x"}`))
	items := got.Raw.(map[string]value.Value)
	if !ok || len(items) != 2 || len(items["b"].Raw.([]value.Value)) != 1 {
		t.Fatalf("infer = %v", got)
	}
	if _, ok := infer(nil); ok {
		t.Error("a top-level null was recorded")
	}
}

// TestReturnedValuesBecomeTheKindTheSchemaDeclares. A property typed as a union is a string attribute (spec §3.1):
// whatever AWS returns for it is recorded as JSON text, never refused.
func TestReturnedValuesBecomeTheKindTheSchemaDeclares(t *testing.T) {
	vpc, db := mustType(t, "aws.vpc"), mustType(t, "aws.dbinstance")
	cidr, _ := vpc.Attribute("CidrBlock")
	netmask, _ := vpc.Attribute("Ipv4NetmaskLength")
	dns, _ := vpc.Attribute("EnableDnsSupport")
	class, _ := db.Attribute("DBInstanceClass")

	if got, ok, err := decodeAttr(vpc, cidr, jsonDatum(t, `{"b": 1, "a": [true]}`), nil); err != nil || !ok || got.Raw != `{"a":[true],"b":1}` {
		t.Errorf("object into a string attribute = %v, %v", got, err)
	}
	if got, _, err := decodeAttr(vpc, netmask, jsonDatum(t, `16.0`), nil); err != nil || got.Raw != int64(16) {
		t.Errorf("16.0 into an integer attribute = %v, %v", got, err)
	}
	if got, _, err := decodeAttr(vpc, dns, jsonDatum(t, `"true"`), nil); err != nil || got.Raw != true {
		t.Errorf(`"true" into a boolean attribute = %v, %v`, got, err)
	}
	if _, _, err := decodeAttr(vpc, netmask, jsonDatum(t, `16.5`), nil); err == nil || !strings.Contains(err.Error(), "aws.vpc.Ipv4NetmaskLength") {
		t.Errorf("16.5 into an integer attribute: err = %v", err)
	}
	if got, _, err := decodeAttr(db, class, jsonDatum(t, `7`), nil); err != nil || got.Raw != "7" {
		t.Errorf("a number into a string attribute = %v, %v", got, err)
	}
}

func TestTheDesiredStateHoldsOnlyWhatConfigurationMaySet(t *testing.T) {
	body, err := desiredJSON(mustType(t, "aws.vpc"), map[string]value.Value{
		"region": sv("us-east-1"), "CidrBlock": sv("10.0.0.0/16"), "VpcId": sv("vpc-from-state"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(body) != 1 || body["CidrBlock"] != "10.0.0.0/16" {
		t.Errorf("desired state = %v, want only CidrBlock: never the plugin's region, never a read-only property", body)
	}
}

// TestJSONTextAWSReturnsAsAnObjectKeepsItsWrittenForm. A policy document is a string attribute, written as JSON text;
// AWS returns it as an object. Equivalent documents must keep the written text, and a changed one must not.
func TestJSONTextAWSReturnsAsAnObjectKeepsItsWrittenForm(t *testing.T) {
	vpc := mustType(t, "aws.vpc")
	a, _ := vpc.Attribute("CidrBlock") // any string attribute with no shape
	written := sv(`{ "Version": "2012-10-17",
  "Statement": [ { "Effect": "Allow", "Action": "sts:AssumeRole" } ] }`)
	same := jsonDatum(t, `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow"}],"Version":"2012-10-17"}`)
	if got, ok, err := decodeAttr(vpc, a, same, &written); err != nil || !ok || !got.Equal(written) {
		t.Errorf("an equivalent document = %v, %v; want the written text", got, err)
	}
	changed := jsonDatum(t, `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Deny"}],"Version":"2012-10-17"}`)
	if got, _, _ := decodeAttr(vpc, a, changed, &written); got.Equal(written) || !strings.Contains(got.Raw.(string), "Deny") {
		t.Errorf("a changed document = %v; want AWS's version, so the drift shows", got)
	}
}

// TestNestedJSONTextAWSMinifiesKeepsItsWrittenForm. An ECR lifecycle policy is a JSON document held in a plain string
// inside an object; AWS returns it minified.
func TestNestedJSONTextAWSMinifiesKeepsItsWrittenForm(t *testing.T) {
	bucket := mustType(t, "aws.bucket")
	a, _ := bucket.Attribute("VersioningConfiguration") // an object shape with a scalar property
	written := value.Map(map[string]value.Value{"status": sv(`{
  "rules": [ { "rulePriority": 1, "action": { "type": "expire" } } ]
}`)}, value.SourceExplicit)
	minified := jsonDatum(t, `{"Status":"{\"rules\":[{\"action\":{\"type\":\"expire\"},\"rulePriority\":1}]}"}`)
	if got, ok, err := decodeAttr(bucket, a, minified, &written); err != nil || !ok || !got.Equal(written) {
		t.Errorf("a minified equivalent = %v, %v; want the written text", got, err)
	}
	changed := jsonDatum(t, `{"Status":"{\"rules\":[{\"action\":{\"type\":\"expire\"},\"rulePriority\":2}]}"}`)
	if got, _, _ := decodeAttr(bucket, a, changed, &written); got.Equal(written) {
		t.Error("a changed document was reported as unchanged")
	}
	plain := value.Map(map[string]value.Value{"status": sv("Enabled")}, value.SourceExplicit)
	if got, _, _ := decodeAttr(bucket, a, jsonDatum(t, `{"Status":"Suspended"}`), &plain); got.Equal(plain) {
		t.Error("an ordinary string change was hidden")
	}
}
