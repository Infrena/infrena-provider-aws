# aws.route53globalresolver.firewalldomainlist

**CloudFormation type:** `AWS::Route53GlobalResolver::FirewallDomainList`

Resource schema for AWS::Route53GlobalResolver::FirewallDomainList

Region attribute: `region`

**Import ID:** `<region>/FirewallDomainListId` (AWS::Route53GlobalResolver::FirewallDomainList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainCount` | domain_count | `integer` | computed |  |  |
| `DomainFileUrl` | domain_file_url | `string` | optional, computed, provider-chosen, write-only |  | S3 URL to import domains from. |
| `Domains` |  | `list` | optional, computed, provider-chosen, write-only |  | An inline list of domains to use for this domain list. |
| `FirewallDomainListId` | firewall_domain_list_id | `string` | computed |  |  |
| `GlobalResolverId` | global_resolver_id | `string` | required, replaces on change | aws.globalresolver.GlobalResolverId |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
