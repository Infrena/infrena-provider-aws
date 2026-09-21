# aws.appstream.application

**CloudFormation type:** `AWS::AppStream::Application`

Resource Type definition for AWS::AppStream::Application

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AppStream::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppBlockArn` | app_block_arn | `string` | required | aws.appblock.Arn |  |
| `Arn` |  | `string` | computed |  |  |
| `AttributesToDelete` | attributes_to_delete | `list` | optional, computed, provider-chosen, write-only |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `IconS3Location` | icon_s3_location | `map` | required |  |  |
| `InstanceFamilies` | instance_families | `list` | required, replaces on change |  |  |
| `LaunchParameters` | launch_parameters | `string` | optional, computed, provider-chosen |  |  |
| `LaunchPath` | launch_path | `string` | required |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Platforms` |  | `list` | required, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |
| `WorkingDirectory` | working_directory | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: not supported
