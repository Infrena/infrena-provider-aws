# aws.bridgeoutput

**CloudFormation type:** `AWS::MediaConnect::BridgeOutput`

Resource schema for AWS::MediaConnect::BridgeOutput

Region attribute: `region`

**Import ID:** `<region>/BridgeArn|Name` (AWS::MediaConnect::BridgeOutput)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BridgeArn` | bridge_arn | `string` | required, replaces on change | aws.bridge.BridgeArn | The Amazon Resource Number (ARN) of the bridge. |
| `Name` |  | `string` | required, replaces on change |  | The network output name. |
| `NetworkOutput` | network_output | `map` | required |  | The output of the bridge. A network output is delivered to your premises. |

Supports update: yes

Discovery: not supported
