# aws.monitoringschedule

**CloudFormation type:** `AWS::SageMaker::MonitoringSchedule`

Resource Type definition for AWS::SageMaker::MonitoringSchedule

Region attribute: `region`

**Import ID:** `<region>/MonitoringScheduleArn` (AWS::SageMaker::MonitoringSchedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time at which the schedule was created. |
| `EndpointName` | endpoint_name | `string` | optional, computed, provider-chosen |  | The name of the endpoint used to run the monitoring job. |
| `FailureReason` | failure_reason | `string` | optional, computed, provider-chosen |  | Contains the reason a monitoring job failed, if it failed. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | A timestamp that indicates the last time the monitoring job was modified. |
| `LastMonitoringExecutionSummary` | last_monitoring_execution_summary | `map` | optional, computed, provider-chosen |  | Summary of information about monitoring job |
| `MonitoringScheduleArn` | monitoring_schedule_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the monitoring schedule. |
| `MonitoringScheduleConfig` | monitoring_schedule_config | `map` | required |  | The configuration object that specifies the monitoring schedule and defines the monitoring job. |
| `MonitoringScheduleName` | monitoring_schedule_name | `string` | required, replaces on change |  | The name of the monitoring schedule. |
| `MonitoringScheduleStatus` | monitoring_schedule_status | `string` | optional, computed, provider-chosen |  | The status of a schedule job. |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
