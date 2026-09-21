# aws.studiolifecycleconfig

**CloudFormation type:** `AWS::SageMaker::StudioLifecycleConfig`

Resource Type definition for AWS::SageMaker::StudioLifecycleConfig

Region attribute: `region`

**Import ID:** `<region>/StudioLifecycleConfigName` (AWS::SageMaker::StudioLifecycleConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `StudioLifecycleConfigAppType` | studio_lifecycle_config_app_type | `string` | required, replaces on change |  | The App type that the Lifecycle Configuration is attached to. |
| `StudioLifecycleConfigArn` | studio_lifecycle_config_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Lifecycle Configuration. |
| `StudioLifecycleConfigContent` | studio_lifecycle_config_content | `string` | required, replaces on change |  | The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded. |
| `StudioLifecycleConfigName` | studio_lifecycle_config_name | `string` | required, replaces on change |  | The name of the Amazon SageMaker Studio Lifecycle Configuration. |
| `Tags` |  | `map` | replaces on change, tags map |  | Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API. |

Supports update: no

Discovery: supported
