# aws.transitgatewayvpcattachment

**CloudFormation type:** `AWS::EC2::TransitGatewayVpcAttachment`

Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TransitGatewayVpcAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddSubnetIds` | add_subnet_ids | `list` | optional, computed, provider-chosen, write-only | aws.subnet.SubnetId |  |
| `Id` |  | `string` | computed |  |  |
| `Options` |  | `map` | optional, computed, provider-chosen |  | The options for the transit gateway vpc attachment. |
| `RemoveSubnetIds` | remove_subnet_ids | `list` | optional, computed, provider-chosen, write-only | aws.subnet.SubnetId |  |
| `SubnetIds` | subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id |  |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
