# aws.ruleset

**CloudFormation type:** `AWS::DataBrew::Ruleset`

Resource schema for AWS::DataBrew::Ruleset.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Ruleset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Ruleset |
| `Name` |  | `string` | required, replaces on change |  | Name of the Ruleset |
| `Rules` |  | `list` | required |  | List of the data quality rules in the ruleset |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TargetArn` | target_arn | `string` | required, replaces on change |  | Arn of the target resource (dataset) to apply the ruleset to |

Supports update: yes

Discovery: supported
