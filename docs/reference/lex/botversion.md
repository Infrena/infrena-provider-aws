# aws.botversion

**CloudFormation type:** `AWS::Lex::BotVersion`

Resource Type definition for bot versions, a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.

Region attribute: `region`

**Import ID:** `<region>/BotId|BotVersion` (AWS::Lex::BotVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BotId` | bot_id | `string` | required, replaces on change | aws.bot.Id | Unique ID of resource |
| `BotVersion` | bot_version | `string` | computed |  | The version of a bot. |
| `BotVersionLocaleSpecification` | bot_version_locale_specification | `list` | required, replaces on change, write-only |  | Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the version. Use the description to help identify the version in lists. |

Supports update: no

Discovery: supported (parent resource required)
