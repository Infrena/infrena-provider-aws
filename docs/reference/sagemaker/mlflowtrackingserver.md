# aws.mlflowtrackingserver

**CloudFormation type:** `AWS::SageMaker::MlflowTrackingServer`

Resource Type definition for AWS::SageMaker::MlflowTrackingServer

Region attribute: `region`

**Import ID:** `<region>/TrackingServerName` (AWS::SageMaker::MlflowTrackingServer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArtifactStoreUri` | artifact_store_uri | `string` | required |  | The Amazon S3 URI for MLFlow Tracking Server artifacts. |
| `AutomaticModelRegistration` | automatic_model_registration | `boolean` | optional, computed, provider-chosen |  | A flag to enable Automatic SageMaker Model Registration. |
| `MlflowVersion` | mlflow_version | `string` | optional, computed, provider-chosen |  | The MLFlow Version used on the MLFlow Tracking Server. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TrackingServerArn` | tracking_server_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the MLFlow Tracking Server. |
| `TrackingServerName` | tracking_server_name | `string` | required, replaces on change |  | The name of the MLFlow Tracking Server. |
| `TrackingServerSize` | tracking_server_size | `string` | optional, computed, provider-chosen |  | The size of the MLFlow Tracking Server. |
| `WeeklyMaintenanceWindowStart` | weekly_maintenance_window_start | `string` | optional, computed, provider-chosen |  | The start of the time window for maintenance of the MLFlow Tracking Server in UTC time. |

Supports update: yes

Discovery: supported
