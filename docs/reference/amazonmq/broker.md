# aws.broker

**CloudFormation type:** `AWS::AmazonMQ::Broker`

Resource type definition for AWS::AmazonMQ::Broker

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AmazonMQ::Broker)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmqpEndpoints` | amqp_endpoints | `list` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AuthenticationStrategy` | authentication_strategy | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `AutoMinorVersionUpgrade` | auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  |  |
| `BrokerName` | broker_name | `string` | required, replaces on change |  |  |
| `Configuration` |  | `map` | optional, computed, provider-chosen, write-only |  | The intended configuration (ID and revision) to be set when creating or updating. |
| `ConfigurationId` | configuration_id | `string` | computed |  | The ID of the current actual configuration. |
| `ConfigurationRevision` | configuration_revision | `string` | computed |  | The revision of the current actual configuration. |
| `ConsoleURLs` | console_ur_ls | `list` | computed |  |  |
| `DataReplicationMode` | data_replication_mode | `string` | optional, computed, provider-chosen |  |  |
| `DataReplicationPrimaryBrokerArn` | data_replication_primary_broker_arn | `string` | optional, computed, provider-chosen, write-only | aws.broker.Arn | The ARN of the primary broker that is used to replicate data from in a data replication pair when creating a replica. |
| `DeploymentMode` | deployment_mode | `string` | required, replaces on change |  |  |
| `EncryptionOptions` | encryption_options | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `EngineType` | engine_type | `string` | required, replaces on change |  |  |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen, write-only |  | The version specified to use. See also EngineVersionCurrent. |
| `EngineVersionCurrent` | engine_version_current | `string` | computed |  | The version in use. This may have more precision than the specified EngineVersion. |
| `HostInstanceType` | host_instance_type | `string` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `IpAddresses` | ip_addresses | `list` | computed |  |  |
| `LdapServerMetadata` | ldap_server_metadata | `map` | optional, computed, provider-chosen |  |  |
| `Logs` |  | `map` | optional, computed, provider-chosen |  |  |
| `MaintenanceWindowStartTime` | maintenance_window_start_time | `map` | optional, computed, provider-chosen |  |  |
| `MqttEndpoints` | mqtt_endpoints | `list` | computed |  |  |
| `OpenWireEndpoints` | open_wire_endpoints | `list` | computed |  |  |
| `PubliclyAccessible` | publicly_accessible | `boolean` | required, replaces on change |  |  |
| `ResourceShareArns` | resource_share_arns | `list` | optional, computed, provider-chosen | aws.resourceshare.Arn | The ARNs of the resource shares to be associated with the broker. |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen |  |  |
| `StompEndpoints` | stomp_endpoints | `list` | computed |  |  |
| `StorageSize` | storage_size | `integer` | optional, computed, provider-chosen |  | The broker's storage size in GB. |
| `StorageType` | storage_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Users` |  | `list` | optional, computed, provider-chosen, write-only |  | Users to configure on the broker. |
| `WssEndpoints` | wss_endpoints | `list` | computed |  |  |

Supports update: yes

Discovery: supported
