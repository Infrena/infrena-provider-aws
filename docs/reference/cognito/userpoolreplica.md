# aws.userpoolreplica

**CloudFormation type:** `AWS::Cognito::UserPoolReplica`

Resource Type definition for AWS::Cognito::UserPoolReplica

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|RegionName` (AWS::Cognito::UserPoolReplica)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RegionName` | region_name | `string` | required, replaces on change |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |
| `UserPoolTagsAtCreate` | user_pool_tags_at_create | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
