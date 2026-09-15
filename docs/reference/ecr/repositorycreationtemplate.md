# aws.repositorycreationtemplate

**CloudFormation type:** `AWS::ECR::RepositoryCreationTemplate`

The details of the repository creation template associated with the request.

Region attribute: `region`

**Import ID:** `<region>/Prefix` (AWS::ECR::RepositoryCreationTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppliedFor` | applied_for | `list` | required |  | A list of enumerable Strings representing the repository creation scenarios that this template will apply towards. The supported scenarios are PULL_THROUGH_CACHE, REPLICATION, and CREATE_ON_PUSH |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CustomRoleArn` | custom_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the role to be assumed by Amazon ECR. Amazon ECR will assume your supplied role when the customRoleArn is specified. When this field isn't specified, Amazon ECR will use the service-linked role for the repository creation template. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description associated with the repository creation template. |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen |  | The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. |
| `ImageTagMutability` | image_tag_mutability | `string` | optional, computed, provider-chosen |  | The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten. |
| `ImageTagMutabilityExclusionFilters` | image_tag_mutability_exclusion_filters | `list` | optional, computed, provider-chosen |  | A list of filters that specify which image tags are excluded from the repository creation template's image tag mutability setting. |
| `LifecyclePolicy` | lifecycle_policy | `string` | optional, computed, provider-chosen |  | The lifecycle policy to use for repositories created using the template. |
| `Prefix` |  | `string` | required, replaces on change |  | The repository namespace prefix associated with the repository creation template. |
| `RepositoryPolicy` | repository_policy | `string` | optional, computed, provider-chosen |  | The repository policy to apply to repositories created using the template. A repository policy is a permissions policy associated with a repository to control access permissions. |
| `ResourceTags` | resource_tags | `list` | optional, computed, provider-chosen |  | The metadata to apply to the repository to help you categorize and organize. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters. |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
