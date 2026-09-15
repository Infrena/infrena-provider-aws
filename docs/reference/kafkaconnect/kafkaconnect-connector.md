# aws.kafkaconnect.connector

**CloudFormation type:** `AWS::KafkaConnect::Connector`

Resource Type definition for AWS::KafkaConnect::Connector

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::KafkaConnect::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Capacity` |  | `map` | required |  | Information about the capacity allocated to the connector. |
| `ConnectorArn` | connector_arn | `string` | computed |  | Amazon Resource Name for the created Connector. |
| `ConnectorConfiguration` | connector_configuration | `map` | required |  | The configuration for the connector. |
| `ConnectorDescription` | connector_description | `string` | optional, computed, provider-chosen, replaces on change |  | A summary description of the connector. |
| `ConnectorName` | connector_name | `string` | required, replaces on change |  | The name of the connector. |
| `KafkaCluster` | kafka_cluster | `map` | required, replaces on change |  | Details of how to connect to the Kafka cluster. |
| `KafkaClusterClientAuthentication` | kafka_cluster_client_authentication | `map` | required, replaces on change |  | Details of the client authentication used by the Kafka cluster. |
| `KafkaClusterEncryptionInTransit` | kafka_cluster_encryption_in_transit | `map` | required, replaces on change |  | Details of encryption in transit to the Kafka cluster. |
| `KafkaConnectVersion` | kafka_connect_version | `string` | required, replaces on change |  | The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins. |
| `LogDelivery` | log_delivery | `map` | optional, computed, provider-chosen, replaces on change |  | Details of what logs are delivered and where they are delivered. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The network type of the Connector. |
| `Plugins` |  | `list` | required, replaces on change |  | List of plugins to use with the connector. |
| `ServiceExecutionRoleArn` | service_execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource |
| `WorkerConfiguration` | worker_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the worker configuration to use with the connector. |

Supports update: yes

Discovery: supported
