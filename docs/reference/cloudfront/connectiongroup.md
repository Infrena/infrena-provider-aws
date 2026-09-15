# aws.connectiongroup

**CloudFormation type:** `AWS::CloudFront::ConnectionGroup`

The connection group for your distribution tenants. When you first create a distribution tenant and you don't specify a connection group, CloudFront will automatically create a default connection group for you. When you create a new distribution tenant and don't specify a connection group, the default one will be associated with your distribution tenant.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::ConnectionGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnycastIpListId` | anycast_ip_list_id | `string` | optional, computed, provider-chosen | aws.anycastiplist.Id | The ID of the Anycast static IP list. |
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `ETag` | e_tag | `string` | computed |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | Whether the connection group is enabled. |
| `Id` |  | `string` | computed |  |  |
| `Ipv6Enabled` | ipv6_enabled | `boolean` | optional, computed, provider-chosen |  | IPv6 is enabled for the connection group. |
| `IsDefault` | is_default | `boolean` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  | The name of the connection group. |
| `RoutingEndpoint` | routing_endpoint | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
