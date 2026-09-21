# aws.mlflowapp

**CloudFormation type:** `AWS::SageMaker::MlflowApp`

Resource type definition for AWS::SageMaker::MlflowApp

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SageMaker::MlflowApp)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the MLflow App. |
| `ArtifactStoreUri` | artifact_store_uri | `string` | required |  | The S3 URI for a general purpose bucket to use as the MLflow App artifact store. |
| `CreationTime` | creation_time | `string` | computed |  | The date and time that the MLflow App was created. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The date and time that the MLflow App was last modified. |
| `MlflowAppId` | mlflow_app_id | `string` | computed |  | The server-generated identifier of the MLflow App. |
| `MlflowVersion` | mlflow_version | `string` | computed |  | The MLflow version used by the MLflow App. |
| `ModelRegistrationMode` | model_registration_mode | `string` | optional, computed, provider-chosen |  | Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry. |
| `Name` |  | `string` | required, replaces on change |  | The name of the MLflow App. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) for an IAM role in your account that the MLflow App uses to access the artifact store in Amazon S3. |
| `Status` |  | `string` | computed |  | The status of the MLflow App. |
| `Tags` |  | `map` | tags map |  | Tags to associate with the MLflow App. |
| `WeeklyMaintenanceWindowStart` | weekly_maintenance_window_start | `string` | optional, computed, provider-chosen |  | The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled. For example: Tue:03:30. |

Supports update: yes

Discovery: supported
