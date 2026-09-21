# aws.connect.workspace

**CloudFormation type:** `AWS::Connect::Workspace`

Resource Type definition for AWS::Connect::Workspace

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Connect::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) for the workspace. |
| `Associations` |  | `list` | optional, computed, provider-chosen |  | The resource ARNs associated with the workspace |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the workspace |
| `Id` |  | `string` | computed |  | The identifier of the workspace. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Media` |  | `list` | optional, computed, provider-chosen |  | The media items for the workspace |
| `Name` |  | `string` | required |  | The name of the workspace. |
| `Pages` |  | `list` | optional, computed, provider-chosen |  | The pages associated with the workspace |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Theme` |  | `map` | optional, computed, provider-chosen |  | The theme configuration for the Connect workspace |
| `Title` |  | `string` | optional, computed, provider-chosen |  | The title of the workspace |
| `Visibility` |  | `string` | optional, computed, provider-chosen |  | The visibility of the Connect workspace |

Supports update: yes

Discovery: supported
