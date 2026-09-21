# aws.customdataidentifier

**CloudFormation type:** `AWS::Macie::CustomDataIdentifier`

Macie CustomDataIdentifier resource schema

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Macie::CustomDataIdentifier)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Custom data identifier ARN. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of custom data identifier. |
| `Id` |  | `string` | computed |  | Custom data identifier ID. |
| `IgnoreWords` | ignore_words | `list` | optional, computed, provider-chosen, replaces on change |  | Words to be ignored. |
| `Keywords` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Keywords to be matched against. |
| `MaximumMatchDistance` | maximum_match_distance | `integer` | optional, computed, provider-chosen, replaces on change |  | Maximum match distance. |
| `Name` |  | `string` | required, replaces on change |  | Name of custom data identifier. |
| `Regex` |  | `string` | required, replaces on change |  | Regular expression for custom data identifier. |
| `Tags` |  | `map` | tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
