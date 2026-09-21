# aws.locationfsxontap

**CloudFormation type:** `AWS::DataSync::LocationFSxONTAP`

Resource Type definition for AWS::DataSync::LocationFSxONTAP.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationFSxONTAP)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FsxFilesystemArn` | fsx_filesystem_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the FSx ONAP file system. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the FSx ONTAP file system that was described. |
| `Protocol` |  | `map` | optional, computed, provider-chosen |  | Configuration settings for NFS or SMB protocol. |
| `SecurityGroupArns` | security_group_arns | `list` | required, replaces on change |  | The ARNs of the security groups that are to use to configure the FSx ONTAP file system. |
| `StorageVirtualMachineArn` | storage_virtual_machine_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) for the FSx ONTAP SVM. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the location's path. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
