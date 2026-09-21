# aws.experimenttemplate

**CloudFormation type:** `AWS::FIS::ExperimentTemplate`

Resource schema for AWS::FIS::ExperimentTemplate

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::FIS::ExperimentTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `map` | optional, computed, provider-chosen |  | The actions for the experiment. |
| `Description` |  | `string` | required |  | A description for the experiment template. |
| `ExperimentOptions` | experiment_options | `map` | optional, computed, provider-chosen |  |  |
| `ExperimentReportConfiguration` | experiment_report_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `LogConfiguration` | log_configuration | `map` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf. |
| `StopConditions` | stop_conditions | `list` | required |  | One or more stop conditions. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Targets` |  | `map` | required |  | The targets for the experiment. |

Supports update: yes

Discovery: supported
