# aws.influxdbcluster

**CloudFormation type:** `AWS::Timestream::InfluxDBCluster`

The AWS::Timestream::InfluxDBCluster resource creates an InfluxDB cluster.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Timestream::InfluxDBCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedStorage` | allocated_storage | `integer` | optional, computed, provider-chosen, replaces on change |  | The allocated storage for the InfluxDB cluster. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is associated with the InfluxDB cluster. |
| `Bucket` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The bucket for the InfluxDB cluster. |
| `DbInstanceType` | db_instance_type | `string` | optional, computed, provider-chosen |  | The compute instance of the InfluxDB cluster. |
| `DbParameterGroupIdentifier` | db_parameter_group_identifier | `string` | optional, computed, provider-chosen |  | The name of an existing InfluxDB parameter group. |
| `DbStorageType` | db_storage_type | `string` | optional, computed, provider-chosen, replaces on change |  | The storage type of the InfluxDB cluster. |
| `DeploymentType` | deployment_type | `string` | optional, computed, provider-chosen, replaces on change |  | Deployment type of the InfluxDB cluster. |
| `Endpoint` |  | `string` | computed |  | The connection endpoint for the InfluxDB cluster. |
| `EngineType` | engine_type | `string` | computed |  | The engine type for the InfluxDB cluster. |
| `FailoverMode` | failover_mode | `string` | optional, computed, provider-chosen |  | Failover mode of the InfluxDB cluster. |
| `Id` |  | `string` | computed |  | The service generated unique identifier for InfluxDB cluster. |
| `InfluxAuthParametersSecretArn` | influx_auth_parameters_secret_arn | `string` | computed |  | The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB cluster. |
| `LogDeliveryConfiguration` | log_delivery_configuration | `map` | optional, computed, provider-chosen |  | Configuration for sending logs to customer account from the InfluxDB cluster. |
| `MaintenanceSchedule` | maintenance_schedule | `map` | optional, computed, provider-chosen |  | The maintenance schedule for the InfluxDB cluster. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name that is associated with the InfluxDB cluster. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | Network type of the InfluxDB cluster. |
| `NextMaintenanceTime` | next_maintenance_time | `string` | computed |  | The timestamp of the next scheduled maintenance event. |
| `Organization` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The organization for the InfluxDB cluster. |
| `Password` |  | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  | The password for the InfluxDB cluster. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port number on which InfluxDB accepts connections. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen, replaces on change |  | Attach a public IP to the customer ENI. |
| `ReaderEndpoint` | reader_endpoint | `string` | computed |  | The reader endpoint for the InfluxDB cluster. |
| `Status` |  | `string` | computed |  | Status of the InfluxDB cluster. |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this DB cluster. |
| `Username` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The username for the InfluxDB cluster. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | A list of Amazon EC2 VPC security groups to associate with this InfluxDB cluster. |
| `VpcSubnetIds` | vpc_subnet_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | A list of EC2 subnet IDs for this InfluxDB cluster. |

Supports update: yes

Discovery: supported
