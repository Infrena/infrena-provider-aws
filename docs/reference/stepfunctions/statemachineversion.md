# aws.statemachineversion

**CloudFormation type:** `AWS::StepFunctions::StateMachineVersion`

Resource schema for StateMachineVersion

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::StepFunctions::StateMachineVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StateMachineArn` | state_machine_arn | `string` | required, replaces on change | aws.statemachine.Arn |  |
| `StateMachineRevisionId` | state_machine_revision_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported (parent resource required)
