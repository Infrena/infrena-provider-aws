# aws.sagemaker.image

**CloudFormation type:** `AWS::SageMaker::Image`

Resource Type definition for AWS::SageMaker::Image

Region attribute: `region`

**Import ID:** `<region>/ImageArn` (AWS::SageMaker::Image)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ImageArn` | image_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the image. |
| `ImageDescription` | image_description | `string` | optional, computed, provider-chosen |  | A description of the image. |
| `ImageDisplayName` | image_display_name | `string` | optional, computed, provider-chosen |  | The display name of the image. |
| `ImageName` | image_name | `string` | required, replaces on change |  | The name of the image. |
| `ImageRoleArn` | image_role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
