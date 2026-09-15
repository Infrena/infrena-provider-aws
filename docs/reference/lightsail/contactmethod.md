# aws.contactmethod

**CloudFormation type:** `AWS::Lightsail::ContactMethod`

Resource Type definition for AWS::Lightsail::ContactMethod

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Lightsail::ContactMethod)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the contact method. |
| `ContactEndpoint` | contact_endpoint | `string` | required, replaces on change |  | The destination of the contact method, such as an email address or a mobile phone number. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the contact method was created. |
| `Name` |  | `string` | computed |  | The name of the contact method. |
| `Protocol` |  | `string` | required, replaces on change |  | The protocol of the contact method, such as Email or SMS (text messaging). |
| `ResourceType` | resource_type | `string` | computed |  | The Lightsail resource type of the contact method. |
| `Status` |  | `string` | computed |  | The current status of the contact method. |
| `SupportCode` | support_code | `string` | computed |  | The support code for the contact method. |

Supports update: no

Discovery: supported
