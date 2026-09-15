# aws.databrew.schedule

**CloudFormation type:** `AWS::DataBrew::Schedule`

Resource schema for AWS::DataBrew::Schedule.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Schedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CronExpression` | cron_expression | `string` | required |  | Schedule cron |
| `JobNames` | job_names | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  | Schedule Name |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
