# aws.deadline.monitor

**CloudFormation type:** `AWS::Deadline::Monitor`

Resource Type definition for AWS::Deadline::Monitor

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Monitor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `IdentityCenterApplicationArn` | identity_center_application_arn | `string` | computed |  |  |
| `IdentityCenterInstanceArn` | identity_center_instance_arn | `string` | required, replaces on change |  |  |
| `IdentityCenterRegion` | identity_center_region | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS region where IAM Identity Center is enabled. Required when Identity Center is in a different region than the monitor. |
| `MonitorId` | monitor_id | `string` | computed |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Subdomain` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Url` |  | `string` | computed |  |  |

Supports update: yes

Discovery: supported
