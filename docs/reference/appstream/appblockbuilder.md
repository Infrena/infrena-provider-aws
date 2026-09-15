# aws.appblockbuilder

**CloudFormation type:** `AWS::AppStream::AppBlockBuilder`

Resource Type definition for AWS::AppStream::AppBlockBuilder.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::AppStream::AppBlockBuilder)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessEndpoints` | access_endpoints | `list` | optional, computed, provider-chosen |  |  |
| `AppBlockArns` | app_block_arns | `list` | optional, computed, provider-chosen, write-only | aws.appblock.Arn |  |
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `EnableDefaultInternetAccess` | enable_default_internet_access | `boolean` | optional, computed, provider-chosen |  |  |
| `IamRoleArn` | iam_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `InstanceType` | instance_type | `string` | required |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Platform` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcConfig` | vpc_config | `map` | required |  |  |

Supports update: yes

Discovery: supported
