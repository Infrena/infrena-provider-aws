# aws.mwaa.environment

**CloudFormation type:** `AWS::MWAA::Environment`

Resource schema for AWS::MWAA::Environment

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::MWAA::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AirflowConfigurationOptions` | airflow_configuration_options | `map` | optional, computed, provider-chosen |  | Key/value pairs representing Airflow configuration variables. |
| `AirflowVersion` | airflow_version | `string` | optional, computed, provider-chosen |  | Version of airflow to deploy to the environment. |
| `Arn` |  | `string` | computed |  | ARN for the MWAA environment. |
| `CeleryExecutorQueue` | celery_executor_queue | `string` | computed |  | The celery executor queue associated with the environment. |
| `DagS3Path` | dag_s3_path | `string` | optional, computed, provider-chosen |  | Represents an S3 prefix relative to the root of an S3 bucket. |
| `DatabaseVpcEndpointService` | database_vpc_endpoint_service | `string` | computed |  | The database VPC endpoint service name. |
| `EndpointManagement` | endpoint_management | `string` | optional, computed, provider-chosen, replaces on change |  | Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA. |
| `EnvironmentClass` | environment_class | `string` | optional, computed, provider-chosen |  | Templated configuration for airflow processes and backing infrastructure. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | IAM role to be used by tasks. |
| `KmsKey` | kms_key | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption. |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  | Logging configuration for the environment. |
| `MaxWebservers` | max_webservers | `integer` | optional, computed, provider-chosen |  | Maximum webserver compute units. |
| `MaxWorkers` | max_workers | `integer` | optional, computed, provider-chosen |  | Maximum worker compute units. |
| `MinWebservers` | min_webservers | `integer` | optional, computed, provider-chosen |  | Minimum webserver compute units. |
| `MinWorkers` | min_workers | `integer` | optional, computed, provider-chosen |  | Minimum worker compute units. |
| `Name` |  | `string` | required, replaces on change |  | Customer-defined identifier for the environment, unique per customer region. |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  | Configures the network resources of the environment. |
| `PluginsS3ObjectVersion` | plugins_s3_object_version | `string` | optional, computed, provider-chosen |  | Represents an version ID for an S3 object. |
| `PluginsS3Path` | plugins_s3_path | `string` | optional, computed, provider-chosen |  | Represents an S3 prefix relative to the root of an S3 bucket. |
| `RequirementsS3ObjectVersion` | requirements_s3_object_version | `string` | optional, computed, provider-chosen |  | Represents an version ID for an S3 object. |
| `RequirementsS3Path` | requirements_s3_path | `string` | optional, computed, provider-chosen |  | Represents an S3 prefix relative to the root of an S3 bucket. |
| `Schedulers` |  | `integer` | optional, computed, provider-chosen |  | Scheduler compute units. |
| `SourceBucketArn` | source_bucket_arn | `string` | optional, computed, provider-chosen |  | ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment. |
| `StartupScriptS3ObjectVersion` | startup_script_s3_object_version | `string` | optional, computed, provider-chosen |  | Represents an version ID for an S3 object. |
| `StartupScriptS3Path` | startup_script_s3_path | `string` | optional, computed, provider-chosen |  | Represents an S3 prefix relative to the root of an S3 bucket. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tags for the environment. |
| `WebserverAccessMode` | webserver_access_mode | `string` | optional, computed, provider-chosen |  | Choice for mode of webserver access including over public internet or via private VPC endpoint. |
| `WebserverUrl` | webserver_url | `string` | computed |  | Url endpoint for the environment's Airflow UI. |
| `WebserverVpcEndpointService` | webserver_vpc_endpoint_service | `string` | computed |  | The webserver VPC endpoint service name, applicable if private webserver access mode selected. |
| `WeeklyMaintenanceWindowStart` | weekly_maintenance_window_start | `string` | optional, computed, provider-chosen |  | Start time for the weekly maintenance window. |
| `WorkerReplacementStrategy` | worker_replacement_strategy | `string` | optional, computed, provider-chosen, write-only |  | The worker replacement strategy to use when updating the environment. Valid values: `FORCED`, `GRACEFUL`. FORCED means Apache Airflow workers will be stopped and replaced without waiting for tasks to complete before an update. GRACEFUL means Apache Airflow workers will be able to complete running tasks for up to 12 hours during an update before being stopped and replaced. |

Supports update: yes

Discovery: supported
