# aws.terms

**CloudFormation type:** `AWS::Cognito::Terms`

Resource Type definition for AWS::Cognito::Terms

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|TermsId` (AWS::Cognito::Terms)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClientId` | client_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Enforcement` |  | `string` | required |  |  |
| `Links` |  | `map` | required |  |  |
| `TermsId` | terms_id | `string` | computed |  |  |
| `TermsName` | terms_name | `string` | required |  |  |
| `TermsSource` | terms_source | `string` | required |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: supported (parent resource required)
