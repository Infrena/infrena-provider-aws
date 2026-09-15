# aws.accesstoken

**CloudFormation type:** `AWS::Route53GlobalResolver::AccessToken`

Resource schema for AWS::Route53GlobalResolver::AccessToken

Region attribute: `region`

**Import ID:** `<region>/AccessTokenId` (AWS::Route53GlobalResolver::AccessToken)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessTokenId` | access_token_id | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DnsViewId` | dns_view_id | `string` | required, replaces on change | aws.dnsview.DnsViewId |  |
| `ExpiresAt` | expires_at | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `GlobalResolverId` | global_resolver_id | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |
| `Value` |  | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
