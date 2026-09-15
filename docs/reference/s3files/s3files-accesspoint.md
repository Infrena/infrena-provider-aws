# aws.s3files.accesspoint

**CloudFormation type:** `AWS::S3Files::AccessPoint`

Resource Type definition for AWS::S3Files::AccessPoint

Region attribute: `region`

**Import ID:** `<region>/AccessPointId` (AWS::S3Files::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPointArn` | access_point_arn | `string` | computed |  |  |
| `AccessPointId` | access_point_id | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change |  | (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation. |
| `FileSystemId` | file_system_id | `string` | required, replaces on change | aws.s3files.filesystem.FileSystemId | The ID of the S3 Files file system that the access point provides access to. |
| `OwnerId` | owner_id | `string` | computed |  |  |
| `PosixUser` | posix_user | `map` | optional, computed, provider-chosen, replaces on change |  | The operating system user and group applied to all compute drive requests made using the access point. |
| `RootDirectory` | root_directory | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the directory on the Amazon EFS file system that the access point exposes as the root directory of your file system to NFS clients using the access point. The clients using the access point can only access the root directory and below. If the RootDirectory>Path specified does not exist, EFS creates it and applies the CreationPermissions settings when a client connects to an access point. When specifying a RootDirectory, you need to provide the Path, and the CreationPermissions is optional. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
