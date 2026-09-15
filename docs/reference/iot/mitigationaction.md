# aws.mitigationaction

**CloudFormation type:** `AWS::IoT::MitigationAction`

Mitigation actions can be used to take actions to mitigate issues that were found in an Audit finding or Detect violation.

Region attribute: `region`

**Import ID:** `<region>/ActionName` (AWS::IoT::MitigationAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionName` | action_name | `string` | optional, computed, provider-chosen, replaces on change |  | A unique identifier for the mitigation action. |
| `ActionParams` | action_params | `map` | required |  | The set of parameters for this mitigation action. You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action). |
| `MitigationActionArn` | mitigation_action_arn | `string` | computed |  |  |
| `MitigationActionId` | mitigation_action_id | `string` | computed |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
