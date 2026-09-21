# aws.glue.job

**CloudFormation type:** `AWS::Glue::Job`

Resource Type definition for AWS::Glue::Job

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::Job)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocatedCapacity` | allocated_capacity | `float` | optional, computed, provider-chosen |  | The number of capacity units that are allocated to this job. |
| `Command` |  | `map` | required |  | The code that executes a job. |
| `Connections` |  | `map` | optional, computed, provider-chosen |  | Specifies the connections used by a job |
| `DefaultArguments` | default_arguments | `map` | optional, computed, provider-chosen |  | The default arguments for this job, specified as name-value pairs. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the job. |
| `ExecutionClass` | execution_class | `string` | optional, computed, provider-chosen |  | Indicates whether the job is run with a standard or flexible execution class. |
| `ExecutionProperty` | execution_property | `map` | optional, computed, provider-chosen |  | The maximum number of concurrent runs that are allowed for this job. |
| `GlueVersion` | glue_version | `string` | optional, computed, provider-chosen |  | Glue version determines the versions of Apache Spark and Python that AWS Glue supports. |
| `JobMode` | job_mode | `string` | optional, computed, provider-chosen |  | Property description not available. |
| `JobRunQueuingEnabled` | job_run_queuing_enabled | `boolean` | optional, computed, provider-chosen |  | Property description not available. |
| `LogUri` | log_uri | `string` | optional, computed, provider-chosen |  | This field is reserved for future use. |
| `MaintenanceWindow` | maintenance_window | `string` | optional, computed, provider-chosen |  | Property description not available. |
| `MaxCapacity` | max_capacity | `float` | optional, computed, provider-chosen |  | The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. |
| `MaxRetries` | max_retries | `float` | optional, computed, provider-chosen |  | The maximum number of times to retry this job after a JobRun fails |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name you assign to the job definition |
| `NonOverridableArguments` | non_overridable_arguments | `map` | optional, computed, provider-chosen |  | Non-overridable arguments for this job, specified as name-value pairs. |
| `NotificationProperty` | notification_property | `map` | optional, computed, provider-chosen |  | Specifies configuration properties of a notification. |
| `NumberOfWorkers` | number_of_workers | `integer` | optional, computed, provider-chosen |  | The number of workers of a defined workerType that are allocated when a job runs. |
| `Role` |  | `string` | required |  | The name or Amazon Resource Name (ARN) of the IAM role associated with this job. |
| `SecurityConfiguration` | security_configuration | `string` | optional, computed, provider-chosen |  | The name of the SecurityConfiguration structure to be used with this job. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags to use with this job. |
| `Timeout` |  | `integer` | optional, computed, provider-chosen |  | The maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. |
| `WorkerType` | worker_type | `string` | optional, computed, provider-chosen |  | TThe type of predefined worker that is allocated when a job runs. |

Supports update: yes

Discovery: supported
