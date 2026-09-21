# aws.m2.environment

**CloudFormation type:** `AWS::M2::Environment`

Represents a runtime environment that can run migrated mainframe applications.

Region attribute: `region`

**Import ID:** `<region>/EnvironmentArn` (AWS::M2::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the environment. |
| `EngineType` | engine_type | `string` | required, replaces on change |  | The target platform for the environment. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen |  | The version of the runtime engine for the environment. |
| `EnvironmentArn` | environment_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the runtime environment. |
| `EnvironmentId` | environment_id | `string` | computed |  | The unique identifier of the environment. |
| `HighAvailabilityConfig` | high_availability_config | `map` | optional, computed, provider-chosen |  | Defines the details of a high availability configuration. |
| `InstanceType` | instance_type | `string` | required |  | The type of instance underlying the environment. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources. |
| `Name` |  | `string` | required, replaces on change |  | The name of the environment. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the environment is publicly accessible. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The list of security groups for the VPC associated with this environment. |
| `StorageConfigurations` | storage_configurations | `list` | optional, computed, provider-chosen, replaces on change |  | The storage configurations defined for the runtime environment. |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | The unique identifiers of the subnets assigned to this runtime environment. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Defines tags associated to an environment. |

Supports update: yes

Discovery: supported
