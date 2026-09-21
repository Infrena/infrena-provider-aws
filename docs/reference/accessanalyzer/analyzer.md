# aws.analyzer

**CloudFormation type:** `AWS::AccessAnalyzer::Analyzer`

The AWS::AccessAnalyzer::Analyzer type specifies an analyzer of the user's account

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AccessAnalyzer::Analyzer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalyzerConfiguration` | analyzer_configuration | `map` | optional, computed, provider-chosen |  | The configuration for the analyzer |
| `AnalyzerName` | analyzer_name | `string` | optional, computed, provider-chosen, replaces on change |  | Analyzer name |
| `ArchiveRules` | archive_rules | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the analyzer |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the analyzer, must be one of ACCOUNT, ORGANIZATION, ACCOUNT_INTERNAL_ACCESS, ORGANIZATION_INTERNAL_ACCESS, ACCOUNT_UNUSED_ACCESS and ORGANIZATION_UNUSED_ACCESS |

Supports update: yes

Discovery: supported
