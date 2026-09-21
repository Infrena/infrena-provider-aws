# aws.directconnect.connection

**CloudFormation type:** `AWS::DirectConnect::Connection`

Resource Type definition for AWS::DirectConnect::Connection

Region attribute: `region`

**Import ID:** `<region>/ConnectionArn` (AWS::DirectConnect::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Bandwidth` |  | `string` | required, replaces on change |  | The bandwidth of the connection. |
| `ConnectionArn` | connection_arn | `string` | computed |  | The ARN of the connection. |
| `ConnectionId` | connection_id | `string` | computed |  | The ID of the connection. |
| `ConnectionName` | connection_name | `string` | required |  | The name of the connection. |
| `ConnectionState` | connection_state | `string` | computed |  | The state of the connection. |
| `LagId` | lag_id | `string` | optional, computed, provider-chosen | aws.lag.LagId | The ID or ARN of the LAG to associate the connection with. |
| `Location` |  | `string` | required, replaces on change |  | The location of the connection. |
| `ProviderName` | provider_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the service provider associated with the requested connection. |
| `RequestMACSec` | request_mac_sec | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Indicates whether you want the connection to support MAC Security (MACsec). |
| `Tags` |  | `map` | tags map |  | The tags associated with the connection. |

Supports update: yes

Discovery: supported
