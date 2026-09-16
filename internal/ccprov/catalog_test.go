package ccprov

import (
	"testing"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

func scalar() *catalog.Shape { return &catalog.Shape{Kind: catalog.ShapeScalar} }
func opaque() *catalog.Shape { return &catalog.Shape{Kind: catalog.ShapeOpaque} }
func object(props map[string]*catalog.Shape) *catalog.Shape {
	return &catalog.Shape{Kind: catalog.ShapeObject, Props: props}
}
func array(item *catalog.Shape, unordered bool) *catalog.Shape {
	return &catalog.Shape{Kind: catalog.ShapeArray, Item: item, Unordered: unordered}
}

// testCatalog is every shape of type the provider handles:
//
//	aws.vpc              regional, provider-chosen defaults, a write-only property, tags as a map
//	aws.subnet           regional, and one of the two types whose cloud-owned answer only EC2 has
//	aws.securitygroup    an unordered list of objects (nested spelling, AWS-added keys, order)
//	aws.bucket           nested objects inside ordered lists inside objects
//	aws.role             global, an opaque policy document, an unordered list holding opaque values
//	aws.dbinstance       sensitive and write-only, a long create timeout, a computed object
//	aws.test.regioned    its own Region property (so the plugin's is aws_region); no update handler, no list handler
//	aws.test.child       a composite identifier; listing needs a parent model
func testCatalog() *catalog.Catalog {
	return &catalog.Catalog{
		Bundle:          "test",
		DiscoverDefault: []string{"aws.vpc", "aws.role"},
		Types: []*catalog.Type{
			{
				Name: "aws.vpc", CFN: "AWS::EC2::VPC", RegionAttr: "region", Identifier: []string{"VpcId"},
				WriteOnly: []string{"Ipv4NetmaskLength"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "CidrBlock", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"cidr", "cidr_block"}},
					{Name: "EnableDnsSupport", Kind: "boolean", Optional: true, Computed: true, Aliases: []string{"enable_dns_support"}},
					{Name: "InstanceTenancy", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"instance_tenancy"}},
					{Name: "Ipv4NetmaskLength", Kind: "integer", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"ipv4_netmask_length"}},
					{Name: "VpcId", Kind: "string", Computed: true, Aliases: []string{"vpc_id"}},
					{Name: "DefaultSecurityGroup", Kind: "string", Computed: true, Aliases: []string{"default_security_group"}},
					{Name: "CidrBlockAssociations", Kind: "list", Computed: true, Aliases: []string{"cidr_block_associations"}, Shape: array(scalar(), true)},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.subnet", CFN: "AWS::EC2::Subnet", RegionAttr: "region", Identifier: []string{"SubnetId"},
				HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "SubnetId", Kind: "string", Computed: true, Aliases: []string{"subnet_id"}},
					{Name: "VpcId", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"vpc", "vpc_id"}},
					{Name: "CidrBlock", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"cidr", "cidr_block"}},
					{Name: "AvailabilityZone", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"az", "availability_zone"}},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.securitygroup", CFN: "AWS::EC2::SecurityGroup", RegionAttr: "region", Identifier: []string{"GroupId"},
				HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "GroupDescription", Kind: "string", Required: true, ForceNew: true, Aliases: []string{"description", "group_description"}},
					{Name: "GroupName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "group_name"}},
					{Name: "VpcId", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"vpc_id"}},
					{Name: "GroupId", Kind: "string", Computed: true, Aliases: []string{"group_id"}},
					{Name: "SecurityGroupIngress", Kind: "list", Optional: true, Computed: true, Aliases: []string{"ingress", "security_group_ingress"},
						Shape: array(object(map[string]*catalog.Shape{
							"IpProtocol": scalar(), "FromPort": scalar(), "ToPort": scalar(), "CidrIp": scalar(), "Description": scalar(),
						}), true)},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.bucket", CFN: "AWS::S3::Bucket", RegionAttr: "region", Identifier: []string{"BucketName"},
				HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "BucketName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "bucket_name"}},
					{Name: "Arn", Kind: "string", Computed: true},
					{Name: "VersioningConfiguration", Kind: "map", Optional: true, Computed: true, Aliases: []string{"versioning_configuration"},
						Shape: object(map[string]*catalog.Shape{"Status": scalar()})},
					{Name: "LifecycleConfiguration", Kind: "map", Optional: true, Computed: true, Aliases: []string{"lifecycle_configuration"},
						Shape: object(map[string]*catalog.Shape{
							"Rules": array(object(map[string]*catalog.Shape{
								"Id": scalar(), "Status": scalar(), "ExpirationInDays": scalar(),
								"Transitions": array(object(map[string]*catalog.Shape{"StorageClass": scalar(), "TransitionInDays": scalar()}), false),
							}), false),
						})},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.role", CFN: "AWS::IAM::Role", Identifier: []string{"RoleName"}, HasUpdate: true, HasList: true, TagsAsMap: "Tags",
				Attributes: []*catalog.Attribute{
					{Name: "RoleName", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"name", "role_name"}},
					{Name: "Arn", Kind: "string", Computed: true},
					{Name: "AssumeRolePolicyDocument", Kind: "map", Required: true, Aliases: []string{"assume_role_policy", "assume_role_policy_document"}, Shape: opaque()},
					{Name: "Policies", Kind: "list", Optional: true, Computed: true,
						Shape: array(object(map[string]*catalog.Shape{"PolicyName": scalar(), "PolicyDocument": opaque()}), true)},
					{Name: "MaxSessionDuration", Kind: "integer", Optional: true, Computed: true, Aliases: []string{"max_session_duration"}},
					{Name: "Tags", Kind: "map", Optional: true, Computed: true, Shape: opaque()},
				},
			},
			{
				Name: "aws.dbinstance", CFN: "AWS::RDS::DBInstance", RegionAttr: "region", Identifier: []string{"DBInstanceIdentifier"},
				WriteOnly: []string{"MasterUserPassword"}, HasUpdate: true, HasList: true, Timeouts: map[string]int{"create": 2160},
				Attributes: []*catalog.Attribute{
					{Name: "DBInstanceIdentifier", Kind: "string", Optional: true, Computed: true, ForceNew: true, Aliases: []string{"identifier", "db_instance_identifier"}},
					{Name: "DBInstanceClass", Kind: "string", Optional: true, Computed: true, Aliases: []string{"instance_class", "db_instance_class"}},
					{Name: "MasterUserPassword", Kind: "string", Optional: true, Computed: true, Sensitive: true, Aliases: []string{"password", "master_user_password"}},
					{Name: "Endpoint", Kind: "map", Computed: true, Shape: object(map[string]*catalog.Shape{"Address": scalar(), "Port": scalar()})},
				},
			},
			{
				Name: "aws.test.regioned", CFN: "AWS::Test::Regioned", RegionAttr: "aws_region", Identifier: []string{"Id"},
				Attributes: []*catalog.Attribute{
					{Name: "Id", Kind: "string", Computed: true},
					{Name: "Name", Kind: "string", Required: true, ForceNew: true},
					{Name: "Region", Kind: "string", Optional: true, Computed: true, ForceNew: true},
				},
			},
			{
				Name: "aws.test.child", CFN: "AWS::Test::Child", RegionAttr: "region", Identifier: []string{"ParentId", "ChildId"},
				HasUpdate: true, HasList: true, ListNeedsModel: true,
				Attributes: []*catalog.Attribute{
					{Name: "ParentId", Kind: "string", Required: true, ForceNew: true, Aliases: []string{"parent_id"}},
					{Name: "ChildId", Kind: "string", Computed: true, Aliases: []string{"child_id"}},
					{Name: "Setting", Kind: "string", Optional: true, Computed: true},
				},
			},
		},
	}
}

func mustType(t *testing.T, name string) *catalog.Type {
	t.Helper()
	typ, ok := testCatalog().Lookup(name)
	if !ok {
		t.Fatalf("test catalog has no %s", name)
	}
	return typ
}

// TestTheTestCatalogIsOneInfrenaWouldLoad. A test catalog infrena would refuse would make every test built on it moot.
func TestTheTestCatalogIsOneInfrenaWouldLoad(t *testing.T) {
	for _, d := range testCatalog().Definitions() {
		if err := d.Validate(); err != nil {
			t.Errorf("%s: %v", d.Type, err)
		}
	}
}
