# aws.acl

**CloudFormation type:** `AWS::MemoryDB::ACL`

Resource Type definition for AWS::MemoryDB::ACL

Region attribute: `region`

**Import ID:** `<region>/ACLName` (AWS::MemoryDB::ACL)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ACLName` | acl_name | `string` | required, replaces on change |  | The name of the acl. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the acl. |
| `Status` |  | `string` | computed |  | Indicates acl status. Can be "creating", "active", "modifying", "deleting". |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this cluster. |
| `UserNames` | user_names | `list` | optional, computed, provider-chosen |  | List of users associated to this acl. |

Supports update: yes

Discovery: supported
