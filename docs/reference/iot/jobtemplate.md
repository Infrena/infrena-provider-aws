# aws.jobtemplate

**CloudFormation type:** `AWS::IoT::JobTemplate`

Resource Type definition for AWS::IoT::JobTemplate. Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.

Region attribute: `region`

**Import ID:** `<region>/JobTemplateId` (AWS::IoT::JobTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AbortConfig` | abort_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The criteria that determine when and how a job abort takes place. |
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | required, replaces on change |  | A description of the Job Template. |
| `DestinationPackageVersions` | destination_package_versions | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Document` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The job document. Required if you don't specify a value for documentSource. |
| `DocumentSource` | document_source | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | An S3 link to the job document to use in the template. Required if you don't specify a value for document. |
| `JobArn` | job_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.iot.job.Arn | Optional for copying a JobTemplate from a pre-existing Job configuration. |
| `JobExecutionsRetryConfig` | job_executions_retry_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `JobExecutionsRolloutConfig` | job_executions_rollout_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Allows you to create a staged rollout of a job. |
| `JobTemplateId` | job_template_id | `string` | required, replaces on change | aws.jobtemplate.JobTemplateId |  |
| `MaintenanceWindows` | maintenance_windows | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `PresignedUrlConfig` | presigned_url_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Configuration for pre-signed S3 URLs. |
| `Tags` |  | `map` | replaces on change, tags map |  | Metadata that can be used to manage the JobTemplate. |
| `TimeoutConfig` | timeout_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the amount of time each device has to finish its execution of the job. |

Supports update: no

Discovery: supported
