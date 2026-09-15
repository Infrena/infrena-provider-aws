# aws.memorydb.user

**CloudFormation type:** `AWS::MemoryDB::User`

Resource Type definition for AWS::MemoryDB::User

Region attribute: `region`

**Import ID:** `<region>/UserName` (AWS::MemoryDB::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessString` | access_string | `string` | optional, computed, provider-chosen, write-only |  | Access permissions string used for this user account. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the user account. |
| `AuthenticationMode` | authentication_mode | `map` | optional, computed, provider-chosen, write-only |  |  |
| `Status` |  | `string` | computed |  | Indicates the user status. Can be "active", "modifying" or "deleting". |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this user. |
| `UserName` | user_name | `string` | required, replaces on change |  | The name of the user. |

Supports update: yes

Discovery: supported
