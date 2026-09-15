# aws.enabledbaseline

**CloudFormation type:** `AWS::ControlTower::EnabledBaseline`

Definition of AWS::ControlTower::EnabledBaseline Resource Type

Region attribute: `region`

**Import ID:** `<region>/EnabledBaselineIdentifier` (AWS::ControlTower::EnabledBaseline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BaselineIdentifier` | baseline_identifier | `string` | required, replaces on change |  |  |
| `BaselineVersion` | baseline_version | `string` | required |  |  |
| `EnabledBaselineIdentifier` | enabled_baseline_identifier | `string` | computed |  |  |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TargetIdentifier` | target_identifier | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
