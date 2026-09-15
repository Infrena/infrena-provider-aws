# aws.influxdbinstance

**CloudFormation type:** `AWS::Timestream::InfluxDBInstance`

The AWS::Timestream::InfluxDBInstance resource creates an InfluxDB instance.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Timestream::InfluxDBInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedStorage` | allocated_storage | `integer` | optional, computed, provider-chosen |  | The allocated storage for the InfluxDB instance. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) that is associated with the InfluxDB instance. |
| `AvailabilityZone` | availability_zone | `string` | computed |  | The Availability Zone (AZ) where the InfluxDB instance is created. |
| `Bucket` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The bucket for the InfluxDB instance. |
| `DbInstanceType` | db_instance_type | `string` | optional, computed, provider-chosen |  | The compute instance of the InfluxDB instance. |
| `DbParameterGroupIdentifier` | db_parameter_group_identifier | `string` | optional, computed, provider-chosen |  | The name of an existing InfluxDB parameter group. |
| `DbStorageType` | db_storage_type | `string` | optional, computed, provider-chosen |  | The storage type of the InfluxDB instance. |
| `DeploymentType` | deployment_type | `string` | optional, computed, provider-chosen |  | Deployment type of the InfluxDB Instance. |
| `Endpoint` |  | `string` | computed |  | The connection endpoint for the InfluxDB instance. |
| `Id` |  | `string` | computed |  | The service generated unique identifier for InfluxDB instance. |
| `InfluxAuthParametersSecretArn` | influx_auth_parameters_secret_arn | `string` | computed |  | The Auth parameters secret Amazon Resource name (ARN) that is associated with the InfluxDB instance. |
| `LogDeliveryConfiguration` | log_delivery_configuration | `map` | optional, computed, provider-chosen |  | Configuration for sending logs to customer account from the InfluxDB instance. |
| `MaintenanceSchedule` | maintenance_schedule | `map` | optional, computed, provider-chosen |  | The maintenance schedule for the InfluxDB instance. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name that is associated with the InfluxDB instance. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | Network type of the InfluxDB Instance. |
| `NextMaintenanceTime` | next_maintenance_time | `string` | computed |  | The timestamp of the next scheduled maintenance event. |
| `Organization` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The organization for the InfluxDB instance. |
| `Password` |  | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  | The password for the InfluxDB instance. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port number on which InfluxDB accepts connections. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen, replaces on change |  | Attach a public IP to the customer ENI. |
| `SecondaryAvailabilityZone` | secondary_availability_zone | `string` | computed |  | The Secondary Availability Zone (AZ) where the InfluxDB instance is created, if DeploymentType is set as WITH_MULTIAZ_STANDBY. |
| `Status` |  | `string` | computed |  | Status of the InfluxDB Instance. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this DB instance. |
| `Username` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The username for the InfluxDB instance. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | A list of Amazon EC2 VPC security groups to associate with this InfluxDB instance. |
| `VpcSubnetIds` | vpc_subnet_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | A list of EC2 subnet IDs for this InfluxDB instance. |

Supports update: yes

Discovery: supported
