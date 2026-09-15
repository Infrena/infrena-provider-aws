# aws.lag

**CloudFormation type:** `AWS::DirectConnect::Lag`

Resource Type definition for AWS::DirectConnect::Lag

Region attribute: `region`

**Import ID:** `<region>/LagArn` (AWS::DirectConnect::Lag)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionsBandwidth` | connections_bandwidth | `string` | required, replaces on change |  | The bandwidth of the individual physical dedicated connections bundled by the LAG. |
| `LagArn` | lag_arn | `string` | computed |  | The ARN of the LAG. |
| `LagId` | lag_id | `string` | computed |  | The ID of the LAG. |
| `LagName` | lag_name | `string` | required |  | The name of the LAG. |
| `LagState` | lag_state | `string` | computed |  | The state of the LAG. |
| `Location` |  | `string` | required, replaces on change |  | The location for the LAG. |
| `MinimumLinks` | minimum_links | `integer` | optional, computed, provider-chosen |  | The minimum number of physical dedicated connections that must be operational for the LAG itself to be operational. |
| `ProviderName` | provider_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the service provider associated with the requested LAG. |
| `RequestMACSec` | request_mac_sec | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Indicates whether you want the LAG to support MAC Security (MACsec). |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the LAG. |

Supports update: yes

Discovery: supported
