# aws.imagepipeline

**CloudFormation type:** `AWS::ImageBuilder::ImagePipeline`

Resource Type definition for AWS::ImageBuilder::ImagePipeline

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::ImagePipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the image pipeline. |
| `ContainerRecipeArn` | container_recipe_arn | `string` | optional, computed, provider-chosen | aws.containerrecipe.Arn | The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested. |
| `DeploymentId` | deployment_id | `string` | computed |  | The deployment ID of the pipeline, used for resource create/update triggers. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the image pipeline. |
| `DistributionConfigurationArn` | distribution_configuration_arn | `string` | optional, computed, provider-chosen | aws.distributionconfiguration.Arn | The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline. |
| `EnhancedImageMetadataEnabled` | enhanced_image_metadata_enabled | `boolean` | optional, computed, provider-chosen |  | Collects additional information about the image being created, including the operating system (OS) version and package list. |
| `ExecutionRole` | execution_role | `string` | optional, computed, provider-chosen |  | The execution role name/ARN for the image build, if provided |
| `ImageRecipeArn` | image_recipe_arn | `string` | optional, computed, provider-chosen | aws.imagerecipe.Arn | The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed. |
| `ImageScanningConfiguration` | image_scanning_configuration | `map` | optional, computed, provider-chosen |  | Determines if tests should run after building the image. Image Builder defaults to enable tests to run following the image build, before image distribution. |
| `ImageTags` | image_tags | `map` | optional, computed, provider-chosen |  | The tags to be applied to images created by this pipeline. |
| `ImageTestsConfiguration` | image_tests_configuration | `map` | optional, computed, provider-chosen |  | Image tests configuration. |
| `InfrastructureConfigurationArn` | infrastructure_configuration_arn | `string` | optional, computed, provider-chosen | aws.infrastructureconfiguration.Arn | The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline. |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  | The logging configuration settings for the image pipeline. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the image pipeline. |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  | The schedule of the image pipeline. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the image pipeline. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags of this image pipeline. |
| `Workflows` |  | `list` | optional, computed, provider-chosen |  | Workflows to define the image build process |

Supports update: yes

Discovery: supported
