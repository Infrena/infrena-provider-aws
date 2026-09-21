# aws.msk.cluster

**CloudFormation type:** `AWS::MSK::Cluster`

Resource Type definition for AWS::MSK::Cluster

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MSK::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BrokerNodeGroupInfo` | broker_node_group_info | `map` | required |  |  |
| `ClientAuthentication` | client_authentication | `map` | optional, computed, provider-chosen |  |  |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  |  |
| `ConfigurationInfo` | configuration_info | `map` | optional, computed, provider-chosen |  |  |
| `CurrentVersion` | current_version | `string` | computed |  | The current version of the MSK cluster |
| `EncryptionInfo` | encryption_info | `map` | optional, computed, provider-chosen |  |  |
| `EnhancedMonitoring` | enhanced_monitoring | `string` | optional, computed, provider-chosen |  |  |
| `KafkaVersion` | kafka_version | `string` | required |  |  |
| `LoggingInfo` | logging_info | `map` | optional, computed, provider-chosen |  |  |
| `NumberOfBrokerNodes` | number_of_broker_nodes | `integer` | required |  |  |
| `OpenMonitoring` | open_monitoring | `map` | optional, computed, provider-chosen |  |  |
| `Rebalancing` |  | `map` | optional, computed, provider-chosen |  |  |
| `StorageMode` | storage_mode | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `ZookeeperAccess` | zookeeper_access | `map` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
