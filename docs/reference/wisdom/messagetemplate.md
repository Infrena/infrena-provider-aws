# aws.messagetemplate

**CloudFormation type:** `AWS::Wisdom::MessageTemplate`

Definition of AWS::Wisdom::MessageTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/MessageTemplateArn` (AWS::Wisdom::MessageTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelSubtype` | channel_subtype | `string` | required, replaces on change |  | The channel subtype this message template applies to. |
| `Content` |  | `map` | required |  | The content of the message template. |
| `DefaultAttributes` | default_attributes | `map` | optional, computed, provider-chosen |  | An object that specifies the default values to use for variables in the message template. This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the message template. |
| `GroupingConfiguration` | grouping_configuration | `map` | optional, computed, provider-chosen |  | The configuration information of the user groups that the message template is accessible to. |
| `KnowledgeBaseArn` | knowledge_base_arn | `string` | required, replaces on change | aws.wisdom.knowledgebase.KnowledgeBaseArn | The Amazon Resource Name (ARN) of the knowledge base to which the message template belongs. |
| `Language` |  | `string` | optional, computed, provider-chosen |  | The language code value for the language in which the message template is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW |
| `MessageTemplateArn` | message_template_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the message template. |
| `MessageTemplateAttachments` | message_template_attachments | `list` | optional, computed, provider-chosen |  | List of message template attachments |
| `MessageTemplateContentSha256` | message_template_content_sha256 | `string` | computed |  | The content SHA256 of the message template. |
| `MessageTemplateId` | message_template_id | `string` | computed |  | The unique identifier of the message template. |
| `Name` |  | `string` | required |  | The name of the message template. |
| `Tags` |  | `map` | tags map |  | The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }. |

Supports update: yes

Discovery: supported (parent resource required)
