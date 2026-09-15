# aws.mailmanageringresspoint

**CloudFormation type:** `AWS::SES::MailManagerIngressPoint`

Definition of AWS::SES::MailManagerIngressPoint Resource Type

Region attribute: `region`

**Import ID:** `<region>/IngressPointId` (AWS::SES::MailManagerIngressPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARecord` | a_record | `string` | computed |  |  |
| `IngressPointArn` | ingress_point_arn | `string` | computed |  |  |
| `IngressPointConfiguration` | ingress_point_configuration | `string` | optional, computed, provider-chosen, write-only |  |  |
| `IngressPointId` | ingress_point_id | `string` | computed |  |  |
| `IngressPointName` | ingress_point_name | `string` | optional, computed, provider-chosen |  |  |
| `NetworkConfiguration` | network_configuration | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RuleSetId` | rule_set_id | `string` | required |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusToUpdate` | status_to_update | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TlsPolicy` | tls_policy | `string` | optional, computed, provider-chosen |  |  |
| `TrafficPolicyId` | traffic_policy_id | `string` | required |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
