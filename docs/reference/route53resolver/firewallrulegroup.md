# aws.firewallrulegroup

**CloudFormation type:** `AWS::Route53Resolver::FirewallRuleGroup`

Resource schema for AWS::Route53Resolver::FirewallRuleGroup.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::FirewallRuleGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn |
| `CreationTime` | creation_time | `string` | computed |  | Rfc3339TimeString |
| `CreatorRequestId` | creator_request_id | `string` | computed |  | The id of the creator request. |
| `FirewallRules` | firewall_rules | `list` | optional, computed, provider-chosen |  | FirewallRules |
| `Id` |  | `string` | computed |  | ResourceId |
| `ModificationTime` | modification_time | `string` | computed |  | Rfc3339TimeString |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | FirewallRuleGroupName |
| `OwnerId` | owner_id | `string` | computed |  | AccountId |
| `RuleCount` | rule_count | `integer` | computed |  | Count |
| `ShareStatus` | share_status | `string` | computed |  | ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME. |
| `Status` |  | `string` | computed |  | ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED. |
| `StatusMessage` | status_message | `string` | computed |  | FirewallRuleGroupStatus |
| `Tags` |  | `map` | tags map |  | Tags |

Supports update: yes

Discovery: supported
