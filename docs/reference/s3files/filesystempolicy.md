# aws.filesystempolicy

**CloudFormation type:** `AWS::S3Files::FileSystemPolicy`

Resource Type definition for AWS::S3Files::FileSystemPolicy

Region attribute: `region`

**Import ID:** `<region>/FileSystemId` (AWS::S3Files::FileSystemPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FileSystemId` | file_system_id | `string` | required, replaces on change | aws.s3files.filesystem.FileSystemId |  |
| `Policy` |  | `map` | required |  |  |

Supports update: yes

Discovery: supported
