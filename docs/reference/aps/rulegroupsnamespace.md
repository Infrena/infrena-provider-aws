# aws.rulegroupsnamespace

**CloudFormation type:** `AWS::APS::RuleGroupsNamespace`

RuleGroupsNamespace schema for cloudformation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::APS::RuleGroupsNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The RuleGroupsNamespace ARN. |
| `Data` |  | `string` | required |  | The RuleGroupsNamespace data. |
| `Name` |  | `string` | required, replaces on change |  | The RuleGroupsNamespace name. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Workspace` |  | `string` | required, replaces on change |  | Required to identify a specific APS Workspace associated with this RuleGroupsNamespace. |

Supports update: yes

Discovery: supported (parent resource required)
