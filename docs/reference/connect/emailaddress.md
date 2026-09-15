# aws.emailaddress

**CloudFormation type:** `AWS::Connect::EmailAddress`

Resource Type definition for AWS::Connect::EmailAddress

Region attribute: `region`

**Import ID:** `<region>/EmailAddressArn` (AWS::Connect::EmailAddress)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasConfigurations` | alias_configurations | `list` | optional, computed, provider-chosen |  | List of alias configurations for the email address |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the email address. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The display name for the email address. |
| `EmailAddress` | email_address | `string` | required, replaces on change |  | Email address to be created for this instance |
| `EmailAddressArn` | email_address_arn | `string` | computed |  | The identifier of the email address. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |

Supports update: yes

Discovery: supported (parent resource required)
