# aws.identityprovider

**CloudFormation type:** `AWS::WorkSpacesWeb::IdentityProvider`

Definition of AWS::WorkSpacesWeb::IdentityProvider Resource Type

Region attribute: `region`

**Import ID:** `<region>/IdentityProviderArn` (AWS::WorkSpacesWeb::IdentityProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IdentityProviderArn` | identity_provider_arn | `string` | computed |  |  |
| `IdentityProviderDetails` | identity_provider_details | `map` | required |  |  |
| `IdentityProviderName` | identity_provider_name | `string` | required |  |  |
| `IdentityProviderType` | identity_provider_type | `string` | required |  |  |
| `PortalArn` | portal_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.workspacesweb.portal.PortalArn |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
