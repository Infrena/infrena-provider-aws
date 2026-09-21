# aws.usergroup

**CloudFormation type:** `AWS::ElastiCache::UserGroup`

Resource Type definition for AWS::ElastiCache::UserGroup

Region attribute: `region`

**Import ID:** `<region>/UserGroupId` (AWS::ElastiCache::UserGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the user account. |
| `Engine` |  | `string` | required |  | The target cache engine for the user group. |
| `Status` |  | `string` | computed |  | Indicates user group status. Can be "creating", "active", "modifying", "deleting". |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this user. |
| `UserGroupId` | user_group_id | `string` | required, replaces on change | aws.usergroup.UserGroupId | The ID of the user group. |
| `UserIds` | user_ids | `list` | required | aws.elasticache.user.UserId | List of users associated to this user group. |

Supports update: yes

Discovery: supported
