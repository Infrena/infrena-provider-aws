# aws.databrew.dataset

**CloudFormation type:** `AWS::DataBrew::Dataset`

Resource schema for AWS::DataBrew::Dataset.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Format` |  | `string` | optional, computed, provider-chosen |  | Dataset format |
| `FormatOptions` | format_options | `map` | optional, computed, provider-chosen |  | Format options for dataset |
| `Input` |  | `map` | required |  | Input |
| `Name` |  | `string` | required, replaces on change |  | Dataset name |
| `PathOptions` | path_options | `map` | optional, computed, provider-chosen |  | Path options for dataset |
| `Source` |  | `string` | optional, computed, provider-chosen |  | Source type of the dataset |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
