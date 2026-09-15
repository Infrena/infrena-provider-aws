# aws.efs.accesspoint

**CloudFormation type:** `AWS::EFS::AccessPoint`

The ``AWS::EFS::AccessPoint`` resource creates an EFS access point. An access point is an application-specific view into an EFS file system that applies an operating system user and group, and a file system path, to any file system request made through the access point. The operating system user and group override any identity information provided by the NFS client. The file system path is exposed as the access point's root directory. Applications using the access point can only access data in its own directory and below. To learn more, see [Mounting a file system using EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).

Region attribute: `region`

**Import ID:** `<region>/AccessPointId` (AWS::EFS::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPointId` | access_point_id | `string` | computed |  |  |
| `AccessPointTags` | access_point_tags | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Arn` |  | `string` | computed |  |  |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change |  | The opaque string specified in the request to ensure idempotent creation. |
| `FileSystemId` | file_system_id | `string` | required, replaces on change | aws.efs.filesystem.FileSystemId | The ID of the EFS file system that the access point applies to. Accepts only the ID format for input when specifying a file system, for example ``fs-0123456789abcedf2``. |
| `PosixUser` | posix_user | `map` | optional, computed, provider-chosen, replaces on change |  | The full POSIX identity, including the user ID, group ID, and any secondary group IDs, on the access point that is used for all file system operations performed by NFS clients using the access point. |
| `RootDirectory` | root_directory | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the directory on the Amazon EFS file system that the access point provides access to. The access point exposes the specified file system path as the root directory of your file system to applications using the access point. NFS clients using the access point can only access data in the access point's ``RootDirectory`` and its subdirectories. |

Supports update: yes

Discovery: supported
