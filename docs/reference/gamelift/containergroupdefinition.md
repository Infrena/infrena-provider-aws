# aws.containergroupdefinition

**CloudFormation type:** `AWS::GameLift::ContainerGroupDefinition`

The AWS::GameLift::ContainerGroupDefinition resource creates an Amazon GameLift container group definition.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::GameLift::ContainerGroupDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContainerGroupDefinitionArn` | container_group_definition_arn | `string` | computed |  | The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions. |
| `ContainerGroupType` | container_group_type | `string` | optional, computed, provider-chosen, replaces on change |  | The scope of the container group |
| `CreationTime` | creation_time | `string` | computed |  | A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057"). |
| `GameServerContainerDefinition` | game_server_container_definition | `map` | optional, computed, provider-chosen |  | Specifies the information required to run game servers with this container group |
| `Name` |  | `string` | required, replaces on change |  | A descriptive label for the container group definition. |
| `OperatingSystem` | operating_system | `string` | required |  | The operating system of the container group |
| `SourceVersionNumber` | source_version_number | `integer` | optional, computed, provider-chosen, write-only |  | A specific ContainerGroupDefinition version to be updated |
| `Status` |  | `string` | computed |  | A string indicating ContainerGroupDefinition status. |
| `StatusReason` | status_reason | `string` | computed |  | A string indicating the reason for ContainerGroupDefinition status. |
| `SupportContainerDefinitions` | support_container_definitions | `list` | optional, computed, provider-chosen |  | A collection of support container definitions that define the containers in this group. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TotalMemoryLimitMebibytes` | total_memory_limit_mebibytes | `integer` | required |  | The total memory limit of container groups following this definition in MiB |
| `TotalVcpuLimit` | total_vcpu_limit | `float` | required |  | The total amount of virtual CPUs on the container group definition |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen |  | The description of this version |
| `VersionNumber` | version_number | `integer` | computed |  | The version of this ContainerGroupDefinition |

Supports update: yes

Discovery: supported
