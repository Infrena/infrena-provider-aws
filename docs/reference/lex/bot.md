# aws.bot

**CloudFormation type:** `AWS::Lex::Bot`

Resource Type definition for an Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Lex::Bot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AutoBuildBotLocales` | auto_build_bot_locales | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `BotFileS3Location` | bot_file_s3_location | `map` | optional, computed, provider-chosen, write-only |  |  |
| `BotLocales` | bot_locales | `list` | optional, computed, provider-chosen, write-only |  |  |
| `BotMembers` | bot_members | `list` | optional, computed, provider-chosen, write-only |  | The list of bot members in a network to be created. |
| `BotTags` | bot_tags | `map` | optional, computed, provider-chosen, tags map |  |  |
| `BotType` | bot_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of a bot to create. |
| `DataPrivacy` | data_privacy | `map` | required |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the resource |
| `ErrorLogSettings` | error_log_settings | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `IdleSessionTTLInSeconds` | idle_session_ttl_in_seconds | `integer` | required |  |  |
| `Name` |  | `string` | required |  |  |
| `Replication` |  | `map` | optional, computed, provider-chosen, write-only |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `TestBotAliasSettings` | test_bot_alias_settings | `map` | optional, computed, provider-chosen, write-only |  |  |
| `TestBotAliasTags` | test_bot_alias_tags | `list` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
