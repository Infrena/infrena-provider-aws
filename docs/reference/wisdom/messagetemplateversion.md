# aws.messagetemplateversion

**CloudFormation type:** `AWS::Wisdom::MessageTemplateVersion`

A version for the specified customer-managed message template within the specified knowledge base.

Region attribute: `region`

**Import ID:** `<region>/MessageTemplateVersionArn` (AWS::Wisdom::MessageTemplateVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `MessageTemplateArn` | message_template_arn | `string` | required, replaces on change | aws.messagetemplate.MessageTemplateArn | The unqualified Amazon Resource Name (ARN) of the message template. |
| `MessageTemplateContentSha256` | message_template_content_sha256 | `string` | optional, computed, provider-chosen |  | The content SHA256 of the message template. |
| `MessageTemplateVersionArn` | message_template_version_arn | `string` | computed |  | The unqualified Amazon Resource Name (ARN) of the message template version. |
| `MessageTemplateVersionNumber` | message_template_version_number | `float` | computed |  | Current version number of the message template. |

Supports update: yes

Discovery: supported (parent resource required)
