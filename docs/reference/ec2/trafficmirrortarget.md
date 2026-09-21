# aws.trafficmirrortarget

**CloudFormation type:** `AWS::EC2::TrafficMirrorTarget`

The description of the Traffic Mirror target.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TrafficMirrorTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the Traffic Mirror target. |
| `GatewayLoadBalancerEndpointId` | gateway_load_balancer_endpoint_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the Gateway Load Balancer endpoint. |
| `Id` |  | `string` | computed |  |  |
| `NetworkInterfaceId` | network_interface_id | `string` | optional, computed, provider-chosen, replaces on change | aws.networkinterface.Id | The network interface ID that is associated with the target. |
| `NetworkLoadBalancerArn` | network_load_balancer_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target. |
| `Tags` |  | `map` | tags map |  | The tags to assign to the Traffic Mirror target. |

Supports update: yes

Discovery: supported
