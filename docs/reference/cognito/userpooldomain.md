# aws.userpooldomain

**CloudFormation type:** `AWS::Cognito::UserPoolDomain`

Resource Type definition for AWS::Cognito::UserPoolDomain

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|Domain` (AWS::Cognito::UserPoolDomain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudFrontDistribution` | cloud_front_distribution | `string` | computed |  |  |
| `CustomDomainConfig` | custom_domain_config | `map` | optional, computed, provider-chosen |  |  |
| `Domain` |  | `string` | required, replaces on change |  |  |
| `ManagedLoginVersion` | managed_login_version | `integer` | optional, computed, provider-chosen, write-only |  |  |
| `Routing` |  | `map` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: not supported
