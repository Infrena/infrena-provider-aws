# aws.firewallrulegroupassociation

**CloudFormation type:** `AWS::Route53Resolver::FirewallRuleGroupAssociation`

Resource schema for AWS::Route53Resolver::FirewallRuleGroupAssociation.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::FirewallRuleGroupAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn |
| `CreationTime` | creation_time | `string` | computed |  | Rfc3339TimeString |
| `CreatorRequestId` | creator_request_id | `string` | computed |  | The id of the creator request. |
| `FirewallRuleGroupId` | firewall_rule_group_id | `string` | required, replaces on change | aws.firewallrulegroup.Id | FirewallRuleGroupId |
| `Id` |  | `string` | computed |  | Id |
| `ManagedOwnerName` | managed_owner_name | `string` | computed |  | ServicePrincipal |
| `ModificationTime` | modification_time | `string` | computed |  | Rfc3339TimeString |
| `MutationProtection` | mutation_protection | `string` | optional, computed, provider-chosen |  | MutationProtectionStatus |
| `Name` |  | `string` | optional, computed, provider-chosen |  | FirewallRuleGroupAssociationName |
| `Priority` |  | `integer` | required |  | Priority |
| `Status` |  | `string` | computed |  | ResolverFirewallRuleGroupAssociation, possible values are COMPLETE, DELETING, UPDATING, and INACTIVE_OWNER_ACCOUNT_CLOSED. |
| `StatusMessage` | status_message | `string` | computed |  | FirewallDomainListAssociationStatus |
| `Tags` |  | `map` | tags map |  | Tags |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | VpcId |

Supports update: yes

Discovery: supported
