# aws.action

**CloudFormation type:** `AWS::SageMaker::Action`

Resource type definition for AWS::SageMaker::Action. An action is a lineage tracking entity that represents an action or activity, such as a model deployment or an HPO job.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::Action)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionName` | action_name | `string` | required, replaces on change |  | The name of the action. Must be unique to your account in an AWS Region. |
| `ActionType` | action_type | `string` | required, replaces on change |  | The action type. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the action. |
| `CreationTime` | creation_time | `string` | computed |  | When the action was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the action. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | When the action was last modified. |
| `MetadataProperties` | metadata_properties | `map` | optional, computed, provider-chosen, replaces on change |  | Metadata properties of the tracking entity, trial, or trial component. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | A list of properties to add to the action. |
| `Source` |  | `map` | required, replaces on change |  | The source type, ID, and URI. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the action. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the action. |

Supports update: yes

Discovery: supported
