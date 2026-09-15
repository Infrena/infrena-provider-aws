# aws.enabledcontrol

**CloudFormation type:** `AWS::ControlTower::EnabledControl`

Enables a control on a specified target.

Region attribute: `region`

**Import ID:** `<region>/TargetIdentifier|ControlIdentifier` (AWS::ControlTower::EnabledControl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ControlIdentifier` | control_identifier | `string` | required, replaces on change |  | Arn of the control. |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  | Parameters to configure the enabled control behavior. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A set of tags to assign to the enabled control. |
| `TargetIdentifier` | target_identifier | `string` | required, replaces on change |  | Arn for Organizational unit to which the control needs to be applied |

Supports update: yes

Discovery: supported
