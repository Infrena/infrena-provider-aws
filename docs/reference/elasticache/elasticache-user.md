# aws.elasticache.user

**CloudFormation type:** `AWS::ElastiCache::User`

Resource Type definition for AWS::ElastiCache::User

Region attribute: `region`

**Import ID:** `<region>/UserId` (AWS::ElastiCache::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessString` | access_string | `string` | optional, computed, provider-chosen, write-only |  | Access permissions string used for this user account. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the user account. |
| `AuthenticationMode` | authentication_mode | `map` | optional, computed, provider-chosen, write-only |  |  |
| `Engine` |  | `string` | required |  | The target cache engine for the user. |
| `NoPasswordRequired` | no_password_required | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates a password is not required for this user account. |
| `Passwords` |  | `list` | optional, computed, provider-chosen, sensitive, write-only |  | Passwords used for this user account. You can create up to two passwords for each user. |
| `Status` |  | `string` | computed |  | Indicates the user status. Can be "active", "modifying" or "deleting". |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this user. |
| `UserId` | user_id | `string` | required, replaces on change | aws.elasticache.user.UserId | The ID of the user. |
| `UserName` | user_name | `string` | required, replaces on change |  | The username of the user. |

Supports update: yes

Discovery: supported
