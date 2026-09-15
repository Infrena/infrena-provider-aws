# aws.imagebuilder

**CloudFormation type:** `AWS::AppStream::ImageBuilder`

Resource Type definition for AWS::AppStream::ImageBuilder

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::AppStream::ImageBuilder)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessEndpoints` | access_endpoints | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `AppstreamAgentVersion` | appstream_agent_version | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainJoinInfo` | domain_join_info | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `EnableDefaultInternetAccess` | enable_default_internet_access | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `IamRoleArn` | iam_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn |  |
| `ImageArn` | image_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ImageName` | image_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `InstanceType` | instance_type | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RootVolumeConfig` | root_volume_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SoftwaresToInstall` | softwares_to_install | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `SoftwaresToUninstall` | softwares_to_uninstall | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `StreamingUrl` | streaming_url | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported
