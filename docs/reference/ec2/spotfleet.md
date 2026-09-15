# aws.spotfleet

**CloudFormation type:** `AWS::EC2::SpotFleet`

Resource Type definition for AWS::EC2::SpotFleet

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SpotFleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `SpotFleetRequestConfigData` | spot_fleet_request_config_data | `map` | required |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | The tags to specify in SpotFleetRequestConfigData |

Supports update: yes

Discovery: supported
