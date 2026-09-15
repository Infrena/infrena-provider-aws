# aws.recoverygroup

**CloudFormation type:** `AWS::Route53RecoveryReadiness::RecoveryGroup`

AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.

Region attribute: `region`

**Import ID:** `<region>/RecoveryGroupName` (AWS::Route53RecoveryReadiness::RecoveryGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Cells` |  | `list` | optional, computed, provider-chosen |  | A list of the cell Amazon Resource Names (ARNs) in the recovery group. |
| `RecoveryGroupArn` | recovery_group_arn | `string` | computed |  | A collection of tags associated with a resource. |
| `RecoveryGroupName` | recovery_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the recovery group to create. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource. |

Supports update: yes

Discovery: supported
