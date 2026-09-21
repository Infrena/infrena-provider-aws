# aws.routerinput

**CloudFormation type:** `AWS::MediaConnect::RouterInput`

Represents a router input in AWS Elemental MediaConnect that is used to ingest content to be transmitted to router outputs

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaConnect::RouterInput)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone where you want to create the router input. This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided. |
| `Configuration` |  | `string` | required |  | The configuration settings for a router input. |
| `ContentQualityAnalysisConfiguration` | content_quality_analysis_configuration | `string` | optional, computed, provider-chosen |  | The content quality analysis configuration for the router input. The content quality analysis feature only monitors the first video stream and the first audio stream it encounters within the router input source. |
| `ContentQualityAnalysisType` | content_quality_analysis_type | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the router input was created. |
| `Id` |  | `string` | computed |  | The unique identifier of the router input. |
| `InputType` | input_type | `string` | computed |  |  |
| `IpAddress` | ip_address | `string` | computed |  | The IP address of the router input. |
| `MaintenanceConfiguration` | maintenance_configuration | `string` | optional, computed, provider-chosen |  | The configuration settings for maintenance operations, including preferred maintenance windows and schedules. |
| `MaintenanceType` | maintenance_type | `string` | computed |  |  |
| `MaximumBitrate` | maximum_bitrate | `integer` | required |  | The maximum bitrate for the router input. |
| `Name` |  | `string` | required |  | The name of the router input. |
| `RegionName` | region_name | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Web Services Region for the router input. Defaults to the current region if not specified. |
| `RoutedOutputs` | routed_outputs | `integer` | computed |  | The number of router outputs associated with the router input. |
| `RoutingScope` | routing_scope | `string` | required |  |  |
| `State` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to tag and organize this router input. |
| `Tier` |  | `string` | required |  |  |
| `TransitEncryption` | transit_encryption | `map` | optional, computed, provider-chosen |  | The transit encryption settings for a router input. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the router input was last updated. |

Supports update: yes

Discovery: supported
