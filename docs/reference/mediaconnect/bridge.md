# aws.bridge

**CloudFormation type:** `AWS::MediaConnect::Bridge`

Resource schema for AWS::MediaConnect::Bridge

Region attribute: `region`

**Import ID:** `<region>/BridgeArn` (AWS::MediaConnect::Bridge)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BridgeArn` | bridge_arn | `string` | computed |  | The Amazon Resource Number (ARN) of the bridge. |
| `BridgeState` | bridge_state | `string` | computed |  |  |
| `EgressGatewayBridge` | egress_gateway_bridge | `map` | optional, computed, provider-chosen |  |  |
| `IngressGatewayBridge` | ingress_gateway_bridge | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  | The name of the bridge. |
| `Outputs` |  | `list` | optional, computed, provider-chosen |  | The outputs on this bridge. |
| `PlacementArn` | placement_arn | `string` | required |  | The placement Amazon Resource Number (ARN) of the bridge. |
| `SourceFailoverConfig` | source_failover_config | `map` | optional, computed, provider-chosen |  | The settings for source failover. |
| `Sources` |  | `list` | required |  | The sources on this bridge. |

Supports update: yes

Discovery: supported
