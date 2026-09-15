# aws.firewallpolicy

**CloudFormation type:** `AWS::NetworkFirewall::FirewallPolicy`

Resource type definition for AWS::NetworkFirewall::FirewallPolicy

Region attribute: `region`

**Import ID:** `<region>/FirewallPolicyArn` (AWS::NetworkFirewall::FirewallPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `FirewallPolicy` | firewall_policy | `map` | required |  |  |
| `FirewallPolicyArn` | firewall_policy_arn | `string` | computed |  | A resource ARN. |
| `FirewallPolicyId` | firewall_policy_id | `string` | computed |  |  |
| `FirewallPolicyName` | firewall_policy_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
