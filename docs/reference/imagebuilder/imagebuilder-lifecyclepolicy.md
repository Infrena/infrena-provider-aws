# aws.imagebuilder.lifecyclepolicy

**CloudFormation type:** `AWS::ImageBuilder::LifecyclePolicy`

Resource Type definition for AWS::ImageBuilder::LifecyclePolicy

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::LifecyclePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the lifecycle policy. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the lifecycle policy. |
| `ExecutionRole` | execution_role | `string` | required |  | The execution role of the lifecycle policy. |
| `Name` |  | `string` | required, replaces on change |  | The name of the lifecycle policy. |
| `PolicyDetails` | policy_details | `list` | required |  | The policy details of the lifecycle policy. |
| `ResourceSelection` | resource_selection | `map` | required |  | The resource selection for the lifecycle policy. |
| `ResourceType` | resource_type | `string` | required |  | The resource type of the lifecycle policy. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the lifecycle policy. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the lifecycle policy. |

Supports update: yes

Discovery: supported
