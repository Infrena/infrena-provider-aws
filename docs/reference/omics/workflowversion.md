# aws.workflowversion

**CloudFormation type:** `AWS::Omics::WorkflowVersion`

Definition of AWS::Omics::WorkflowVersion Resource Type.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Omics::WorkflowVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Accelerators` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `ContainerRegistryMap` | container_registry_map | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ContainerRegistryMapUri` | container_registry_map_uri | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DefinitionRepository` | definition_repository | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `DefinitionUri` | definition_uri | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Engine` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Main` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ParameterTemplate` | parameter_template | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ParameterTemplatePath` | parameter_template_path | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Path to the primary workflow parameter template JSON file inside the repository |
| `Status` |  | `string` | computed |  |  |
| `StorageCapacity` | storage_capacity | `float` | optional, computed, provider-chosen |  |  |
| `StorageType` | storage_type | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of resource tags |
| `Type` | type_value | `string` | computed |  |  |
| `Uuid` |  | `string` | computed |  |  |
| `VersionName` | version_name | `string` | required, replaces on change |  |  |
| `WorkflowBucketOwnerId` | workflow_bucket_owner_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `WorkflowId` | workflow_id | `string` | required, replaces on change | aws.omics.workflow.Id |  |
| `readmeMarkdown` | readme_markdown | `string` | optional, computed, provider-chosen, write-only |  | The markdown content for the workflow's README file. This provides documentation and usage information for users of the workflow. |
| `readmePath` | readme_path | `string` | optional, computed, provider-chosen, replaces on change |  | The path to the workflow README markdown file within the repository. This file provides documentation and usage information for the workflow. If not specified, the README.md file from the root directory of the repository will be used. |
| `readmeUri` | readme_uri | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The S3 URI of the README file for the workflow. This file provides documentation and usage information for the workflow. The S3 URI must begin with s3://USER-OWNED-BUCKET/. The requester must have access to the S3 bucket and object. The max README content length is 500 KiB. |

Supports update: yes

Discovery: supported
