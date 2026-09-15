# aws.dnsview

**CloudFormation type:** `AWS::Route53GlobalResolver::DnsView`

Resource schema for AWS::Route53GlobalResolver::DnsView

Region attribute: `region`

**Import ID:** `<region>/DnsViewId` (AWS::Route53GlobalResolver::DnsView)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DnsViewId` | dns_view_id | `string` | computed |  |  |
| `DnssecValidation` | dnssec_validation | `string` | optional, computed, provider-chosen |  |  |
| `EdnsClientSubnet` | edns_client_subnet | `string` | optional, computed, provider-chosen |  |  |
| `FirewallRulesFailOpen` | firewall_rules_fail_open | `string` | optional, computed, provider-chosen |  |  |
| `GlobalResolverId` | global_resolver_id | `string` | required, replaces on change | aws.globalresolver.GlobalResolverId |  |
| `Name` |  | `string` | required |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
