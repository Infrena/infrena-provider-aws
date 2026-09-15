# aws.partnerapp

**CloudFormation type:** `AWS::SageMaker::PartnerApp`

Resource Type definition for AWS::SageMaker::PartnerApp

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::PartnerApp)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppVersion` | app_version | `string` | optional, computed, provider-chosen |  | The version of the PartnerApp. |
| `ApplicationConfig` | application_config | `map` | optional, computed, provider-chosen |  | A collection of configuration settings for the PartnerApp. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the created PartnerApp. |
| `AuthType` | auth_type | `string` | required, replaces on change |  | The Auth type of PartnerApp. |
| `BaseUrl` | base_url | `string` | computed |  | The AppServerUrl based on app and account-info. |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, write-only |  | The client token for the PartnerApp. |
| `CurrentVersionEolDate` | current_version_eol_date | `string` | computed |  | The end-of-life date for the current version of the PartnerApp. |
| `EnableAutoMinorVersionUpgrade` | enable_auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  | Enables automatic minor version upgrades for the PartnerApp. |
| `EnableIamSessionBasedIdentity` | enable_iam_session_based_identity | `boolean` | optional, computed, provider-chosen |  | Enables IAM Session based Identity for PartnerApp. |
| `ExecutionRoleArn` | execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The execution role for the user. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS KMS customer managed key used to encrypt the data associated with the PartnerApp. |
| `MaintenanceConfig` | maintenance_config | `map` | optional, computed, provider-chosen |  | A collection of settings that specify the maintenance schedule for the PartnerApp. |
| `Name` |  | `string` | required, replaces on change |  | A name for the PartnerApp. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to apply to the PartnerApp. |
| `Tier` |  | `string` | required |  | The tier of the PartnerApp. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of PartnerApp. |

Supports update: yes

Discovery: supported
