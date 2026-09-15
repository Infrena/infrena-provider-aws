# aws.sagemaker.space

**CloudFormation type:** `AWS::SageMaker::Space`

Resource Type definition for AWS::SageMaker::Space

Region attribute: `region`

**Import ID:** `<region>/DomainId|SpaceName` (AWS::SageMaker::Space)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainId` | domain_id | `string` | required, replaces on change | aws.sagemaker.domain.DomainId | The ID of the associated Domain. |
| `OwnershipSettings` | ownership_settings | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SpaceArn` | space_arn | `string` | computed |  | The space Amazon Resource Name (ARN). |
| `SpaceDisplayName` | space_display_name | `string` | optional, computed, provider-chosen |  |  |
| `SpaceName` | space_name | `string` | required, replaces on change |  | A name for the Space. |
| `SpaceSettings` | space_settings | `map` | optional, computed, provider-chosen, write-only |  | A collection of settings that apply to spaces of Amazon SageMaker Studio. These settings are specified when the CreateSpace API is called. |
| `SpaceSharingSettings` | space_sharing_settings | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to apply to the space. |
| `Url` |  | `string` | computed |  |  |

Supports update: yes

Discovery: supported
