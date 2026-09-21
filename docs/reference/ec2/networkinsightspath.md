# aws.networkinsightspath

**CloudFormation type:** `AWS::EC2::NetworkInsightsPath`

Resource schema for AWS::EC2::NetworkInsightsPath

Region attribute: `region`

**Import ID:** `<region>/NetworkInsightsPathId` (AWS::EC2::NetworkInsightsPath)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedDate` | created_date | `string` | computed |  |  |
| `Destination` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DestinationArn` | destination_arn | `string` | computed |  |  |
| `DestinationIp` | destination_ip | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DestinationPort` | destination_port | `integer` | optional, computed, provider-chosen, replaces on change |  |  |
| `FilterAtDestination` | filter_at_destination | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `FilterAtSource` | filter_at_source | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `NetworkInsightsPathArn` | network_insights_path_arn | `string` | computed |  |  |
| `NetworkInsightsPathId` | network_insights_path_id | `string` | computed |  |  |
| `Protocol` |  | `string` | required, replaces on change |  |  |
| `Source` |  | `string` | required, replaces on change |  |  |
| `SourceArn` | source_arn | `string` | computed |  |  |
| `SourceIp` | source_ip | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
