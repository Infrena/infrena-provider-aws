# aws.accesssource

**CloudFormation type:** `AWS::Route53GlobalResolver::AccessSource`

Resource schema for AWS::Route53GlobalResolver::AccessSource

Region attribute: `region`

**Import ID:** `<region>/AccessSourceId` (AWS::Route53GlobalResolver::AccessSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessSourceId` | access_source_id | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `Cidr` |  | `string` | required |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DnsViewId` | dns_view_id | `string` | required, replaces on change | aws.dnsview.DnsViewId |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Protocol` |  | `string` | required |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
