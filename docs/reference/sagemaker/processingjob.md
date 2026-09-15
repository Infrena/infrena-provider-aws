# aws.processingjob

**CloudFormation type:** `AWS::SageMaker::ProcessingJob`

Resource Type definition for AWS::SageMaker::ProcessingJob

Region attribute: `region`

**Import ID:** `<region>/ProcessingJobArn` (AWS::SageMaker::ProcessingJob)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppSpecification` | app_specification | `map` | required, replaces on change |  | Configures the processing job to run a specified Docker container image. |
| `AutoMLJobArn` | auto_ml_job_arn | `string` | computed |  | The ARN of an AutoML job associated with this processing job. |
| `CreationTime` | creation_time | `string` | computed |  | The time at which the processing job was created. |
| `Environment` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Sets the environment variables in the Docker container |
| `ExitMessage` | exit_message | `string` | computed |  | An optional string, up to one KB in size, that contains metadata from the processing container when the processing job exits. |
| `ExperimentConfig` | experiment_config | `map` | optional, computed, provider-chosen, replaces on change |  | Associates a SageMaker job as a trial component with an experiment and trial. |
| `FailureReason` | failure_reason | `string` | computed |  | A string, up to one KB in size, that contains the reason a processing job failed, if it failed. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the processing job was last modified. |
| `MonitoringScheduleArn` | monitoring_schedule_arn | `string` | computed |  | The ARN of a monitoring schedule for an endpoint associated with this processing job. |
| `NetworkConfig` | network_config | `map` | optional, computed, provider-chosen, replaces on change |  | Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs. |
| `ProcessingEndTime` | processing_end_time | `string` | computed |  | The time at which the processing job completed. |
| `ProcessingInputs` | processing_inputs | `list` | optional, computed, provider-chosen, replaces on change |  | An array of inputs configuring the data to download into the processing container. |
| `ProcessingJobArn` | processing_job_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the processing job. |
| `ProcessingJobName` | processing_job_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the processing job. The name must be unique within an AWS Region in the AWS account. |
| `ProcessingJobStatus` | processing_job_status | `string` | computed |  | Provides the status of a processing job. |
| `ProcessingOutputConfig` | processing_output_config | `map` | optional, computed, provider-chosen, replaces on change |  | Configuration for uploading output from the processing container. |
| `ProcessingResources` | processing_resources | `map` | required, replaces on change |  | Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance. |
| `ProcessingStartTime` | processing_start_time | `string` | computed |  | The time at which the processing job started. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf. |
| `StoppingCondition` | stopping_condition | `map` | optional, computed, provider-chosen, replaces on change |  | Configures conditions under which the processing job should be stopped, such as how long the processing job has been running. After the condition is met, the processing job is stopped. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | (Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags(https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the AWS Billing and Cost Management User Guide. |
| `TrainingJobArn` | training_job_arn | `string` | computed |  | The ARN of a training job associated with this processing job |

Supports update: no

Discovery: supported
