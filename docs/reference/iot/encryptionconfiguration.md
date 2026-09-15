# aws.encryptionconfiguration

**CloudFormation type:** `AWS::IoT::EncryptionConfiguration`

Resource Type definition for AWS::IoT::EncryptionConfiguration

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::IoT::EncryptionConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  |  |
| `ConfigurationDetails` | configuration_details | `map` | computed |  |  |
| `EncryptionType` | encryption_type | `string` | required |  |  |
| `KmsAccessRoleArn` | kms_access_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedDate` | last_modified_date | `string` | computed |  |  |

Supports update: yes

Discovery: supported
