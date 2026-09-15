# aws.locationefs

**CloudFormation type:** `AWS::DataSync::LocationEFS`

Resource schema for AWS::DataSync::LocationEFS.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationEFS)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPointArn` | access_point_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) for the Amazon EFS Access point that DataSync uses when accessing the EFS file system. |
| `Ec2Config` | ec2_config | `map` | required, replaces on change |  | The subnet and security group that DataSync uses to access target EFS file system. |
| `EfsFilesystemArn` | efs_filesystem_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) for the Amazon EFS file system. |
| `FileSystemAccessRoleArn` | file_system_access_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the AWS IAM role that the DataSync will assume when mounting the EFS file system. |
| `InTransitEncryption` | in_transit_encryption | `string` | optional, computed, provider-chosen |  | Protocol that is used for encrypting the traffic exchanged between the DataSync Agent and the EFS file system. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the EFS location that was described. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
