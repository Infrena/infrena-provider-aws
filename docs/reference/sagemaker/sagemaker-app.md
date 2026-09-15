# aws.sagemaker.app

**CloudFormation type:** `AWS::SageMaker::App`

Resource Type definition for AWS::SageMaker::App

Region attribute: `region`

**Import ID:** `<region>/AppName|AppType|DomainId|UserProfileName` (AWS::SageMaker::App)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppArn` | app_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the app. |
| `AppName` | app_name | `string` | required, replaces on change |  | The name of the app. |
| `AppType` | app_type | `string` | required, replaces on change |  | The type of app. |
| `BuiltInLifecycleConfigArn` | built_in_lifecycle_config_arn | `string` | computed |  | The lifecycle configuration that runs before the default lifecycle configuration. |
| `DomainId` | domain_id | `string` | required, replaces on change | aws.sagemaker.domain.DomainId | The domain ID. |
| `RecoveryMode` | recovery_mode | `boolean` | optional, computed, provider-chosen |  | Indicates whether the application is launched in recovery mode. |
| `ResourceSpec` | resource_spec | `map` | optional, computed, provider-chosen, replaces on change |  | The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to apply to the app. |
| `UserProfileName` | user_profile_name | `string` | required, replaces on change |  | The user profile name. |

Supports update: yes

Discovery: supported
