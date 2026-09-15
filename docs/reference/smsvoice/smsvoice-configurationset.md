# aws.smsvoice.configurationset

**CloudFormation type:** `AWS::SMSVOICE::ConfigurationSet`

Resource Type definition for AWS::SMSVOICE::ConfigurationSet

Region attribute: `region`

**Import ID:** `<region>/ConfigurationSetName` (AWS::SMSVOICE::ConfigurationSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ConfigurationSetName` | configuration_set_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name to use for the configuration set. |
| `DefaultSenderId` | default_sender_id | `string` | optional, computed, provider-chosen |  | The default sender ID to set for the ConfigurationSet. |
| `EventDestinations` | event_destinations | `list` | optional, computed, provider-chosen |  | An event destination is a location where you send message events. |
| `MessageFeedbackEnabled` | message_feedback_enabled | `boolean` | optional, computed, provider-chosen |  | Set to true to enable message feedback. |
| `ProtectConfigurationId` | protect_configuration_id | `string` | optional, computed, provider-chosen | aws.protectconfiguration.ProtectConfigurationId | The unique identifier for the protect configuration to be associated to the configuration set. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
