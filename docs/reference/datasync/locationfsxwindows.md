# aws.locationfsxwindows

**CloudFormation type:** `AWS::DataSync::LocationFSxWindows`

Resource Type definition for AWS::DataSync::LocationFSxWindows.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationFSxWindows)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CmkSecretConfig` | cmk_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key. |
| `CustomSecretConfig` | custom_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret. |
| `Domain` |  | `string` | optional, computed, provider-chosen |  | The name of the Windows domain that the FSx for Windows server belongs to. |
| `FsxFilesystemArn` | fsx_filesystem_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) for the FSx for Windows file system. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the FSx for Windows location that was described. |
| `ManagedSecretConfig` | managed_secret_config | `map` | computed |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager. |
| `Password` |  | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password of the user who has the permissions to access files and folders in the FSx for Windows file system. |
| `SecurityGroupArns` | security_group_arns | `list` | required, replaces on change |  | The ARNs of the security groups that are to use to configure the FSx for Windows file system. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | A subdirectory in the location's path. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `User` |  | `string` | required |  | The user who has the permissions to access files and folders in the FSx for Windows file system. |

Supports update: yes

Discovery: supported
