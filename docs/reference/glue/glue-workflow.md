# aws.glue.workflow

**CloudFormation type:** `AWS::Glue::Workflow`

Resource Type definition for AWS::Glue::Workflow

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::Workflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DefaultRunProperties` | default_run_properties | `map` | optional, computed, provider-chosen |  | A collection of properties to be used as part of each execution of the workflow |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the workflow |
| `MaxConcurrentRuns` | max_concurrent_runs | `integer` | optional, computed, provider-chosen |  | You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs. If you leave this parameter blank, there is no limit to the number of concurrent workflow runs. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the workflow representing the flow |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags to use with this workflow. |

Supports update: yes

Discovery: supported
