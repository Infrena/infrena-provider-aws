# aws.iotanalytics.pipeline

**CloudFormation type:** `AWS::IoTAnalytics::Pipeline`

Resource Type definition for AWS::IoTAnalytics::Pipeline

Region attribute: `region`

**Import ID:** `<region>/PipelineName` (AWS::IoTAnalytics::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `PipelineActivities` | pipeline_activities | `list` | required |  |  |
| `PipelineName` | pipeline_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
