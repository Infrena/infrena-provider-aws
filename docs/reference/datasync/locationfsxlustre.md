# aws.locationfsxlustre

**CloudFormation type:** `AWS::DataSync::LocationFSxLustre`

Resource schema for AWS::DataSync::LocationFSxLustre.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationFSxLustre)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FsxFilesystemArn` | fsx_filesystem_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) for the FSx for Lustre file system. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the FSx for Lustre location that was described. |
| `SecurityGroupArns` | security_group_arns | `list` | required, replaces on change |  | The ARNs of the security groups that are to use to configure the FSx for Lustre file system. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the location's path. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
