# aws.transitgatewayattachment

**CloudFormation type:** `AWS::EC2::TransitGatewayAttachment`

Resource Type definition for AWS::EC2::TransitGatewayAttachment

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TransitGatewayAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `Options` |  | `map` | optional, computed, provider-chosen |  | The options for the transit gateway vpc attachment. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | tags map |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | required, replaces on change | aws.transitgateway.Id |  |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
