# aws.appstream.entitlement

**CloudFormation type:** `AWS::AppStream::Entitlement`

Resource Type definition for AWS::AppStream::Entitlement

Region attribute: `region`

**Import ID:** `<region>/StackName|Name` (AWS::AppStream::Entitlement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppVisibility` | app_visibility | `string` | required |  |  |
| `Attributes` |  | `list` | required |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `StackName` | stack_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: not supported
