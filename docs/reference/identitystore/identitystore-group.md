# aws.identitystore.group

**CloudFormation type:** `AWS::IdentityStore::Group`

Resource Type definition for AWS::IdentityStore::Group

Region attribute: `region`

**Import ID:** `<region>/GroupId|IdentityStoreId` (AWS::IdentityStore::Group)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A string containing the description of the group. |
| `DisplayName` | display_name | `string` | required |  | A string containing the name of the group. This value is commonly displayed when the group is referenced. |
| `GroupId` | group_id | `string` | computed |  | The unique identifier for a group in the identity store. |
| `IdentityStoreId` | identity_store_id | `string` | required, replaces on change |  | The globally unique identifier for the identity store. |

Supports update: yes

Discovery: supported (parent resource required)
