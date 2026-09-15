# aws.recipe

**CloudFormation type:** `AWS::DataBrew::Recipe`

Resource schema for AWS::DataBrew::Recipe.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Recipe)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the recipe |
| `Name` |  | `string` | required, replaces on change |  | Recipe name |
| `Steps` |  | `list` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
