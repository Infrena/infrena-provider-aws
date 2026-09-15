# aws.modelpackagegroup

**CloudFormation type:** `AWS::SageMaker::ModelPackageGroup`

Resource Type definition for AWS::SageMaker::ModelPackageGroup

Region attribute: `region`

**Import ID:** `<region>/ModelPackageGroupArn` (AWS::SageMaker::ModelPackageGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time at which the model package group was created. |
| `ModelPackageGroupArn` | model_package_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the model package group. |
| `ModelPackageGroupDescription` | model_package_group_description | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the model package group. |
| `ModelPackageGroupName` | model_package_group_name | `string` | required, replaces on change |  | The name of the model package group. |
| `ModelPackageGroupPolicy` | model_package_group_policy | `string` | optional, computed, provider-chosen |  |  |
| `ModelPackageGroupStatus` | model_package_group_status | `string` | computed |  | The status of a modelpackage group job. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
