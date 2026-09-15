# aws.staticip

**CloudFormation type:** `AWS::Lightsail::StaticIp`

Resource Type definition for AWS::Lightsail::StaticIp

Region attribute: `region`

**Import ID:** `<region>/StaticIpName` (AWS::Lightsail::StaticIp)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachedTo` | attached_to | `string` | optional, computed, provider-chosen |  | The instance where the static IP is attached. |
| `IpAddress` | ip_address | `string` | computed |  | The static IP address. |
| `IsAttached` | is_attached | `boolean` | computed |  | A Boolean value indicating whether the static IP is attached. |
| `StaticIpArn` | static_ip_arn | `string` | computed |  |  |
| `StaticIpName` | static_ip_name | `string` | required, replaces on change |  | The name of the static IP address. |

Supports update: yes

Discovery: supported
