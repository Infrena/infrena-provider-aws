# aws.timestream.scheduledquery

**CloudFormation type:** `AWS::Timestream::ScheduledQuery`

The AWS::Timestream::ScheduledQuery resource creates a Timestream Scheduled Query.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Timestream::ScheduledQuery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name of the scheduled query that is generated upon creation. |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, write-only |  | Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result. Making multiple identical CreateScheduledQuery requests has the same effect as making a single request. If CreateScheduledQuery is called without a ClientToken, the Query SDK generates a ClientToken on your behalf. After 8 hours, any request with the same ClientToken is treated as a new request. |
| `ErrorReportConfiguration` | error_report_configuration | `map` | required, replaces on change |  | Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest. |
| `NotificationConfiguration` | notification_configuration | `map` | required, replaces on change |  | Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it. |
| `QueryString` | query_string | `string` | required, replaces on change |  | The query string to run. Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter @scheduled_runtime is reserved and can be used in the query to get the time at which the query is scheduled to run. The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of @scheduled_runtime paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the @scheduled_runtime parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query. |
| `SQErrorReportConfiguration` | sq_error_report_configuration | `string` | computed |  | Configuration for error reporting. Error reports will be generated when a problem is encountered when writing the query results. |
| `SQKmsKeyId` | sq_kms_key_id | `string` | computed |  | The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with alias/. If ErrorReportConfiguration uses SSE_KMS as encryption type, the same KmsKeyId is used to encrypt the error report at rest. |
| `SQName` | sq_name | `string` | computed |  | The name of the scheduled query. Scheduled query names must be unique within each Region. |
| `SQNotificationConfiguration` | sq_notification_configuration | `string` | computed |  | Notification configuration for the scheduled query. A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it. |
| `SQQueryString` | sq_query_string | `string` | computed |  | The query string to run. Parameter names can be specified in the query string @ character followed by an identifier. The named Parameter @scheduled_runtime is reserved and can be used in the query to get the time at which the query is scheduled to run. The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of @scheduled_runtime paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the @scheduled_runtime parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query. |
| `SQScheduleConfiguration` | sq_schedule_configuration | `string` | computed |  | Configuration for when the scheduled query is executed. |
| `SQScheduledQueryExecutionRoleArn` | sq_scheduled_query_execution_role_arn | `string` | computed |  | The ARN for the IAM role that Timestream will assume when running the scheduled query. |
| `SQTargetConfiguration` | sq_target_configuration | `string` | computed |  | Configuration of target store where scheduled query results are written to. |
| `ScheduleConfiguration` | schedule_configuration | `map` | required, replaces on change |  | Configuration for when the scheduled query is executed. |
| `ScheduledQueryExecutionRoleArn` | scheduled_query_execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The ARN for the IAM role that Timestream will assume when running the scheduled query. |
| `ScheduledQueryName` | scheduled_query_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the scheduled query. Scheduled query names must be unique within each Region. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs to label the scheduled query. |
| `TargetConfiguration` | target_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | Configuration of target store where scheduled query results are written to. |

Supports update: yes

Discovery: supported
