# aws.smsvoice.phonenumber

**CloudFormation type:** `AWS::SMSVOICE::PhoneNumber`

Resource Type definition for AWS::SMSVOICE::PhoneNumber

Region attribute: `region`

**Import ID:** `<region>/PhoneNumberId` (AWS::SMSVOICE::PhoneNumber)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | When set to true the sender ID can't be deleted. By default this is set to false. |
| `IsoCountryCode` | iso_country_code | `string` | required, replaces on change |  | The two-character code, in ISO 3166-1 alpha-2 format, for the country or region. |
| `MandatoryKeywords` | mandatory_keywords | `map` | required |  | A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. Keywords "HELP" and "STOP" are mandatory keywords |
| `NumberCapabilities` | number_capabilities | `list` | required, replaces on change |  | Indicates if the phone number will be used for text messages, voice messages, or both. |
| `NumberType` | number_type | `string` | required, replaces on change |  | The type of phone number to request. |
| `OptOutListName` | opt_out_list_name | `string` | optional, computed, provider-chosen |  | The name of the OptOutList to associate with the phone number. You can use the OptOutListName or OptOutListArn. |
| `OptionalKeywords` | optional_keywords | `list` | optional, computed, provider-chosen |  | A keyword is a word that you can search for on a particular phone number or pool. It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, AWS End User Messaging SMS and Voice responds with a customizable message. |
| `PhoneNumber` | phone_number | `string` | computed |  |  |
| `PhoneNumberId` | phone_number_id | `string` | computed |  |  |
| `SelfManagedOptOutsEnabled` | self_managed_opt_outs_enabled | `boolean` | optional, computed, provider-chosen |  | By default this is set to false. When an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, AWS End User Messaging SMS and Voice automatically replies with a customizable message and adds the end recipient to the OptOutList. When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out requests. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TwoWay` | two_way | `map` | optional, computed, provider-chosen |  | When you set up two-way SMS, you can receive incoming messages from your customers. When one of your customers sends a message to your phone number, the message body is sent to an Amazon SNS topic or Amazon Connect for processing. |

Supports update: yes

Discovery: supported
