# aws.replicationconfig

**CloudFormation type:** `AWS::DMS::ReplicationConfig`

A replication configuration that you later provide to configure and start a AWS DMS Serverless replication

Region attribute: `region`

**Import ID:** `<region>/ReplicationConfigArn` (AWS::DMS::ReplicationConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputeConfig` | compute_config | `map` | required |  | Configuration parameters for provisioning a AWS DMS Serverless replication |
| `ReplicationConfigArn` | replication_config_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Replication Config |
| `ReplicationConfigIdentifier` | replication_config_identifier | `string` | required |  | A unique identifier of replication configuration |
| `ReplicationSettings` | replication_settings | `map` | optional, computed, provider-chosen |  | JSON settings for Servereless replications that are provisioned using this replication configuration |
| `ReplicationType` | replication_type | `string` | required |  | The type of AWS DMS Serverless replication to provision using this replication configuration |
| `ResourceIdentifier` | resource_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource |
| `SourceEndpointArn` | source_endpoint_arn | `string` | required | aws.dms.endpoint.EndpointArn | The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration |
| `SupplementalSettings` | supplemental_settings | `map` | optional, computed, provider-chosen |  | JSON settings for specifying supplemental data |
| `TableMappings` | table_mappings | `map` | required |  | JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p> |
| `TargetEndpointArn` | target_endpoint_arn | `string` | required | aws.dms.endpoint.EndpointArn | The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration |

Supports update: yes

Discovery: supported
