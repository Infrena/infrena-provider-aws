# aws.networkinsightsaccessscopeanalysis

**CloudFormation type:** `AWS::EC2::NetworkInsightsAccessScopeAnalysis`

Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis

Region attribute: `region`

**Import ID:** `<region>/NetworkInsightsAccessScopeAnalysisId` (AWS::EC2::NetworkInsightsAccessScopeAnalysis)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalyzedEniCount` | analyzed_eni_count | `integer` | computed |  |  |
| `EndDate` | end_date | `string` | computed |  |  |
| `FindingsFound` | findings_found | `string` | computed |  |  |
| `NetworkInsightsAccessScopeAnalysisArn` | network_insights_access_scope_analysis_arn | `string` | computed |  |  |
| `NetworkInsightsAccessScopeAnalysisId` | network_insights_access_scope_analysis_id | `string` | computed |  |  |
| `NetworkInsightsAccessScopeId` | network_insights_access_scope_id | `string` | required, replaces on change | aws.networkinsightsaccessscope.NetworkInsightsAccessScopeId |  |
| `StartDate` | start_date | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
