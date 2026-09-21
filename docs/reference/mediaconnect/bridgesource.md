# aws.bridgesource

**CloudFormation type:** `AWS::MediaConnect::BridgeSource`

Resource schema for AWS::MediaConnect::BridgeSource

Region attribute: `region`

**Import ID:** `<region>/BridgeArn|Name` (AWS::MediaConnect::BridgeSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BridgeArn` | bridge_arn | `string` | required, replaces on change | aws.bridge.BridgeArn | The Amazon Resource Number (ARN) of the bridge. |
| `FlowSource` | flow_source | `map` | optional, computed, provider-chosen |  | The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow. |
| `Name` |  | `string` | required, replaces on change |  | The name of the source. |
| `NetworkSource` | network_source | `map` | optional, computed, provider-chosen |  | The source of the bridge. A network source originates at your premises. |

Supports update: yes

Discovery: not supported
