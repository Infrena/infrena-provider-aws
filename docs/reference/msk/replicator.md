# aws.replicator

**CloudFormation type:** `AWS::MSK::Replicator`

Resource Type definition for AWS::MSK::Replicator

Region attribute: `region`

**Import ID:** `<region>/ReplicatorArn` (AWS::MSK::Replicator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CurrentVersion` | current_version | `string` | computed |  | The current version of the MSK replicator. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A summary description of the replicator. |
| `KafkaClusters` | kafka_clusters | `list` | required, replaces on change |  | Specifies a list of Kafka clusters which are targets of the replicator. |
| `LogDelivery` | log_delivery | `map` | optional, computed, provider-chosen |  | Configuration for log delivery for the replicator. |
| `ReplicationInfoList` | replication_info_list | `list` | required |  | A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow. |
| `ReplicatorArn` | replicator_arn | `string` | computed |  | Amazon Resource Name for the created replicator. |
| `ReplicatorName` | replicator_name | `string` | required, replaces on change |  | The name of the replicator. |
| `ServiceExecutionRoleArn` | service_execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources. |
| `Tags` |  | `map` | tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
