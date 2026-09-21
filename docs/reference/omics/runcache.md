# aws.runcache

**CloudFormation type:** `AWS::Omics::RunCache`

Definition of AWS::Omics::RunCache Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Omics::RunCache)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The run cache ARN. |
| `CacheBehavior` | cache_behavior | `string` | optional, computed, provider-chosen |  | The default cache behavior for runs using this cache. |
| `CacheBucketOwnerId` | cache_bucket_owner_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS account ID of the expected owner of the S3 bucket for the run cache. |
| `CacheS3Location` | cache_s3_location | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The S3 location for storing the cached task outputs. |
| `CreationTime` | creation_time | `string` | computed |  | Creation time of the run cache (an ISO 8601 formatted string). |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the run cache. |
| `Id` |  | `string` | computed |  | The run cache ID. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A name for the run cache. |
| `Status` |  | `string` | computed |  | The run cache status. |
| `Tags` |  | `map` | tags map |  | Tags for the run cache. |

Supports update: yes

Discovery: supported
