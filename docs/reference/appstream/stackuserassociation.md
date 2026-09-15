# aws.stackuserassociation

**CloudFormation type:** `AWS::AppStream::StackUserAssociation`

Resource Type definition for AWS::AppStream::StackUserAssociation

Region attribute: `region`

**Import ID:** `<region>/StackName|UserName|AuthenticationType` (AWS::AppStream::StackUserAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthenticationType` | authentication_type | `string` | required, replaces on change |  | The authentication type for the user who is associated with the stack. You must specify USERPOOL. |
| `SendEmailNotification` | send_email_notification | `boolean` | optional, computed, provider-chosen, replaces on change |  | Specifies whether a welcome email is sent to a user after the user is created in the user pool. |
| `StackName` | stack_name | `string` | required, replaces on change |  | The name of the stack that is associated with the user. |
| `UserName` | user_name | `string` | required, replaces on change |  | The name of the user who is associated with the stack. |

Supports update: no

Discovery: supported
