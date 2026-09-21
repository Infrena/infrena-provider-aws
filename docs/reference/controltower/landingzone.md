# aws.landingzone

**CloudFormation type:** `AWS::ControlTower::LandingZone`

Definition of AWS::ControlTower::LandingZone Resource Type

Region attribute: `region`

**Import ID:** `<region>/LandingZoneIdentifier` (AWS::ControlTower::LandingZone)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DriftStatus` | drift_status | `string` | computed |  |  |
| `LandingZoneIdentifier` | landing_zone_identifier | `string` | computed |  |  |
| `LatestAvailableVersion` | latest_available_version | `string` | computed |  |  |
| `Manifest` |  | `string` | required |  |  |
| `RemediationTypes` | remediation_types | `list` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Version` |  | `string` | required |  |  |

Supports update: yes

Discovery: supported
