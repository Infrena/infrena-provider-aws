# aws.customverificationemailtemplate

**CloudFormation type:** `AWS::SES::CustomVerificationEmailTemplate`

Resource Type definition for AWS::SES::CustomVerificationEmailTemplate.

Region attribute: `region`

**Import ID:** `<region>/TemplateName` (AWS::SES::CustomVerificationEmailTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FailureRedirectionURL` | failure_redirection_url | `string` | required |  | The URL that the recipient of the verification email is sent to if his or her address is not successfully verified. |
| `FromEmailAddress` | from_email_address | `string` | required |  | The email address that the custom verification email is sent from. |
| `SuccessRedirectionURL` | success_redirection_url | `string` | required |  | The URL that the recipient of the verification email is sent to if his or her address is successfully verified. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the tenant. |
| `TemplateContent` | template_content | `string` | required |  | The content of the custom verification email. The total size of the email must be less than 10 MB. The message body may contain HTML, with some limitations. |
| `TemplateName` | template_name | `string` | required, replaces on change |  | The name of the custom verification email template. |
| `TemplateSubject` | template_subject | `string` | required |  | The subject line of the custom verification email. |

Supports update: yes

Discovery: supported
