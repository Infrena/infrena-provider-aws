# aws.bedrockmantle.project

**CloudFormation type:** `AWS::BedrockMantle::Project`

Resource type definition for AWS::BedrockMantle::Project

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BedrockMantle::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the project. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the project was created. |
| `Id` |  | `string` | computed |  | The unique identifier of the project. |
| `Name` |  | `string` | required |  | The name of the project. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
