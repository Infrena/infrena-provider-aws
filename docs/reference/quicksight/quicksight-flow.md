# aws.quicksight.flow

**CloudFormation type:** `AWS::QuickSight::Flow`

Definition of AWS::QuickSight::Flow Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::QuickSight::Flow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `FlowDefinition` | flow_definition | `string` | required |  |  |
| `FlowId` | flow_id | `string` | computed |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen, write-only |  |  |
| `PublishState` | publish_state | `string` | computed |  |  |
| `StepAliases` | step_aliases | `list` | computed |  |  |

Supports update: yes

Discovery: supported
