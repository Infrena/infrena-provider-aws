# aws.configuredtable

**CloudFormation type:** `AWS::CleanRooms::ConfiguredTable`

Represents a table that can be associated with collaborations

Region attribute: `region`

**Import ID:** `<region>/ConfiguredTableIdentifier` (AWS::CleanRooms::ConfiguredTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedColumns` | allowed_columns | `list` | required |  |  |
| `AnalysisMethod` | analysis_method | `string` | required |  |  |
| `AnalysisRules` | analysis_rules | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `ConfiguredTableIdentifier` | configured_table_identifier | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `SelectedAnalysisMethods` | selected_analysis_methods | `list` | optional, computed, provider-chosen |  |  |
| `TableReference` | table_reference | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration. |

Supports update: yes

Discovery: supported
