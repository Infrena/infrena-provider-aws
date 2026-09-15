# aws.datasync.task

**CloudFormation type:** `AWS::DataSync::Task`

Resource schema for AWS::DataSync::Task.

Region attribute: `region`

**Import ID:** `<region>/TaskArn` (AWS::DataSync::Task)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudWatchLogGroupArn` | cloud_watch_log_group_arn | `string` | optional, computed, provider-chosen | aws.loggroup.Arn | The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task. |
| `DestinationLocationArn` | destination_location_arn | `string` | required, replaces on change |  | The ARN of an AWS storage resource's location. |
| `DestinationNetworkInterfaceArns` | destination_network_interface_arns | `list` | computed |  | The Amazon Resource Names (ARNs) of the destination ENIs (Elastic Network Interfaces) that were created for your subnet. |
| `Excludes` |  | `list` | optional, computed, provider-chosen |  |  |
| `Includes` |  | `list` | optional, computed, provider-chosen |  |  |
| `ManifestConfig` | manifest_config | `map` | optional, computed, provider-chosen |  | Configures a manifest, which is a list of files or objects that you want DataSync to transfer. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of a task. This value is a text reference that is used to identify the task in the console. |
| `Options` |  | `map` | optional, computed, provider-chosen |  | Represents the options that are available to control the behavior of a StartTaskExecution operation. |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  | Specifies the schedule you want your task to use for repeated executions. |
| `SourceLocationArn` | source_location_arn | `string` | required, replaces on change |  | The ARN of the source location for the task. |
| `SourceNetworkInterfaceArns` | source_network_interface_arns | `list` | computed |  | The Amazon Resource Names (ARNs) of the source ENIs (Elastic Network Interfaces) that were created for your subnet. |
| `Status` |  | `string` | computed |  | The status of the task that was described. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TaskArn` | task_arn | `string` | computed |  | The ARN of the task. |
| `TaskMode` | task_mode | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the task mode for the task. |
| `TaskReportConfig` | task_report_config | `map` | optional, computed, provider-chosen |  | Specifies how you want to configure a task report, which provides detailed information about for your Datasync transfer. |

Supports update: yes

Discovery: supported
