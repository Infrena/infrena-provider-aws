# aws.botalias

**CloudFormation type:** `AWS::Lex::BotAlias`

Resource Type definition for a Bot Alias, which enables you to change the version of a bot without updating applications that use the bot

Region attribute: `region`

**Import ID:** `<region>/BotAliasId|BotId` (AWS::Lex::BotAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the bot alias. |
| `BotAliasId` | bot_alias_id | `string` | computed |  | Unique ID of resource |
| `BotAliasLocaleSettings` | bot_alias_locale_settings | `list` | optional, computed, provider-chosen |  | A list of bot alias locale settings to add to the bot alias. |
| `BotAliasName` | bot_alias_name | `string` | required |  | A unique identifier for a resource. |
| `BotAliasStatus` | bot_alias_status | `string` | computed |  |  |
| `BotAliasTags` | bot_alias_tags | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to add to the bot alias. |
| `BotId` | bot_id | `string` | required, replaces on change | aws.bot.Id | Unique ID of resource |
| `BotVersion` | bot_version | `string` | optional, computed, provider-chosen |  | The version of a bot. |
| `ConversationLogSettings` | conversation_log_settings | `map` | optional, computed, provider-chosen |  | Contains information about code hooks that Amazon Lex calls during a conversation. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the bot alias. Use the description to help identify the bot alias in lists. |
| `SentimentAnalysisSettings` | sentiment_analysis_settings | `map` | optional, computed, provider-chosen |  | Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances. |

Supports update: yes

Discovery: supported (parent resource required)
