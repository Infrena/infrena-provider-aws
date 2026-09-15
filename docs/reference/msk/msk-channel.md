# aws.msk.channel

**CloudFormation type:** `AWS::MSK::Channel`

Resource Type definition for AWS::MSK::Channel

Region attribute: `region`

**Import ID:** `<region>/ChannelArn` (AWS::MSK::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelArn` | channel_arn | `string` | computed |  | The Amazon Resource Name (ARN) that uniquely identifies the channel |
| `ChannelName` | channel_name | `string` | required, replaces on change |  | Name of the channel |
| `ClusterArn` | cluster_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.msk.cluster.Arn | The Amazon Resource Name (ARN) of the cluster |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Encryption configuration |
| `IcebergDestinationConfiguration` | iceberg_destination_configuration | `map` | optional, computed, provider-chosen |  | Iceberg destination configuration |
| `LoggingInfo` | logging_info | `map` | optional, computed, provider-chosen, replaces on change |  | Log configuration details for Channel |
| `S3DestinationConfiguration` | s3_destination_configuration | `map` | optional, computed, provider-chosen |  | S3 destination configuration |
| `StateInfo` | state_info | `map` | computed |  | Includes information about the channel state |
| `Status` |  | `string` | computed |  | Status of a channel resource |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Tags attached to the channel |
| `TopicConfigurationList` | topic_configuration_list | `list` | required, replaces on change |  | Topic configuration |

Supports update: yes

Discovery: supported
