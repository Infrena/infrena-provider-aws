# aws.iot.job

**CloudFormation type:** `AWS::IoT::Job`

Use the AWS::IoT::Job resource to declare an AWS IoT job. A job can be used to define a set of remote operations that are sent to and run on one or more devices (things or thing groups) connected to AWS IoT.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IoT::Job)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AbortConfig` | abort_config | `map` | optional, computed, provider-chosen |  | The criteria that determine when and how a job abort takes place. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the job. |
| `CreatedAt` | created_at | `string` | computed |  | The time when the job was created, in ISO 8601 date-time format. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A short text description of the job. |
| `DestinationPackageVersions` | destination_package_versions | `list` | optional, computed, provider-chosen, replaces on change |  | The package version Amazon Resource Names (ARNs) that are installed on the device when the job successfully completes. |
| `Document` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The job document. Required if you don't specify a value for documentSource. |
| `DocumentParameters` | document_parameters | `map` | optional, computed, provider-chosen, replaces on change |  | Parameters of an Amazon Web Services managed template that you can specify to create the job document. |
| `DocumentSource` | document_source | `string` | optional, computed, provider-chosen, replaces on change |  | An S3 link, or S3 object URL, to the job document. The link is an Amazon S3 object URL and is required if you don't specify a value for document. |
| `JobExecutionsRetryConfig` | job_executions_retry_config | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration that determines how many retries are allowed for each failure type for a job. |
| `JobExecutionsRolloutConfig` | job_executions_rollout_config | `map` | optional, computed, provider-chosen |  | Allows you to create a staged rollout of a job. |
| `JobId` | job_id | `string` | required, replaces on change | aws.iot.job.JobId | A job identifier which must be unique for your AWS account. We recommend using a UUID. Alpha-numeric characters, '-' and '_' are valid for use here. |
| `JobTemplateArn` | job_template_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.jobtemplate.Arn | The ARN of the job template used to create the job. |
| `PresignedUrlConfig` | presigned_url_config | `map` | optional, computed, provider-chosen |  | Configuration for pre-signed S3 URLs. |
| `SchedulingConfig` | scheduling_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the date and time that a job will begin the rollout of the job document to all devices in the target group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata which can be used to manage the job. |
| `TargetSelection` | target_selection | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the job will continue to run (CONTINUOUS), or will be complete after all those things specified as targets have completed the job (SNAPSHOT). |
| `Targets` |  | `list` | required, replaces on change |  | A list of things and thing groups to which the job should be sent. |
| `TimeoutConfig` | timeout_config | `map` | optional, computed, provider-chosen |  | Specifies the amount of time each device has to finish its execution of the job. |

Supports update: yes

Discovery: supported
