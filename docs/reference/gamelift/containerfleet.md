# aws.containerfleet

**CloudFormation type:** `AWS::GameLift::ContainerFleet`

The AWS::GameLift::ContainerFleet resource creates an Amazon GameLift (GameLift) container fleet to host game servers.

Region attribute: `region`

**Import ID:** `<region>/FleetId` (AWS::GameLift::ContainerFleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BillingType` | billing_type | `string` | optional, computed, provider-chosen, replaces on change |  | Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet. |
| `CreationTime` | creation_time | `string` | computed |  | A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057"). |
| `DeploymentConfiguration` | deployment_configuration | `map` | optional, computed, provider-chosen, write-only |  | Provides details about how to drain old tasks and replace them with new updated tasks. |
| `DeploymentDetails` | deployment_details | `map` | computed |  | Provides information about the last deployment ID and its status. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A human-readable description of a fleet. |
| `FleetArn` | fleet_arn | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container fleet resource and uniquely identifies it across all AWS Regions. |
| `FleetId` | fleet_id | `string` | computed |  | Unique fleet ID |
| `FleetRoleArn` | fleet_role_arn | `string` | required | aws.role.Arn | A unique identifier for an AWS IAM role that manages access to your AWS services. Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console. |
| `GameServerContainerGroupDefinitionArn` | game_server_container_group_definition_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the game server container group definition. This field will be empty if GameServerContainerGroupDefinitionName is not specified. |
| `GameServerContainerGroupDefinitionName` | game_server_container_group_definition_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the container group definition that will be created per game server. You must specify GAME_SERVER container group. You have the option to also specify one PER_INSTANCE container group. |
| `GameServerContainerGroupsPerInstance` | game_server_container_groups_per_instance | `integer` | optional, computed, provider-chosen |  | The number of desired game server container groups per instance, a number between 1-5000. |
| `GameSessionCreationLimitPolicy` | game_session_creation_limit_policy | `map` | optional, computed, provider-chosen |  | A policy that limits the number of game sessions a player can create on the same fleet. This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: "An individual player can create a maximum number of new game sessions within a specified time period". |
| `InstanceConnectionPortRange` | instance_connection_port_range | `map` | optional, computed, provider-chosen |  | Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet. |
| `InstanceInboundPermissions` | instance_inbound_permissions | `list` | optional, computed, provider-chosen |  | A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server. |
| `InstanceType` | instance_type | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions. |
| `Locations` |  | `list` | optional, computed, provider-chosen |  |  |
| `LogConfiguration` | log_configuration | `map` | optional, computed, provider-chosen |  | A policy the location and provider of logs from the fleet. |
| `MaximumGameServerContainerGroupsPerInstance` | maximum_game_server_container_groups_per_instance | `integer` | computed |  | The maximum number of game server container groups per instance, a number between 1-5000. |
| `MetricGroups` | metric_groups | `list` | optional, computed, provider-chosen |  | The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string. |
| `NewGameSessionProtectionPolicy` | new_game_session_protection_policy | `string` | optional, computed, provider-chosen |  | A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions. |
| `PerInstanceContainerGroupDefinitionArn` | per_instance_container_group_definition_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the per instance container group definition. This field will be empty if PerInstanceContainerGroupDefinitionName is not specified. |
| `PerInstanceContainerGroupDefinitionName` | per_instance_container_group_definition_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the container group definition that will be created per instance. This field is optional if you specify GameServerContainerGroupDefinitionName. |
| `PlayerGatewayMode` | player_gateway_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The player gateway mode for the container fleet. |
| `ScalingPolicies` | scaling_policies | `list` | optional, computed, provider-chosen |  | A list of rules that control how a fleet is scaled. |
| `Status` |  | `string` | computed |  | The current status of the container fleet. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
