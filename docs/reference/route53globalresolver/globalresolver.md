# aws.globalresolver

**CloudFormation type:** `AWS::Route53GlobalResolver::GlobalResolver`

Resource schema for AWS::Route53GlobalResolver::GlobalResolver

Region attribute: `region`

**Import ID:** `<region>/GlobalResolverId` (AWS::Route53GlobalResolver::GlobalResolver)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DnsName` | dns_name | `string` | computed |  |  |
| `GlobalResolverId` | global_resolver_id | `string` | computed |  |  |
| `IPv4Addresses` | i_pv4_addresses | `list` | computed |  |  |
| `IPv6Addresses` | i_pv6_addresses | `list` | computed |  |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `ObservabilityRegion` | observability_region | `string` | optional, computed, provider-chosen |  |  |
| `Regions` |  | `list` | required |  | The list of regions the Global Resolver exists in. Regions can be added or removed on update; the order of this list is not significant. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
