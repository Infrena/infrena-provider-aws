# aws.s3files.filesystem

**CloudFormation type:** `AWS::S3Files::FileSystem`

Resource Type definition for AWS::S3Files::FileSystem

Region attribute: `region`

**Import ID:** `<region>/FileSystemArn` (AWS::S3Files::FileSystem)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptBucketWarning` | accept_bucket_warning | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Bucket` |  | `string` | required, replaces on change |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `FileSystemArn` | file_system_arn | `string` | computed |  |  |
| `FileSystemId` | file_system_id | `string` | computed |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `OwnerId` | owner_id | `string` | computed |  |  |
| `Prefix` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `SynchronizationConfiguration` | synchronization_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
