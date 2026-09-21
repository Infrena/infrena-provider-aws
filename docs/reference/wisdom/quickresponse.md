# aws.quickresponse

**CloudFormation type:** `AWS::Wisdom::QuickResponse`

Definition of AWS::Wisdom::QuickResponse Resource Type.

Region attribute: `region`

**Import ID:** `<region>/QuickResponseArn` (AWS::Wisdom::QuickResponse)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Channels` |  | `list` | optional, computed, provider-chosen |  | The Amazon Connect contact channels this quick response applies to. |
| `Content` |  | `map` | required |  | The container of quick response content. |
| `ContentType` | content_type | `string` | optional, computed, provider-chosen |  | The media type of the quick response content. |
| `Contents` |  | `map` | computed |  | The content of the quick response stored in different media types. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the quick response. |
| `GroupingConfiguration` | grouping_configuration | `map` | optional, computed, provider-chosen |  | The configuration information of the user groups that the quick response is accessible to. |
| `IsActive` | is_active | `boolean` | optional, computed, provider-chosen |  | Whether the quick response is active. |
| `KnowledgeBaseArn` | knowledge_base_arn | `string` | required, replaces on change | aws.wisdom.knowledgebase.KnowledgeBaseArn | The Amazon Resource Name (ARN) of the knowledge base. |
| `Language` |  | `string` | optional, computed, provider-chosen |  | The language code value for the language in which the quick response is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW |
| `Name` |  | `string` | required |  | The name of the quick response. |
| `QuickResponseArn` | quick_response_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the quick response. |
| `QuickResponseId` | quick_response_id | `string` | computed |  | The identifier of the quick response. |
| `ShortcutKey` | shortcut_key | `string` | optional, computed, provider-chosen |  | The shortcut key of the quick response. The value should be unique across the knowledge base. |
| `Status` |  | `string` | computed |  | The status of the quick response data. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
