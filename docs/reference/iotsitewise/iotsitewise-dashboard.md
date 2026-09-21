# aws.iotsitewise.dashboard

**CloudFormation type:** `AWS::IoTSiteWise::Dashboard`

Resource schema for AWS::IoTSiteWise::Dashboard

Region attribute: `region`

**Import ID:** `<region>/DashboardId` (AWS::IoTSiteWise::Dashboard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DashboardArn` | dashboard_arn | `string` | computed |  | The ARN of the dashboard. |
| `DashboardDefinition` | dashboard_definition | `string` | required |  | The dashboard definition specified in a JSON literal. |
| `DashboardDescription` | dashboard_description | `string` | required |  | A description for the dashboard. |
| `DashboardId` | dashboard_id | `string` | computed |  | The ID of the dashboard. |
| `DashboardName` | dashboard_name | `string` | required |  | A friendly name for the dashboard. |
| `ProjectId` | project_id | `string` | optional, computed, provider-chosen, replaces on change | aws.iotsitewise.project.ProjectId | The ID of the project in which to create the dashboard. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the dashboard. |

Supports update: yes

Discovery: supported
