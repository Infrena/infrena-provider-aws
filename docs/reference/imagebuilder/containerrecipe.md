# aws.containerrecipe

**CloudFormation type:** `AWS::ImageBuilder::ContainerRecipe`

Resource Type definition for AWS::ImageBuilder::ContainerRecipe

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::ContainerRecipe)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the container recipe. |
| `Components` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Components for build and test that are included in the container recipe. |
| `ContainerType` | container_type | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the type of container, such as Docker. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the container recipe. |
| `DockerfileTemplateData` | dockerfile_template_data | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe. |
| `DockerfileTemplateUri` | dockerfile_template_uri | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The S3 URI for the Dockerfile that will be used to build your container image. |
| `ImageOsVersionOverride` | image_os_version_override | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the operating system version for the source image. |
| `InstanceConfiguration` | instance_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | A group of options that can be used to configure an instance for building and testing container images. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | Identifies which KMS key is used to encrypt the container image. |
| `LatestVersion` | latest_version | `map` | computed |  | The latest version references of the container recipe. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the container recipe. |
| `ParentImage` | parent_image | `string` | optional, computed, provider-chosen, replaces on change |  | The source image for the container recipe. |
| `PlatformOverride` | platform_override | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the operating system platform when you use a custom source image. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Tags that are attached to the container recipe. |
| `TargetRepository` | target_repository | `map` | optional, computed, provider-chosen, replaces on change |  | The container repository where the output container image is stored. |
| `Version` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The semantic version of the container recipe (<major>.<minor>.<patch>). |
| `WorkingDirectory` | working_directory | `string` | optional, computed, provider-chosen, replaces on change |  | The working directory to be used during build and test workflows. |

Supports update: yes

Discovery: supported
