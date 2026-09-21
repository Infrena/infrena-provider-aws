# aws.sequencestore

**CloudFormation type:** `AWS::Omics::SequenceStore`

Resource Type definition for AWS::Omics::SequenceStore

Region attribute: `region`

**Import ID:** `<region>/SequenceStoreId` (AWS::Omics::SequenceStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessLogLocation` | access_log_location | `string` | optional, computed, provider-chosen |  | Location of the access logs. |
| `Arn` |  | `string` | computed |  | The store's ARN. |
| `CreationTime` | creation_time | `string` | computed |  | When the store was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the store. |
| `ETagAlgorithmFamily` | e_tag_algorithm_family | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `FallbackLocation` | fallback_location | `string` | optional, computed, provider-chosen |  | An S3 location that is used to store files that have failed a direct upload. |
| `Name` |  | `string` | required |  | A name for the store. |
| `PropagatedSetLevelTags` | propagated_set_level_tags | `list` | optional, computed, provider-chosen |  | The tags keys to propagate to the S3 objects associated with read sets in the sequence store. |
| `S3AccessPointArn` | s3_access_point_arn | `string` | computed |  | This is ARN of the access point associated with the S3 bucket storing read sets. |
| `S3AccessPolicy` | s3_access_policy | `map` | optional, computed, provider-chosen |  | The resource policy that controls S3 access on the store |
| `S3Uri` | s3_uri | `string` | computed |  | The S3 URI of the sequence store. |
| `SequenceStoreId` | sequence_store_id | `string` | computed |  |  |
| `SseConfig` | sse_config | `map` | optional, computed, provider-chosen, replaces on change |  | Server-side encryption (SSE) settings for a store. |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  | The status message of the sequence store. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `UpdateTime` | update_time | `string` | computed |  | The last-updated time of the sequence store. |

Supports update: yes

Discovery: supported
