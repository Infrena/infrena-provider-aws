# aws.archiverule

**CloudFormation type:** `AWS::AccessAnalyzer::ArchiveRule`

Creates an archive rule for the specified analyzer. Archive rules automatically archive new findings that meet the criteria you define when you create the rule.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AccessAnalyzer::ArchiveRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalyzerName` | analyzer_name | `string` | required, replaces on change |  | The name of the analyzer for the archive rule. |
| `Arn` |  | `string` | computed |  | The ARN of the archive rule. |
| `CreatedAt` | created_at | `string` | computed |  | The time at which the archive rule was created. |
| `Filter` |  | `map` | required |  | The criteria for the archive rule. A map of filter criteria property names to their criterion values. |
| `RuleName` | rule_name | `string` | required, replaces on change |  | The name of the archive rule. |
| `UpdatedAt` | updated_at | `string` | computed |  | The time at which the archive rule was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
