# aws.emailidentity

**CloudFormation type:** `AWS::SES::EmailIdentity`

Resource Type definition for AWS::SES::EmailIdentity

Region attribute: `region`

**Import ID:** `<region>/EmailIdentity` (AWS::SES::EmailIdentity)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigurationSetAttributes` | configuration_set_attributes | `map` | optional, computed, provider-chosen |  | Used to associate a configuration set with an email identity. |
| `DkimAttributes` | dkim_attributes | `map` | optional, computed, provider-chosen |  | Used to enable or disable DKIM authentication for an email identity. |
| `DkimDNSTokenName1` | dkim_dns_token_name1 | `string` | computed |  |  |
| `DkimDNSTokenName2` | dkim_dns_token_name2 | `string` | computed |  |  |
| `DkimDNSTokenName3` | dkim_dns_token_name3 | `string` | computed |  |  |
| `DkimDNSTokenValue1` | dkim_dns_token_value1 | `string` | computed |  |  |
| `DkimDNSTokenValue2` | dkim_dns_token_value2 | `string` | computed |  |  |
| `DkimDNSTokenValue3` | dkim_dns_token_value3 | `string` | computed |  |  |
| `DkimSigningAttributes` | dkim_signing_attributes | `map` | optional, computed, provider-chosen |  | If your request includes this object, Amazon SES configures the identity to use Bring Your Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be used for Easy DKIM. |
| `EmailIdentity` | email_identity | `string` | required, replaces on change |  | The email address or domain to verify. |
| `FeedbackAttributes` | feedback_attributes | `map` | optional, computed, provider-chosen |  | Used to enable or disable feedback forwarding for an identity. |
| `MailFromAttributes` | mail_from_attributes | `map` | optional, computed, provider-chosen |  | Used to enable or disable the custom Mail-From domain configuration for an email identity. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the email identity. |

Supports update: yes

Discovery: supported
