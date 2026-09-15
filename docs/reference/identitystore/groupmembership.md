# aws.groupmembership

**CloudFormation type:** `AWS::IdentityStore::GroupMembership`

Resource Type Definition for AWS:IdentityStore::GroupMembership

Region attribute: `region`

**Import ID:** `<region>/MembershipId|IdentityStoreId` (AWS::IdentityStore::GroupMembership)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupId` | group_id | `string` | required, replaces on change | aws.identitystore.group.GroupId | The unique identifier for a group in the identity store. |
| `IdentityStoreId` | identity_store_id | `string` | required, replaces on change |  | The globally unique identifier for the identity store. |
| `MemberId` | member_id | `map` | required, replaces on change |  | An object containing the identifier of a group member. |
| `MembershipId` | membership_id | `string` | computed |  | The identifier for a GroupMembership in the identity store. |

Supports update: no

Discovery: supported (parent resource required)
