# aws.msk.topic

**CloudFormation type:** `AWS::MSK::Topic`

Resource Type definition for AWS::MSK::Topic

Region attribute: `region`

**Import ID:** `<region>/TopicArn` (AWS::MSK::Topic)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | required, replaces on change | aws.msk.cluster.Arn | The Amazon Resource Name (ARN) of the MSK cluster |
| `Configs` |  | `string` | optional, computed, provider-chosen, write-only |  | Base64 encoded configuration properties of the topic |
| `PartitionCount` | partition_count | `integer` | required |  | The number of partitions for the topic |
| `ReplicationFactor` | replication_factor | `integer` | required, replaces on change |  | The replication factor for the topic |
| `TopicArn` | topic_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the topic |
| `TopicName` | topic_name | `string` | required, replaces on change |  | The name of the topic |

Supports update: yes

Discovery: supported
