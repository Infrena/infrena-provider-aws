# aws.replicationtask

**CloudFormation type:** `AWS::DMS::ReplicationTask`

Resource Type definition for AWS::DMS::ReplicationTask

Region attribute: `region`

**Import ID:** `<region>/ReplicationTaskArn` (AWS::DMS::ReplicationTask)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CdcStartPosition` | cdc_start_position | `string` | optional, computed, provider-chosen |  | Indicates when you want a change data capture (CDC) operation to start. Use either CdcStartPosition or CdcStartTime to specify when you want a CDC operation to start. Specifying both values results in an error. |
| `CdcStartTime` | cdc_start_time | `float` | optional, computed, provider-chosen, write-only |  | Indicates the start time for a change data capture (CDC) operation. Use either CdcStartTime or CdcStartPosition to specify when you want a CDC operation to start. Specifying both values results in an error. |
| `CdcStopPosition` | cdc_stop_position | `string` | optional, computed, provider-chosen |  | Indicates when you want a change data capture (CDC) operation to stop. The value can be either server time or commit time. |
| `MigrationType` | migration_type | `string` | required, replaces on change |  | The migration type. |
| `ReplicationInstanceArn` | replication_instance_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of a replication instance. |
| `ReplicationTaskArn` | replication_task_arn | `string` | computed |  | The ARN of the ReplicationTask. Also serves the purpise of Primary Identifier. |
| `ReplicationTaskIdentifier` | replication_task_identifier | `string` | optional, computed, provider-chosen |  | An identifier for the replication task. |
| `ReplicationTaskSettings` | replication_task_settings | `string` | optional, computed, provider-chosen |  | Overall settings for the task, in JSON format |
| `ResourceIdentifier` | resource_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | A friendly name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object. |
| `SourceEndpointArn` | source_endpoint_arn | `string` | required, replaces on change | aws.dms.endpoint.EndpointArn | An Amazon Resource Name (ARN) that uniquely identifies the source endpoint. |
| `TableMappings` | table_mappings | `string` | required |  | The table mappings for the task, in JSON format. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TargetEndpointArn` | target_endpoint_arn | `string` | required, replaces on change | aws.dms.endpoint.EndpointArn | An Amazon Resource Name (ARN) that uniquely identifies the target endpoint. |
| `TaskData` | task_data | `string` | optional, computed, provider-chosen |  | Supplemental information that the task requires to migrate the data for certain source and target endpoints. |

Supports update: yes

Discovery: supported
