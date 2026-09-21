# aws.networkinsightsaccessscope

**CloudFormation type:** `AWS::EC2::NetworkInsightsAccessScope`

Resource schema for AWS::EC2::NetworkInsightsAccessScope

Region attribute: `region`

**Import ID:** `<region>/NetworkInsightsAccessScopeId` (AWS::EC2::NetworkInsightsAccessScope)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedDate` | created_date | `string` | computed |  |  |
| `ExcludePaths` | exclude_paths | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `MatchPaths` | match_paths | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `NetworkInsightsAccessScopeArn` | network_insights_access_scope_arn | `string` | computed |  |  |
| `NetworkInsightsAccessScopeId` | network_insights_access_scope_id | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedDate` | updated_date | `string` | computed |  |  |

Supports update: yes

Discovery: supported
