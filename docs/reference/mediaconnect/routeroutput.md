# aws.routeroutput

**CloudFormation type:** `AWS::MediaConnect::RouterOutput`

Represents a router input in AWS Elemental MediaConnect that can be used to egress content transmitted from router inputs

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaConnect::RouterOutput)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone where you want to create the router output. This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided. |
| `Configuration` |  | `string` | required |  | The configuration settings for a router output. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the router output was created. |
| `FabricConfiguration` | fabric_configuration | `map` | optional, computed, provider-chosen |  | The fabric configuration settings for the router output. |
| `Id` |  | `string` | computed |  | The unique identifier of the router output. |
| `IpAddress` | ip_address | `string` | computed |  | The IP address of the router output. |
| `MaintenanceConfiguration` | maintenance_configuration | `string` | optional, computed, provider-chosen |  | The configuration settings for maintenance operations, including preferred maintenance windows and schedules. |
| `MaintenanceType` | maintenance_type | `string` | computed |  |  |
| `MaximumBitrate` | maximum_bitrate | `integer` | required |  | The maximum bitrate for the router output. |
| `Name` |  | `string` | required |  | The name of the router output. |
| `OutputType` | output_type | `string` | computed |  |  |
| `RegionName` | region_name | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Web Services Region for the router output. Defaults to the current region if not specified. |
| `RoutedState` | routed_state | `string` | computed |  |  |
| `RoutingScope` | routing_scope | `string` | required |  |  |
| `State` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to tag this router output. |
| `Tier` |  | `string` | required |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the router output was last updated. |

Supports update: yes

Discovery: supported
