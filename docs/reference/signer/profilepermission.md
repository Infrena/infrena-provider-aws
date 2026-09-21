# aws.profilepermission

**CloudFormation type:** `AWS::Signer::ProfilePermission`

Resource Type definition for AWS::Signer::ProfilePermission

Region attribute: `region`

**Import ID:** `<region>/StatementId|ProfileName` (AWS::Signer::ProfilePermission)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | required, replaces on change |  |  |
| `Principal` |  | `string` | required, replaces on change |  |  |
| `ProfileName` | profile_name | `string` | required, replaces on change |  |  |
| `ProfileVersion` | profile_version | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `StatementId` | statement_id | `string` | required, replaces on change |  |  |

Supports update: no

Discovery: supported
