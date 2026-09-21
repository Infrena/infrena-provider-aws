# aws.customentitytype

**CloudFormation type:** `AWS::Glue::CustomEntityType`

Resource Type definition for AWS::Glue::CustomEntityType

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::CustomEntityType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContextWords` | context_words | `list` | optional, computed, provider-chosen |  | A list of context words. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the custom entity type. |
| `RegexString` | regex_string | `string` | optional, computed, provider-chosen |  | A regular expression string that is used for detecting sensitive data in a custom pattern. |
| `Tags` |  | `map` | optional, computed, provider-chosen, write-only |  | Tags to associate with the custom entity type. |

Supports update: yes

Discovery: supported
