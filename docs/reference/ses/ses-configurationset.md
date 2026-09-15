# aws.ses.configurationset

**CloudFormation type:** `AWS::SES::ConfigurationSet`

Resource schema for AWS::SES::ConfigurationSet.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::SES::ConfigurationSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArchivingOptions` | archiving_options | `map` | optional, computed, provider-chosen |  | An object that defines a MailManager archive that is used to preserve emails that you send using the configuration set. |
| `DeliveryOptions` | delivery_options | `map` | optional, computed, provider-chosen |  | An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the configuration set. |
| `ReputationOptions` | reputation_options | `map` | optional, computed, provider-chosen |  | An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set. |
| `SendingOptions` | sending_options | `map` | optional, computed, provider-chosen |  | An object that defines whether or not Amazon SES can send email that you send using the configuration set. |
| `SuppressionOptions` | suppression_options | `map` | optional, computed, provider-chosen |  | An object that contains information about the suppression list preferences for your account. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the contact list. |
| `TrackingOptions` | tracking_options | `map` | optional, computed, provider-chosen |  | An object that defines the open and click tracking options for emails that you send using the configuration set. |
| `VdmOptions` | vdm_options | `map` | optional, computed, provider-chosen |  | An object that contains Virtual Deliverability Manager (VDM) settings for this configuration set. |

Supports update: yes

Discovery: supported
