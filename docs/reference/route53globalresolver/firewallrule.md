# aws.firewallrule

**CloudFormation type:** `AWS::Route53GlobalResolver::FirewallRule`

Resource schema for AWS::Route53GlobalResolver::FirewallRule

Region attribute: `region`

**Import ID:** `<region>/FirewallRuleId` (AWS::Route53GlobalResolver::FirewallRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | required |  |  |
| `BlockOverrideDnsType` | block_override_dns_type | `string` | optional, computed, provider-chosen |  |  |
| `BlockOverrideDomain` | block_override_domain | `string` | optional, computed, provider-chosen |  |  |
| `BlockOverrideTtl` | block_override_ttl | `integer` | optional, computed, provider-chosen |  |  |
| `BlockResponse` | block_response | `string` | optional, computed, provider-chosen |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ConfidenceThreshold` | confidence_threshold | `string` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DnsAdvancedProtection` | dns_advanced_protection | `string` | optional, computed, provider-chosen |  |  |
| `DnsViewId` | dns_view_id | `string` | required, replaces on change | aws.dnsview.DnsViewId |  |
| `FirewallDomainListId` | firewall_domain_list_id | `string` | optional, computed, provider-chosen, replaces on change | aws.route53globalresolver.firewalldomainlist.FirewallDomainListId |  |
| `FirewallRuleId` | firewall_rule_id | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Priority` |  | `integer` | optional, computed, provider-chosen |  |  |
| `QType` | q_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `QueryType` | query_type | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
