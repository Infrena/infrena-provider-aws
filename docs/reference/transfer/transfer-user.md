# aws.transfer.user

**CloudFormation type:** `AWS::Transfer::User`

Definition of AWS::Transfer::User Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Transfer::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `HomeDirectory` | home_directory | `string` | optional, computed, provider-chosen |  |  |
| `HomeDirectoryMappings` | home_directory_mappings | `list` | optional, computed, provider-chosen |  |  |
| `HomeDirectoryType` | home_directory_type | `string` | optional, computed, provider-chosen |  |  |
| `Policy` |  | `string` | optional, computed, provider-chosen |  |  |
| `PosixProfile` | posix_profile | `map` | optional, computed, provider-chosen |  |  |
| `Role` |  | `string` | required |  |  |
| `ServerId` | server_id | `string` | required, replaces on change | aws.server.ServerId |  |
| `SshPublicKeys` | ssh_public_keys | `list` | optional, computed, provider-chosen |  | This represents the SSH User Public Keys for CloudFormation resource |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UserName` | user_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
