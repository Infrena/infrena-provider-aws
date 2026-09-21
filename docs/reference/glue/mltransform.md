# aws.mltransform

**CloudFormation type:** `AWS::Glue::MLTransform`

Resource Type definition for AWS::Glue::MLTransform

Region attribute: `region`

**Import ID:** `<region>/TransformId` (AWS::Glue::MLTransform)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A user-defined, long-form description text for the machine learning transform. |
| `GlueVersion` | glue_version | `string` | optional, computed, provider-chosen |  | The version of AWS Glue this machine learning transform is compatible with. |
| `InputRecordTables` | input_record_tables | `map` | required, replaces on change |  | A list of AWS Glue table definitions used by the transform. |
| `MaxCapacity` | max_capacity | `float` | optional, computed, provider-chosen |  | The number of AWS Glue DPUs allocated to task runs for this transform. |
| `MaxRetries` | max_retries | `integer` | optional, computed, provider-chosen |  | The maximum number of times to retry after an MLTaskRun fails. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | A user-defined name for the machine learning transform. |
| `NumberOfWorkers` | number_of_workers | `integer` | optional, computed, provider-chosen |  | The number of workers of a defined workerType that are allocated when a task runs. |
| `Role` |  | `string` | required |  | The name or ARN of the IAM role with the required permissions. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags to use with this machine learning transform. |
| `Timeout` |  | `integer` | optional, computed, provider-chosen |  | The timeout in minutes of the machine learning transform. |
| `TransformEncryption` | transform_encryption | `map` | optional, computed, provider-chosen, replaces on change |  | The encryption-at-rest settings of the transform that apply to accessing user data. |
| `TransformId` | transform_id | `string` | computed |  | The unique identifier for the transform. |
| `TransformParameters` | transform_parameters | `map` | required |  | The algorithm-specific parameters that are associated with the machine learning transform. |
| `WorkerType` | worker_type | `string` | optional, computed, provider-chosen |  | The type of predefined worker that is allocated when a task runs. |

Supports update: yes

Discovery: supported
