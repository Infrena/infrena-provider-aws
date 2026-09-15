# aws.statemachinealias

**CloudFormation type:** `AWS::StepFunctions::StateMachineAlias`

Resource schema for StateMachineAlias

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::StepFunctions::StateMachineAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the alias. |
| `DeploymentPreference` | deployment_preference | `map` | optional, computed, provider-chosen, write-only |  | The settings to enable gradual state machine deployments. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description of the alias. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The alias name. |
| `RoutingConfiguration` | routing_configuration | `list` | optional, computed, provider-chosen |  | The routing configuration of the alias. One or two versions can be mapped to an alias to split StartExecution requests of the same state machine. |
| `StateMachineArn` | state_machine_arn | `string` | optional, computed, provider-chosen | aws.statemachine.Arn |  |

Supports update: yes

Discovery: supported (parent resource required)
