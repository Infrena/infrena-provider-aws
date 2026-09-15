# aws.resourcedatasync

**CloudFormation type:** `AWS::SSM::ResourceDataSync`

Resource Type definition for AWS::SSM::ResourceDataSync

Region attribute: `region`

**Import ID:** `<region>/SyncName` (AWS::SSM::ResourceDataSync)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BucketName` | bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `BucketPrefix` | bucket_prefix | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `BucketRegion` | bucket_region | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `KMSKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `S3Destination` | s3_destination | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SyncFormat` | sync_format | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SyncName` | sync_name | `string` | required, replaces on change |  |  |
| `SyncSource` | sync_source | `map` | optional, computed, provider-chosen |  |  |
| `SyncType` | sync_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
