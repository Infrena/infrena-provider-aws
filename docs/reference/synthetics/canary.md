# aws.canary

**CloudFormation type:** `AWS::Synthetics::Canary`

Resource Type definition for AWS::Synthetics::Canary

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Synthetics::Canary)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArtifactConfig` | artifact_config | `map` | optional, computed, provider-chosen |  | Provide artifact configuration |
| `ArtifactS3Location` | artifact_s3_location | `string` | required |  | Provide the s3 bucket output location for test results |
| `BrowserConfigs` | browser_configs | `list` | optional, computed, provider-chosen |  | List of browser configurations for the canary |
| `Code` |  | `map` | required |  | Provide the canary script source |
| `DeleteLambdaResourcesOnCanaryDeletion` | delete_lambda_resources_on_canary_deletion | `boolean` | optional, computed, provider-chosen, write-only |  | Deletes associated lambda resources created by Synthetics if set to True. Default is False |
| `DryRunAndUpdate` | dry_run_and_update | `boolean` | optional, computed, provider-chosen, write-only |  | Setting to control if UpdateCanary will perform a DryRun and validate it is PASSING before performing the Update. Default is FALSE. |
| `ExecutionRoleArn` | execution_role_arn | `string` | required | aws.role.Arn | Lambda Execution role used to run your canaries |
| `FailureRetentionPeriod` | failure_retention_period | `integer` | optional, computed, provider-chosen |  | Retention period of failed canary runs represented in number of days |
| `Id` |  | `string` | computed |  | Id of the canary |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | KMS key ARN for encrypting the canary's Lambda function environment variables at rest. If omitted, Lambda uses an AWS-managed key. |
| `Name` |  | `string` | required, replaces on change |  | Name of the canary. |
| `ProvisionedResourceCleanup` | provisioned_resource_cleanup | `string` | optional, computed, provider-chosen |  | Setting to control if provisioned resources created by Synthetics are deleted alongside the canary. Default is AUTOMATIC. |
| `Replicas` |  | `list` | optional, computed, provider-chosen |  | List of replica locations for multi-location canary execution |
| `ResourcesToReplicateTags` | resources_to_replicate_tags | `list` | optional, computed, provider-chosen, write-only |  | List of resources which canary tags should be replicated to. |
| `RunConfig` | run_config | `map` | optional, computed, provider-chosen |  | Provide canary run configuration |
| `RuntimeVersion` | runtime_version | `string` | required |  | Runtime version of Synthetics Library |
| `Schedule` |  | `map` | required |  | Frequency to run your canaries |
| `StartCanaryAfterCreation` | start_canary_after_creation | `boolean` | optional, computed, provider-chosen, write-only |  | Runs canary if set to True. Default is False |
| `State` |  | `string` | computed |  | State of the canary |
| `SuccessRetentionPeriod` | success_retention_period | `integer` | optional, computed, provider-chosen |  | Retention period of successful canary runs represented in number of days |
| `Tags` |  | `map` | tags map |  |  |
| `VPCConfig` | vpc_config | `map` | optional, computed, provider-chosen |  | Provide VPC Configuration if enabled. |
| `VisualReference` | visual_reference | `map` | optional, computed, provider-chosen, write-only |  | Visual reference configuration for visual testing |
| `VisualReferences` | visual_references | `list` | optional, computed, provider-chosen, write-only |  | List of visual references for the canary |

Supports update: yes

Discovery: supported
