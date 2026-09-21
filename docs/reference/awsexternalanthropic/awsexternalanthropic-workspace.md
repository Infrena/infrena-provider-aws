# aws.awsexternalanthropic.workspace

**CloudFormation type:** `AWS::AWSExternalAnthropic::Workspace`

Resource type definition for AWS::AWSExternalAnthropic::Workspace

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AWSExternalAnthropic::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the workspace. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the workspace was created. |
| `DataResidency` | data_residency | `map` | optional, computed, provider-chosen, replaces on change |  | Data residency configuration for the workspace. WorkspaceGeo is immutable after creation. |
| `Id` |  | `string` | computed |  | The unique identifier of the workspace. |
| `Name` |  | `string` | required |  | The name of the workspace. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
