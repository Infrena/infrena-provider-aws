# aws.contactchannel

**CloudFormation type:** `AWS::SSMContacts::ContactChannel`

Resource Type definition for AWS::SSMContacts::ContactChannel

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMContacts::ContactChannel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the engagement to a contact channel. |
| `ChannelAddress` | channel_address | `string` | optional, computed, provider-chosen |  | The details that SSM Incident Manager uses when trying to engage the contact channel. |
| `ChannelName` | channel_name | `string` | optional, computed, provider-chosen |  | The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters. |
| `ChannelType` | channel_type | `string` | optional, computed, provider-chosen, replaces on change |  | Device type, which specify notification channel. Currently supported values: “SMS”, “VOICE”, “EMAIL”, “CHATBOT. |
| `ContactId` | contact_id | `string` | optional, computed, provider-chosen, replaces on change |  | ARN of the contact resource |
| `DeferActivation` | defer_activation | `boolean` | optional, computed, provider-chosen, write-only |  | If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated. |

Supports update: yes

Discovery: supported
