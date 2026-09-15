# aws.detective.graph

**CloudFormation type:** `AWS::Detective::Graph`

Resource schema for AWS::Detective::Graph

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Detective::Graph)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Detective graph ARN |
| `AutoEnableMembers` | auto_enable_members | `boolean` | optional, computed, provider-chosen |  | Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
