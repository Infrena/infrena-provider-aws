# aws.rbin.rule

**CloudFormation type:** `AWS::Rbin::Rule`

Resource Type definition for AWS::Rbin::Rule

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Rbin::Rule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Rule Arn is unique for each rule. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the retention rule. |
| `ExcludeResourceTags` | exclude_resource_tags | `list` | optional, computed, provider-chosen |  | Information about the exclude resource tags used to identify resources that are excluded by the retention rule. |
| `Identifier` |  | `string` | computed |  | The unique ID of the retention rule. |
| `LockConfiguration` | lock_configuration | `map` | optional, computed, provider-chosen, write-only |  | Information about the retention rule lock configuration. |
| `LockState` | lock_state | `string` | computed |  | The lock state for the retention rule. |
| `ResourceTags` | resource_tags | `list` | optional, computed, provider-chosen |  | Information about the resource tags used to identify resources that are retained by the retention rule. |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | The resource type retained by the retention rule. |
| `RetentionPeriod` | retention_period | `map` | required |  | The retention period of the rule. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The state of the retention rule. Only retention rules that are in the available state retain resources. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Information about the tags assigned to the retention rule. |

Supports update: yes

Discovery: supported (parent resource required)
