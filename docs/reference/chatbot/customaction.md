# aws.customaction

**CloudFormation type:** `AWS::Chatbot::CustomAction`

Definition of AWS::Chatbot::CustomAction Resource Type

Region attribute: `region`

**Import ID:** `<region>/CustomActionArn` (AWS::Chatbot::CustomAction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionName` | action_name | `string` | required, replaces on change |  |  |
| `AliasName` | alias_name | `string` | optional, computed, provider-chosen |  |  |
| `Attachments` |  | `list` | optional, computed, provider-chosen |  |  |
| `CustomActionArn` | custom_action_arn | `string` | computed |  |  |
| `Definition` |  | `map` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
