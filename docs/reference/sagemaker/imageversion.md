# aws.imageversion

**CloudFormation type:** `AWS::SageMaker::ImageVersion`

Resource Type definition for AWS::SageMaker::ImageVersion

Region attribute: `region`

**Import ID:** `<region>/ImageVersionArn` (AWS::SageMaker::ImageVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | optional, computed, provider-chosen, write-only |  | The alias of the image version. |
| `Aliases` |  | `list` | optional, computed, provider-chosen, write-only |  | List of aliases for the image version. |
| `BaseImage` | base_image | `string` | required, replaces on change |  | The registry path of the container image on which this image version is based. |
| `ContainerImage` | container_image | `string` | computed |  | The registry path of the container image that contains this image version. |
| `Horovod` |  | `boolean` | optional, computed, provider-chosen |  | Indicates Horovod compatibility. |
| `ImageArn` | image_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the parent image. |
| `ImageName` | image_name | `string` | required, replaces on change |  | The name of the image this version belongs to. |
| `ImageVersionArn` | image_version_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the image version. |
| `JobType` | job_type | `string` | optional, computed, provider-chosen |  | Indicates SageMaker job type compatibility. |
| `MLFramework` | ml_framework | `string` | optional, computed, provider-chosen |  | The machine learning framework vended in the image version. |
| `Processor` |  | `string` | optional, computed, provider-chosen |  | Indicates CPU or GPU compatibility. |
| `ProgrammingLang` | programming_lang | `string` | optional, computed, provider-chosen |  | The supported programming language and its version. |
| `ReleaseNotes` | release_notes | `string` | optional, computed, provider-chosen |  | The maintainer description of the image version. |
| `VendorGuidance` | vendor_guidance | `string` | optional, computed, provider-chosen |  | The availability of the image version specified by the maintainer. |
| `Version` |  | `integer` | computed |  | The version number of the image version. |

Supports update: yes

Discovery: supported (parent resource required)
