# aws.experimentrun

**CloudFormation type:** `AWS::AppConfig::ExperimentRun`

Resource Type definition for AWS::AppConfig::ExperimentRun

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|ExperimentDefinitionId|Run` (AWS::AppConfig::ExperimentRun)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | computed |  | The resolved application ID. |
| `ApplicationIdentifier` | application_identifier | `string` | required, replaces on change |  | The application name or ID used to create the experiment run. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the experiment run. |
| `ExperimentDefinitionId` | experiment_definition_id | `string` | computed |  | The resolved experiment definition ID. |
| `ExperimentDefinitionIdentifier` | experiment_definition_identifier | `string` | required, replaces on change |  | The experiment definition name or ID used to create the experiment run. |
| `ExposurePercentage` | exposure_percentage | `float` | required |  | Percentage of traffic exposed to the experiment (0-100). |
| `Run` |  | `string` | computed |  | The run number (auto-assigned by the service). |
| `StartedAt` | started_at | `string` | computed |  | ISO-8601 timestamp when the run started. |
| `Status` |  | `string` | computed |  | Current status of the run. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to associate with the experiment run. |
| `TreatmentOverrides` | treatment_overrides | `map` | optional, computed, provider-chosen |  | Treatment overrides for specific entities. |
| `UpdatedAt` | updated_at | `string` | computed |  | ISO-8601 timestamp when the run was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
