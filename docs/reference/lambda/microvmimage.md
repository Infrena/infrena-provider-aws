# aws.microvmimage

**CloudFormation type:** `AWS::Lambda::MicrovmImage`

Resource Type definition for AWS::Lambda::MicrovmImage

Region attribute: `region`

**Import ID:** `<region>/ImageArn` (AWS::Lambda::MicrovmImage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalOsCapabilities` | additional_os_capabilities | `list` | required, write-only |  |  |
| `BaseImageArn` | base_image_arn | `string` | required, write-only |  | ARN of the base MicroVM image. |
| `BaseImageVersion` | base_image_version | `string` | required, write-only |  | Specific version of the base MicroVM image to use. |
| `BuildRoleArn` | build_role_arn | `string` | required, write-only | aws.role.Arn | ARN of the IAM build role. |
| `CodeArtifact` | code_artifact | `map` | required, write-only |  | Code artifact for the active MicroVM image. |
| `CpuConfigurations` | cpu_configurations | `list` | required, write-only |  |  |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the MicroVM image was created. |
| `Description` |  | `string` | required, write-only |  | Human-readable description of the MicroVM image and its purpose. |
| `EgressNetworkConnectors` | egress_network_connectors | `list` | required, write-only |  |  |
| `EnvironmentVariables` | environment_variables | `list` | required, write-only |  | Environment variables to set in the container during the snapshot build. |
| `Hooks` |  | `map` | required, write-only |  |  |
| `ImageArn` | image_arn | `string` | computed |  | ARN of the MicroVM image. |
| `LatestActiveImageVersion` | latest_active_image_version | `string` | computed |  | The latest active version of the MicroVM image. |
| `LatestFailedImageVersion` | latest_failed_image_version | `string` | computed |  | The latest failed version of the MicroVM image. |
| `Logging` |  | `map` | required, write-only |  | Configuration for MicroVM image logging. |
| `Name` |  | `string` | required, replaces on change |  | Unique name for the MicroVM image within the account. |
| `Resources` |  | `list` | required, write-only |  |  |
| `State` |  | `string` | computed |  | Current state of the MicroVM image. |
| `Tags` |  | `map` | tags map |  | Key-value pairs to associate with the MicroVM image for organization and management. |
| `UpdatedAt` | updated_at | `string` | computed |  | Timestamp when the MicroVM image was updated. |

Supports update: yes

Discovery: supported
