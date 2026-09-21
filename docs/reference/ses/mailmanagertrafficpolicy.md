# aws.mailmanagertrafficpolicy

**CloudFormation type:** `AWS::SES::MailManagerTrafficPolicy`

Definition of AWS::SES::MailManagerTrafficPolicy Resource Type

Region attribute: `region`

**Import ID:** `<region>/TrafficPolicyId` (AWS::SES::MailManagerTrafficPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DefaultAction` | default_action | `string` | required |  |  |
| `MaxMessageSizeBytes` | max_message_size_bytes | `float` | optional, computed, provider-chosen |  |  |
| `PolicyStatements` | policy_statements | `list` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TrafficPolicyArn` | traffic_policy_arn | `string` | computed |  |  |
| `TrafficPolicyId` | traffic_policy_id | `string` | computed |  |  |
| `TrafficPolicyName` | traffic_policy_name | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
