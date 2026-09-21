# aws.locationfsxopenzfs

**CloudFormation type:** `AWS::DataSync::LocationFSxOpenZFS`

Resource schema for AWS::DataSync::LocationFSxOpenZFS.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationFSxOpenZFS)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FsxFilesystemArn` | fsx_filesystem_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) for the FSx OpenZFS file system. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the FSx OpenZFS that was described. |
| `Protocol` |  | `map` | required |  | Configuration settings for an NFS or SMB protocol, currently only support NFS |
| `SecurityGroupArns` | security_group_arns | `list` | required, replaces on change |  | The ARNs of the security groups that are to use to configure the FSx OpenZFS file system. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the location's path. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
