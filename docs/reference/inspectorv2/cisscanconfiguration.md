# aws.cisscanconfiguration

**CloudFormation type:** `AWS::InspectorV2::CisScanConfiguration`

CIS Scan Configuration resource schema

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::InspectorV2::CisScanConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | CIS Scan configuration unique identifier |
| `ScanName` | scan_name | `string` | required |  | Name of the scan |
| `Schedule` |  | `map` | required |  | Choose a Schedule cadence |
| `SecurityLevel` | security_level | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Targets` |  | `map` | required |  |  |

Supports update: yes

Discovery: supported
