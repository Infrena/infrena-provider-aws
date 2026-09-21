# aws.iotwireless.taskdefinition

**CloudFormation type:** `AWS::IoTWireless::TaskDefinition`

Creates a gateway task definition.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::TaskDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | TaskDefinition arn. Returned after successful create. |
| `AutoCreateTasks` | auto_create_tasks | `boolean` | required |  | Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask. |
| `Id` |  | `string` | computed |  | The ID of the new wireless gateway task definition |
| `LoRaWANUpdateGatewayTaskEntry` | lo_ra_wan_update_gateway_task_entry | `map` | optional, computed, provider-chosen |  | The list of task definitions. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the new resource. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the destination. |
| `TaskDefinitionType` | task_definition_type | `string` | optional, computed, provider-chosen |  | A filter to list only the wireless gateway task definitions that use this task definition type |
| `Update` |  | `map` | optional, computed, provider-chosen |  | Information about the gateways to update. |

Supports update: yes

Discovery: supported
