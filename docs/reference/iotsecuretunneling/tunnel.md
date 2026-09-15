# aws.tunnel

**CloudFormation type:** `AWS::IoTSecureTunneling::Tunnel`

A connection between a source computer and a destination device using AWS IoT Secure Tunneling.

Region attribute: `region`

**Import ID:** `<region>/TunnelArn` (AWS::IoTSecureTunneling::Tunnel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A short text description of the tunnel. |
| `DestinationConfig` | destination_config | `map` | optional, computed, provider-chosen, replaces on change |  | The destination configuration. |
| `Status` |  | `string` | computed |  | The status of the tunnel. Valid values are OPEN and CLOSED. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tag metadata. |
| `TimeoutConfig` | timeout_config | `map` | optional, computed, provider-chosen, replaces on change |  | Tunnel timeout configuration. |
| `TunnelArn` | tunnel_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the tunnel. |
| `TunnelId` | tunnel_id | `string` | computed |  | A unique alpha-numeric tunnel ID. |

Supports update: yes

Discovery: supported
