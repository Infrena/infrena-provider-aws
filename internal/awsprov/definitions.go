package awsprov

import (
	"github.com/infrata/infrata/pkg/schema"
	"github.com/infrata/infrata/pkg/value"
)

const (
	typeVPC    = "aws.vpc"
	typeSubnet = "aws.subnet"
)

// regionAttr is declared by every regional type. ForceNew because an AWS resource cannot move
// between regions; the default comes from the instance's `defaults: {region: …}`, applied by
// infrata at compile time and never sent to this plugin.
var regionAttr = schema.Attribute{Kind: value.KindString, Required: true, ForceNew: true,
	Description: "AWS region, e.g. us-east-1. Usually set once with the provider's defaults: {region: …}"}

var tagsAttr = schema.Attribute{Kind: value.KindMap,
	Description: "Tags. Keys starting aws: are reserved by AWS and are neither reported nor accepted"}

func definitions() []*schema.ResourceDefinition {
	return []*schema.ResourceDefinition{
		{
			Type:        typeVPC,
			Description: "An Amazon VPC.",
			Attributes: map[string]schema.Attribute{
				"region": regionAttr,
				// The primary CIDR block cannot be changed or disassociated after creation.
				"cidr":       {Kind: value.KindString, Required: true, ForceNew: true, Description: "Primary IPv4 CIDR block"},
				"tags":       tagsAttr,
				"id":         {Kind: value.KindString, Computed: true, Description: "VPC ID, e.g. vpc-0abc123"},
				"owner_id":   {Kind: value.KindString, Computed: true, Description: "Owning AWS account ID"},
				"is_default": {Kind: value.KindBool, Computed: true, Description: "Whether this is the region's default VPC"},
			},
			Capabilities: schema.Capabilities{Create: true, Read: true, Update: true, Delete: true, Import: true},
			ImportID:     schema.ImportSpec{Description: "<region>/<vpc id>, e.g. us-east-1/vpc-0abc123"},
		},
		{
			Type:        typeSubnet,
			Description: "A subnet inside an Amazon VPC.",
			Attributes: map[string]schema.Attribute{
				"region": regionAttr,
				"vpc_id": {Kind: value.KindString, Required: true, ForceNew: true,
					Description: "VPC to create the subnet in, usually ${<vpc>.id}"},
				"cidr": {Kind: value.KindString, Required: true, ForceNew: true, Description: "IPv4 CIDR block"},
				// Required, not optional-and-chosen-by-AWS: AWS always reports it, so an omitted
				// availability zone would plan a replacement on every run.
				"availability_zone": {Kind: value.KindString, Required: true, ForceNew: true,
					Description: "Availability zone, e.g. us-east-1a"},
				"map_public_ip_on_launch": {Kind: value.KindBool, Default: false,
					Description: "Give instances launched here a public IPv4 address"},
				"tags":     tagsAttr,
				"id":       {Kind: value.KindString, Computed: true, Description: "Subnet ID, e.g. subnet-0abc123"},
				"arn":      {Kind: value.KindString, Computed: true, Description: "Subnet ARN"},
				"owner_id": {Kind: value.KindString, Computed: true, Description: "Owning AWS account ID"},
			},
			// A pre-flight hint only: satisfied by any aws.vpc in the same provider instance,
			// whatever its region, and never by state. The real dependency is vpc_id.
			Requirements: []schema.Requirement{{
				Name: "vpc", Types: []string{typeVPC},
				Description: "A subnet must be created inside a VPC",
			}},
			Capabilities: schema.Capabilities{Create: true, Read: true, Update: true, Delete: true, Import: true},
			ImportID:     schema.ImportSpec{Description: "<region>/<subnet id>, e.g. us-east-1/subnet-0abc123"},
		},
	}
}
