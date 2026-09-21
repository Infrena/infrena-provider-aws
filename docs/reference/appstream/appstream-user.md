# aws.appstream.user

**CloudFormation type:** `AWS::AppStream::User`

Resource Type definition for AWS::AppStream::User

Region attribute: `region`

**Import ID:** `<region>/UserName|AuthenticationType` (AWS::AppStream::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Returns the Amazon Resource Name (ARN) for the Amazon AppStream User resource. |
| `AuthenticationType` | authentication_type | `string` | required, replaces on change |  | The authentication type for the user. |
| `FirstName` | first_name | `string` | optional, computed, provider-chosen, replaces on change |  | The first name, or given name, of the user. |
| `LastName` | last_name | `string` | optional, computed, provider-chosen, replaces on change |  | The last name, or surname, of the user. |
| `MessageAction` | message_action | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The action to take for the welcome email that is sent to a user after the user is created in the user pool. If you specify SUPPRESS, no email is sent. If you specify RESEND, do not specify the first name or last name of the user. If the value is null, the email is sent. |
| `UserName` | user_name | `string` | required, replaces on change |  | The email address of the user. |

Supports update: no

Discovery: supported
