# aws.amplify.app

**CloudFormation type:** `AWS::Amplify::App`

The AWS::Amplify::App resource creates Apps in the Amplify Console. An App is a collection of branches.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Amplify::App)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessToken` | access_token | `string` | optional, computed, provider-chosen, sensitive, write-only |  |  |
| `AppId` | app_id | `string` | computed |  |  |
| `AppName` | app_name | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AutoBranchCreationConfig` | auto_branch_creation_config | `map` | optional, computed, provider-chosen, write-only |  |  |
| `BasicAuthConfig` | basic_auth_config | `map` | optional, computed, provider-chosen, write-only |  |  |
| `BuildSpec` | build_spec | `string` | optional, computed, provider-chosen |  |  |
| `CacheConfig` | cache_config | `map` | optional, computed, provider-chosen |  |  |
| `ComputeRoleArn` | compute_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `CustomHeaders` | custom_headers | `string` | optional, computed, provider-chosen |  |  |
| `CustomRules` | custom_rules | `list` | optional, computed, provider-chosen |  |  |
| `DefaultDomain` | default_domain | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnableBranchAutoDeletion` | enable_branch_auto_deletion | `boolean` | optional, computed, provider-chosen |  |  |
| `EnvironmentVariables` | environment_variables | `list` | optional, computed, provider-chosen |  |  |
| `IAMServiceRole` | iam_service_role | `string` | optional, computed, provider-chosen |  |  |
| `JobConfig` | job_config | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `OauthToken` | oauth_token | `string` | optional, computed, provider-chosen, sensitive, write-only |  |  |
| `Platform` |  | `string` | optional, computed, provider-chosen |  |  |
| `Repository` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
