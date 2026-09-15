# aws.carriergateway

**CloudFormation type:** `AWS::EC2::CarrierGateway`

Resource Type definition for Carrier Gateway which describes the Carrier Gateway resource

Region attribute: `region`

**Import ID:** `<region>/CarrierGatewayId` (AWS::EC2::CarrierGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CarrierGatewayId` | carrier_gateway_id | `string` | computed |  | The ID of the carrier gateway. |
| `OwnerId` | owner_id | `string` | computed |  | The ID of the owner. |
| `State` |  | `string` | computed |  | The state of the carrier gateway. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the carrier gateway. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: yes

Discovery: supported
