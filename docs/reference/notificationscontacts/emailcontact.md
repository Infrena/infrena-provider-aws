# aws.emailcontact

**CloudFormation type:** `AWS::NotificationsContacts::EmailContact`

Definition of AWS::NotificationsContacts::EmailContact Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::NotificationsContacts::EmailContact)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `EmailAddress` | email_address | `string` | required, replaces on change |  |  |
| `EmailContact` | email_contact | `map` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | replaces on change, tags map |  | A list of tags that are attached to the role. |

Supports update: no

Discovery: supported
