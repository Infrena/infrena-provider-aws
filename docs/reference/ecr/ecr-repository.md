# aws.ecr.repository

**CloudFormation type:** `AWS::ECR::Repository`

The ``AWS::ECR::Repository`` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.

Region attribute: `region`

**Import ID:** `<region>/RepositoryName` (AWS::ECR::Repository)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `EmptyOnDelete` | empty_on_delete | `boolean` | optional, computed, provider-chosen, write-only |  | If true, deleting the repository force deletes the contents of the repository. Without a force delete, you can only delete empty repositories. |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. |
| `ImageScanningConfiguration` | image_scanning_configuration | `map` | optional, computed, provider-chosen |  | The image scanning configuration for a repository. |
| `ImageTagMutability` | image_tag_mutability | `string` | optional, computed, provider-chosen |  | The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten. |
| `ImageTagMutabilityExclusionFilters` | image_tag_mutability_exclusion_filters | `list` | optional, computed, provider-chosen |  | A list of filters that specify which image tags are excluded from the repository's image tag mutability setting. |
| `LifecyclePolicy` | lifecycle_policy | `map` | optional, computed, provider-chosen |  | The ``LifecyclePolicy`` property type specifies a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html) in the *Amazon ECR User Guide*. |
| `RepositoryName` | repository_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). |
| `RepositoryPolicyText` | repository_policy_text | `string` | optional, computed, provider-chosen |  | The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*. |
| `RepositoryUri` | repository_uri | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
