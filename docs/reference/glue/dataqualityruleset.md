# aws.dataqualityruleset

**CloudFormation type:** `AWS::Glue::DataQualityRuleset`

Resource Type definition for AWS::Glue::DataQualityRuleset

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::DataQualityRuleset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, write-only |  | A unique token for idempotency. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the data quality ruleset. |
| `Name` |  | `string` | required, replaces on change |  | A unique name for the data quality ruleset. |
| `Ruleset` |  | `string` | optional, computed, provider-chosen |  | A Data Quality Definition Language (DQDL) ruleset. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of key-value pairs to apply to this resource. |
| `TargetTable` | target_table | `map` | optional, computed, provider-chosen |  | An object representing an AWS Glue table. |

Supports update: yes

Discovery: supported
