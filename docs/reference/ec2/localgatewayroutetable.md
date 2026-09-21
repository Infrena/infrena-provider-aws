# aws.localgatewayroutetable

**CloudFormation type:** `AWS::EC2::LocalGatewayRouteTable`

Resource Type definition for Local Gateway Route Table which describes a route table for a local gateway.

Region attribute: `region`

**Import ID:** `<region>/LocalGatewayRouteTableId` (AWS::EC2::LocalGatewayRouteTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocalGatewayId` | local_gateway_id | `string` | required, replaces on change |  | The ID of the local gateway. |
| `LocalGatewayRouteTableArn` | local_gateway_route_table_arn | `string` | computed |  | The ARN of the local gateway route table. |
| `LocalGatewayRouteTableId` | local_gateway_route_table_id | `string` | computed |  | The ID of the local gateway route table. |
| `Mode` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The mode of the local gateway route table. |
| `OutpostArn` | outpost_arn | `string` | computed |  | The ARN of the outpost. |
| `OwnerId` | owner_id | `string` | computed |  | The owner of the local gateway route table. |
| `State` |  | `string` | computed |  | The state of the local gateway route table. |
| `Tags` |  | `map` | tags map |  | The tags for the local gateway route table. |

Supports update: yes

Discovery: supported
