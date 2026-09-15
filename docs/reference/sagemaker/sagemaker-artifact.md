# aws.sagemaker.artifact

**CloudFormation type:** `AWS::SageMaker::Artifact`

Resource type definition for AWS::SageMaker::Artifact. An artifact is a lineage tracking entity that represents a URI addressable object or data, such as an S3 URI of a dataset or the ECR registry path of an image.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::Artifact)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the artifact. |
| `ArtifactName` | artifact_name | `string` | optional, computed, provider-chosen |  | The name of the artifact. Must be unique to your account in an AWS Region. |
| `ArtifactType` | artifact_type | `string` | required, replaces on change |  | The artifact type. |
| `CreationTime` | creation_time | `string` | computed |  | When the artifact was created. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | When the artifact was last modified. |
| `MetadataProperties` | metadata_properties | `map` | optional, computed, provider-chosen, replaces on change |  | Metadata properties of the tracking entity, trial, or trial component. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | A list of properties to add to the artifact. |
| `Source` |  | `map` | required, replaces on change |  | The source of the artifact. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to apply to the artifact. |

Supports update: yes

Discovery: supported
