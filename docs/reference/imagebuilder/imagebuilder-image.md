# aws.imagebuilder.image

**CloudFormation type:** `AWS::ImageBuilder::Image`

Resource Type definition for AWS::ImageBuilder::Image

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::Image)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the image. |
| `ContainerRecipeArn` | container_recipe_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.containerrecipe.Arn | The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested. |
| `DeletionSettings` | deletion_settings | `map` | optional, computed, provider-chosen, write-only |  | The deletion settings of the image, indicating whether to delete the underlying resources in addition to the image. |
| `DistributionConfigurationArn` | distribution_configuration_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.distributionconfiguration.Arn | The Amazon Resource Name (ARN) of the distribution configuration. |
| `EnhancedImageMetadataEnabled` | enhanced_image_metadata_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Collects additional information about the image being created, including the operating system (OS) version and package list. |
| `ExecutionRole` | execution_role | `string` | optional, computed, provider-chosen |  | The execution role name/ARN for the image build, if provided |
| `ImageId` | image_id | `string` | computed |  | The AMI ID of the EC2 AMI in current region. |
| `ImagePipelineExecutionSettings` | image_pipeline_execution_settings | `map` | optional, computed, provider-chosen, write-only |  | The settings for starting an image pipeline execution. |
| `ImageRecipeArn` | image_recipe_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.imagerecipe.Arn | The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed. |
| `ImageScanningConfiguration` | image_scanning_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Contains settings for Image Builder image resource and container image scans. |
| `ImageTestsConfiguration` | image_tests_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The image tests configuration used when creating this image. |
| `ImageUri` | image_uri | `string` | computed |  | URI for containers created in current Region with default ECR image tag |
| `InfrastructureConfigurationArn` | infrastructure_configuration_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.infrastructureconfiguration.Arn | The Amazon Resource Name (ARN) of the infrastructure configuration. |
| `LatestVersion` | latest_version | `map` | computed |  | The latest version references of the image. |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  | The logging configuration settings for the image. |
| `Name` |  | `string` | computed |  | The name of the image. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the image. |
| `Workflows` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Workflows to define the image build process |

Supports update: yes

Discovery: supported
