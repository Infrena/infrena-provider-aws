# aws.branch

**CloudFormation type:** `AWS::Amplify::Branch`

The AWS::Amplify::Branch resource creates a new branch within an app.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Amplify::Branch)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppId` | app_id | `string` | required, replaces on change | aws.amplify.app.AppId |  |
| `Arn` |  | `string` | computed |  |  |
| `Backend` |  | `map` | optional, computed, provider-chosen |  |  |
| `BasicAuthConfig` | basic_auth_config | `map` | optional, computed, provider-chosen, write-only |  |  |
| `BranchName` | branch_name | `string` | required, replaces on change |  |  |
| `BuildSpec` | build_spec | `string` | optional, computed, provider-chosen |  |  |
| `ComputeRoleArn` | compute_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnableAutoBuild` | enable_auto_build | `boolean` | optional, computed, provider-chosen |  |  |
| `EnablePerformanceMode` | enable_performance_mode | `boolean` | optional, computed, provider-chosen |  |  |
| `EnablePullRequestPreview` | enable_pull_request_preview | `boolean` | optional, computed, provider-chosen |  |  |
| `EnableSkewProtection` | enable_skew_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `EnvironmentVariables` | environment_variables | `list` | optional, computed, provider-chosen |  |  |
| `Framework` |  | `string` | optional, computed, provider-chosen |  |  |
| `PullRequestEnvironmentName` | pull_request_environment_name | `string` | optional, computed, provider-chosen |  |  |
| `Stage` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
