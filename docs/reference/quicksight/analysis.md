# aws.analysis

**CloudFormation type:** `AWS::QuickSight::Analysis`

Definition of the AWS::QuickSight::Analysis Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AnalysisId|AwsAccountId` (AWS::QuickSight::Analysis)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalysisId` | analysis_id | `string` | required, replaces on change | aws.analysis.AnalysisId |  |
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the analysis.</p> |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that the analysis was created.</p> |
| `DataSetArns` | data_set_arns | `list` | computed |  | <p>The ARNs of the datasets of the analysis.</p> |
| `Definition` |  | `map` | optional, computed, provider-chosen, write-only |  |  |
| `Errors` |  | `list` | optional, computed, provider-chosen |  | <p>Errors associated with the analysis.</p> |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, write-only | aws.folder.Arn |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The time that the analysis was last updated.</p> |
| `Name` |  | `string` | required |  | <p>The descriptive name of the analysis.</p> |
| `Parameters` |  | `map` | optional, computed, provider-chosen, write-only |  | <p>A list of Amazon QuickSight parameters and the list's override values.</p> |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `Sheets` |  | `list` | optional, computed, provider-chosen |  | <p>A list of the associated sheets with the unique identifier and name of each sheet.</p> |
| `SourceEntity` | source_entity | `map` | optional, computed, provider-chosen, write-only |  | <p>The source entity of an analysis.</p> |
| `Status` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `ThemeArn` | theme_arn | `string` | optional, computed, provider-chosen | aws.quicksight.theme.Arn | <p>The ARN of the theme of the analysis.</p> |
| `ValidationStrategy` | validation_strategy | `map` | optional, computed, provider-chosen, write-only |  | <p>The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to <code>LENIENT</code>, validation is skipped for specific errors.</p> |

Supports update: yes

Discovery: supported (parent resource required)
