# aws.networkinsightsanalysis

**CloudFormation type:** `AWS::EC2::NetworkInsightsAnalysis`

Resource schema for AWS::EC2::NetworkInsightsAnalysis

Region attribute: `region`

**Import ID:** `<region>/NetworkInsightsAnalysisId` (AWS::EC2::NetworkInsightsAnalysis)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalAccounts` | additional_accounts | `list` | optional, computed, provider-chosen |  |  |
| `AlternatePathHints` | alternate_path_hints | `list` | computed |  |  |
| `Explanations` |  | `list` | computed |  |  |
| `FilterInArns` | filter_in_arns | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `FilterOutArns` | filter_out_arns | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `ForwardPathComponents` | forward_path_components | `list` | computed |  |  |
| `NetworkInsightsAnalysisArn` | network_insights_analysis_arn | `string` | computed |  |  |
| `NetworkInsightsAnalysisId` | network_insights_analysis_id | `string` | computed |  |  |
| `NetworkInsightsPathId` | network_insights_path_id | `string` | required, replaces on change | aws.networkinsightspath.NetworkInsightsPathId |  |
| `NetworkPathFound` | network_path_found | `boolean` | computed |  |  |
| `ReturnPathComponents` | return_path_components | `list` | computed |  |  |
| `StartDate` | start_date | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `SuggestedAccounts` | suggested_accounts | `list` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
