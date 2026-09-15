# aws.rotationschedule

**CloudFormation type:** `AWS::SecretsManager::RotationSchedule`

Resource Type definition for AWS::SecretsManager::RotationSchedule

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SecretsManager::RotationSchedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ExternalSecretRotationMetadata` | external_secret_rotation_metadata | `list` | optional, computed, provider-chosen |  | The list of metadata needed to successfully rotate a managed external secret. |
| `ExternalSecretRotationRoleArn` | external_secret_rotation_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the IAM role that is used by Secrets Manager to rotate a managed external secret. |
| `HostedRotationLambda` | hosted_rotation_lambda | `map` | optional, computed, provider-chosen, write-only |  | Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead. |
| `Id` |  | `string` | computed |  | The ARN of the secret. |
| `RotateImmediatelyOnUpdate` | rotate_immediately_on_update | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window. |
| `RotationLambdaARN` | rotation_lambda_arn | `string` | optional, computed, provider-chosen |  | The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function. |
| `RotationRules` | rotation_rules | `map` | optional, computed, provider-chosen |  | A structure that defines the rotation configuration for this secret. |
| `SecretId` | secret_id | `string` | required, replaces on change | aws.secret.Id | The ARN or name of the secret to rotate. |

Supports update: yes

Discovery: supported
