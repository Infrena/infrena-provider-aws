# aws.imagerecipe

**CloudFormation type:** `AWS::ImageBuilder::ImageRecipe`

Resource Type definition for AWS::ImageBuilder::ImageRecipe

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::ImageRecipe)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalInstanceConfiguration` | additional_instance_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Specify additional settings and launch scripts for your build instances. |
| `AmiTags` | ami_tags | `map` | optional, computed, provider-chosen, replaces on change |  | The tags to apply to the AMI created by this image recipe. |
| `AmiWatermarks` | ami_watermarks | `list` | optional, computed, provider-chosen, replaces on change |  | The AMI watermark names to attach to the output AMI from this recipe. AMI watermarks are lineage markers that automatically propagate to derivative AMIs when the source AMI is copied or distributed. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the image recipe. |
| `BlockDeviceMappings` | block_device_mappings | `list` | optional, computed, provider-chosen, replaces on change |  | The block device mappings to apply when creating images from this recipe. |
| `Components` |  | `list` | optional, computed, provider-chosen, replaces on change |  | The components of the image recipe. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the image recipe. |
| `LatestVersion` | latest_version | `map` | computed |  | The latest version references of the image recipe. |
| `Name` |  | `string` | required, replaces on change |  | The name of the image recipe. |
| `ParentImage` | parent_image | `string` | required, replaces on change |  | The parent image of the image recipe. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags of the image recipe. |
| `Version` |  | `string` | required, replaces on change |  | The version of the image recipe. |
| `WorkingDirectory` | working_directory | `string` | optional, computed, provider-chosen, replaces on change |  | The working directory to be used during build and test workflows. |

Supports update: yes

Discovery: supported
