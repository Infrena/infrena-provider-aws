# aws.kinesisanalyticsv2.application

**CloudFormation type:** `AWS::KinesisAnalyticsV2::Application`

Creates an Amazon Kinesis Data Analytics application. For information about creating a Kinesis Data Analytics application, see [Creating an Application](https://docs.aws.amazon.com/kinesisanalytics/latest/java/getting-started.html).

Region attribute: `region`

**Import ID:** `<region>/ApplicationName` (AWS::KinesisAnalyticsV2::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationConfiguration` | application_configuration | `map` | optional, computed, provider-chosen |  | Specifies the creation parameters for a Kinesis Data Analytics application. |
| `ApplicationDescription` | application_description | `string` | optional, computed, provider-chosen |  | The description of the application. |
| `ApplicationMaintenanceConfiguration` | application_maintenance_configuration | `map` | optional, computed, provider-chosen |  | Describes the maintenance configuration for the application. |
| `ApplicationMode` | application_mode | `string` | optional, computed, provider-chosen, replaces on change |  | To create a Kinesis Data Analytics Studio notebook, you must set the mode to `INTERACTIVE`. However, for a Kinesis Data Analytics for Apache Flink application, the mode is optional. |
| `ApplicationName` | application_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the application. |
| `RunConfiguration` | run_configuration | `map` | optional, computed, provider-chosen, write-only |  | Identifies the run configuration (start parameters) of a Kinesis Data Analytics application. This section is evaluated only on stack updates for applications in running RUNNING state and has no effect during manual application start. |
| `RuntimeEnvironment` | runtime_environment | `string` | required |  | The runtime environment for the application. |
| `ServiceExecutionRole` | service_execution_role | `string` | required |  | The Amazon Resource Name |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of one or more tags to assign to the application. A tag is a key-value pair that identifies an application. Note that the maximum number of application tags includes system tags. The maximum number of user-defined application tags is 50. |

Supports update: yes

Discovery: supported
