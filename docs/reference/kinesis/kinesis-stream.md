# aws.kinesis.stream

**CloudFormation type:** `AWS::Kinesis::Stream`

Resource Type definition for AWS::Kinesis::Stream

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Kinesis::Stream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon resource name (ARN) of the Kinesis stream |
| `DesiredShardLevelMetrics` | desired_shard_level_metrics | `list` | optional, computed, provider-chosen |  | The final list of shard-level metrics |
| `MaxRecordSizeInKiB` | max_record_size_in_ki_b | `integer` | optional, computed, provider-chosen |  | Maximum size of a data record in KiB allowed to be put into Kinesis stream. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Kinesis stream. |
| `RetentionPeriodHours` | retention_period_hours | `integer` | optional, computed, provider-chosen |  | The number of hours for the data records that are stored in shards to remain accessible. |
| `ShardCount` | shard_count | `integer` | optional, computed, provider-chosen |  | The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed. |
| `StreamEncryption` | stream_encryption | `map` | optional, computed, provider-chosen |  | When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream. Removing this property from your stack template and updating your stack disables encryption. |
| `StreamModeDetails` | stream_mode_details | `map` | optional, computed, provider-chosen |  | When specified, enables or updates the mode of stream. Default is PROVISIONED. |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream. |
| `WarmThroughputMiBps` | warm_throughput_mi_bps | `integer` | optional, computed, provider-chosen, write-only |  | Target warm throughput in MiB/s for the stream. This property can ONLY be set when StreamMode is ON_DEMAND. |
| `WarmThroughputObject` | warm_throughput_object | `map` | computed |  | Warm throughput configuration details for the stream. Only present for ON_DEMAND streams. |

Supports update: yes

Discovery: supported
