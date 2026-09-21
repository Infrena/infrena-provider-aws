# aws.context

**CloudFormation type:** `AWS::SageMaker::Context`

Resource type definition for AWS::SageMaker::Context. A context is a lineage tracking entity that represents a logical grouping of other tracking or experiment entities.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::Context)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the context. |
| `ContextName` | context_name | `string` | required, replaces on change |  | The name of the context. Must be unique to your account in an AWS Region. |
| `ContextType` | context_type | `string` | required, replaces on change |  | The context type. |
| `CreationTime` | creation_time | `string` | computed |  | When the context was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the context. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | When the context was last modified. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | A list of properties to add to the context. |
| `Source` |  | `map` | required, replaces on change |  | The source type, ID, and URI. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the context. |

Supports update: yes

Discovery: supported
