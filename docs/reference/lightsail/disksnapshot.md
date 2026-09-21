# aws.disksnapshot

**CloudFormation type:** `AWS::Lightsail::DiskSnapshot`

Resource Type definition for AWS::Lightsail::DiskSnapshot

Region attribute: `region`

**Import ID:** `<region>/DiskSnapshotName` (AWS::Lightsail::DiskSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the disk snapshot was created. |
| `DiskName` | disk_name | `string` | required, replaces on change |  | The name of the source disk from which the snapshot was created. |
| `DiskSnapshotArn` | disk_snapshot_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the disk snapshot. |
| `DiskSnapshotName` | disk_snapshot_name | `string` | required, replaces on change |  | The name of the disk snapshot (e.g., my-disk-snapshot). |
| `FromDiskName` | from_disk_name | `string` | computed |  | The name of the source disk from which the disk snapshot was created. |
| `IsFromAutoSnapshot` | is_from_auto_snapshot | `boolean` | computed |  | A Boolean value indicating whether the snapshot was created from an automatic snapshot. |
| `Location` |  | `map` | computed |  | The AWS Region and Availability Zone where the disk snapshot was created. |
| `Progress` |  | `string` | computed |  | The progress of the disk snapshot creation operation. |
| `ResourceType` | resource_type | `string` | computed |  | The Lightsail resource type (DiskSnapshot). |
| `SizeInGb` | size_in_gb | `integer` | computed |  | The size of the disk snapshot in GB. |
| `State` |  | `string` | computed |  | The status of the disk snapshot operation. |
| `SupportCode` | support_code | `string` | computed |  | The support code. Include this code in your email to support when you have questions about an instance or another resource in Lightsail. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
