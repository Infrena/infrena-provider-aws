# aws.imagebuilder.workflow

**CloudFormation type:** `AWS::ImageBuilder::Workflow`

Resource schema for AWS::ImageBuilder::Workflow

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::Workflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the workflow. |
| `ChangeDescription` | change_description | `string` | optional, computed, provider-chosen, replaces on change |  | The change description of the workflow. |
| `Data` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The data of the workflow. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the workflow. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The KMS key identifier used to encrypt the workflow. |
| `LatestVersion` | latest_version | `map` | computed |  | The latest version references of the workflow. |
| `Name` |  | `string` | required, replaces on change |  | The name of the workflow. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the workflow. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the workflow denotes whether the workflow is used to build, test, or distribute. |
| `Uri` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The uri of the workflow. |
| `Version` |  | `string` | required, replaces on change |  | The version of the workflow. |

Supports update: yes

Discovery: supported
