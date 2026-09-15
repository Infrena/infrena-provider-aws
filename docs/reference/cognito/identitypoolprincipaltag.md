# aws.identitypoolprincipaltag

**CloudFormation type:** `AWS::Cognito::IdentityPoolPrincipalTag`

Resource Type definition for AWS::Cognito::IdentityPoolPrincipalTag

Region attribute: `region`

**Import ID:** `<region>/IdentityPoolId|IdentityProviderName` (AWS::Cognito::IdentityPoolPrincipalTag)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IdentityPoolId` | identity_pool_id | `string` | required, replaces on change | aws.identitypool.Id |  |
| `IdentityProviderName` | identity_provider_name | `string` | required, replaces on change |  |  |
| `PrincipalTags` | principal_tags | `map` | optional, computed, provider-chosen |  |  |
| `UseDefaults` | use_defaults | `boolean` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
