# aws.statemachine

**CloudFormation type:** `AWS::StepFunctions::StateMachine`

Resource schema for StateMachine

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::StepFunctions::StateMachine)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Definition` |  | `map` | optional, computed, provider-chosen, write-only |  |  |
| `DefinitionS3Location` | definition_s3_location | `map` | optional, computed, provider-chosen, write-only |  |  |
| `DefinitionString` | definition_string | `string` | optional, computed, provider-chosen |  |  |
| `DefinitionSubstitutions` | definition_substitutions | `map` | optional, computed, provider-chosen, write-only |  |  |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen |  |  |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | computed |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `StateMachineName` | state_machine_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StateMachineRevisionId` | state_machine_revision_id | `string` | computed |  |  |
| `StateMachineType` | state_machine_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TracingConfiguration` | tracing_configuration | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
