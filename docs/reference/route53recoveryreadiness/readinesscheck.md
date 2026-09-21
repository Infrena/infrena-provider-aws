# aws.readinesscheck

**CloudFormation type:** `AWS::Route53RecoveryReadiness::ReadinessCheck`

Aws Route53 Recovery Readiness Check Schema and API specification.

Region attribute: `region`

**Import ID:** `<region>/ReadinessCheckName` (AWS::Route53RecoveryReadiness::ReadinessCheck)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ReadinessCheckArn` | readiness_check_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the readiness check. |
| `ReadinessCheckName` | readiness_check_name | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the ReadinessCheck to create. |
| `ResourceSetName` | resource_set_name | `string` | optional, computed, provider-chosen |  | The name of the resource set to check. |
| `Tags` |  | `map` | tags map |  | A collection of tags associated with a resource. |

Supports update: yes

Discovery: supported
