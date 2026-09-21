# aws.multiplex

**CloudFormation type:** `AWS::MediaLive::Multiplex`

Resource schema for AWS::MediaLive::Multiplex

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaLive::Multiplex)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The unique arn of the multiplex. |
| `AvailabilityZones` | availability_zones | `list` | required, replaces on change |  | A list of availability zones for the multiplex. |
| `Destinations` |  | `list` | optional, computed, provider-chosen |  | A list of the multiplex output destinations. |
| `Id` |  | `string` | computed |  | The unique id of the multiplex. |
| `MultiplexSettings` | multiplex_settings | `map` | required |  | A key-value pair to associate with a resource. |
| `Name` |  | `string` | required |  | Name of multiplex. |
| `PipelinesRunningCount` | pipelines_running_count | `integer` | computed |  | The number of currently healthy pipelines. |
| `ProgramCount` | program_count | `integer` | computed |  | The number of programs in the multiplex. |
| `State` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | A collection of key-value pairs. |

Supports update: yes

Discovery: supported
