# aws.senderid

**CloudFormation type:** `AWS::SMSVOICE::SenderId`

Resource Type definition for AWS::SMSVOICE::SenderId

Region attribute: `region`

**Import ID:** `<region>/IsoCountryCode|SenderId` (AWS::SMSVOICE::SenderId)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) associated with the SenderId. |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | When set to true the sender ID can't be deleted. By default this is set to false. |
| `IsoCountryCode` | iso_country_code | `string` | required, replaces on change |  | The two-character code, in ISO 3166-1 alpha-2 format, for the country or region. |
| `SenderId` | sender_id | `string` | required, replaces on change |  | The sender ID string to request. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
