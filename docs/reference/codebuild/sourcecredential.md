# aws.sourcecredential

**CloudFormation type:** `AWS::CodeBuild::SourceCredential`

Resource Type definition for AWS::CodeBuild::SourceCredential

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeBuild::SourceCredential)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the SourceCredential resource. |
| `AuthType` | auth_type | `string` | required |  | The type of authentication used by the credentials. Valid options are OAUTH, BASIC_AUTH, PERSONAL_ACCESS_TOKEN, CODECONNECTIONS, or SECRETS_MANAGER. |
| `ServerType` | server_type | `string` | required, replaces on change |  | The type of source provider. The valid options are GITHUB, GITHUB_ENTERPRISE, GITLAB, GITLAB_SELF_MANAGED, or BITBUCKET. |
| `Token` |  | `string` | required, sensitive, write-only |  | For GitHub or GitHub Enterprise, this is the personal access token. For Bitbucket, this is either the access token or the app password. For the authType CODECONNECTIONS, this is the connectionArn. For the authType SECRETS_MANAGER, this is the secretArn. |
| `Username` |  | `string` | optional, computed, provider-chosen, write-only |  | The Bitbucket username when the authType is BASIC_AUTH. This parameter is not valid for other types of source providers or connections. |

Supports update: yes

Discovery: supported
