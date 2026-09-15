# aws.rungroup

**CloudFormation type:** `AWS::Omics::RunGroup`

Definition of AWS::Omics::RunGroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Omics::RunGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `MaxCpus` | max_cpus | `float` | optional, computed, provider-chosen |  |  |
| `MaxDuration` | max_duration | `float` | optional, computed, provider-chosen |  |  |
| `MaxGpus` | max_gpus | `float` | optional, computed, provider-chosen |  |  |
| `MaxRuns` | max_runs | `float` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of resource tags |

Supports update: yes

Discovery: supported
