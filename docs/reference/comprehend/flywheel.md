# aws.flywheel

**CloudFormation type:** `AWS::Comprehend::Flywheel`

The AWS::Comprehend::Flywheel resource creates an Amazon Comprehend Flywheel that enables customer to train their model.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Comprehend::Flywheel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActiveModelArn` | active_model_arn | `string` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `DataAccessRoleArn` | data_access_role_arn | `string` | required | aws.role.Arn |  |
| `DataLakeS3Uri` | data_lake_s3_uri | `string` | required, replaces on change |  |  |
| `DataSecurityConfig` | data_security_config | `map` | optional, computed, provider-chosen |  |  |
| `FlywheelName` | flywheel_name | `string` | required, replaces on change |  |  |
| `ModelType` | model_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TaskConfig` | task_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
