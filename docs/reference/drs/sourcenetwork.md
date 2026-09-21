# aws.sourcenetwork

**CloudFormation type:** `AWS::DRS::SourceNetwork`

A Source Network resource represents a VPC that is protected by AWS Elastic Disaster Recovery.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::DRS::SourceNetwork)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Source Network. |
| `OriginAccountID` | origin_account_id | `string` | required, replaces on change |  | The account ID containing the VPC to protect. |
| `OriginRegion` | origin_region | `string` | required, replaces on change |  | The region containing the VPC to protect. |
| `SourceNetworkID` | source_network_id | `string` | computed |  | The ID of the Source Network. |
| `Tags` |  | `map` | tags map |  | A set of tags associated with the Source Network. |
| `VpcID` | vpc_id | `string` | required, replaces on change |  | The VPC ID to protect. |

Supports update: yes

Discovery: supported
