# aws.configurationprofile

**CloudFormation type:** `AWS::AppConfig::ConfigurationProfile`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|ConfigurationProfileId` (AWS::AppConfig::ConfigurationProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.appconfig.application.ApplicationId | The application ID. |
| `ConfigurationProfileId` | configuration_profile_id | `string` | computed |  | The configuration profile ID |
| `DeletionProtectionCheck` | deletion_protection_check | `string` | optional, computed, provider-chosen, write-only |  | On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the configuration profile. |
| `KmsKeyArn` | kms_key_arn | `string` | computed |  | The Amazon Resource Name of the AWS Key Management Service key to encrypt new configuration data versions in the AWS AppConfig hosted configuration store. This attribute is only used for hosted configuration types. To encrypt data managed in other configuration stores, see the documentation for how to specify an AWS KMS key for that particular service. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen |  | The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated. |
| `LocationUri` | location_uri | `string` | required, replaces on change |  | A URI to locate the configuration. You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. |
| `Name` |  | `string` | required |  | A name for the configuration profile. |
| `RetrievalRoleArn` | retrieval_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of an IAM role with permission to access the configuration at the specified LocationUri. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata to assign to the configuration profile. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The type of configurations contained in the profile. When calling this API, enter one of the following values for Type: AWS.AppConfig.FeatureFlags, AWS.Freeform |
| `Validators` |  | `list` | optional, computed, provider-chosen |  | A list of methods for validating the configuration. |

Supports update: yes

Discovery: supported (parent resource required)
