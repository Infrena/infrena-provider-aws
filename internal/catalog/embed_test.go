package catalog

import (
	"encoding/json"
	"testing"
)

func TestTheEmbeddedCatalogLoadsAndValidates(t *testing.T) {
	cat, err := Embedded()
	if err != nil {
		t.Fatal(err)
	}
	if len(cat.Types) < 1500 {
		t.Fatalf("embedded catalog has %d types; expected the ~1,584 Cloud Control supports", len(cat.Types))
	}
	defs := cat.Definitions()
	for _, d := range defs {
		if err := d.Validate(); err != nil {
			t.Fatalf("%s: %v", d.Type, err)
		}
	}
	for name, cfn := range map[string]string{
		"aws.vpc": "AWS::EC2::VPC", "aws.subnet": "AWS::EC2::Subnet", "aws.securitygroup": "AWS::EC2::SecurityGroup",
		"aws.role": "AWS::IAM::Role", "aws.ec2.instance": "AWS::EC2::Instance", "aws.s3.bucket": "AWS::S3::Bucket",
	} {
		if typ, ok := cat.Lookup(name); !ok || typ.CFN != cfn {
			t.Errorf("%s = %+v", name, typ)
		}
	}
	for _, name := range cat.DiscoverDefault {
		if _, ok := cat.Lookup(name); !ok {
			t.Errorf("discover default names %s, which the catalog lacks", name)
		}
	}
	raw, _ := json.Marshal(defs)
	t.Logf("%d definitions, %d bytes as JSON", len(defs), len(raw))
}
