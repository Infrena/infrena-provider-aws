# aws.cloudwatch.dashboard

**CloudFormation type:** `AWS::CloudWatch::Dashboard`

Resource Type definition for AWS::CloudWatch::Dashboard

Region attribute: `region`

**Import ID:** `<region>/DashboardName` (AWS::CloudWatch::Dashboard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DashboardBody` | dashboard_body | `string` | required |  | The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard |
| `DashboardName` | dashboard_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs to associate with the cloudwatch dashboard. You can associate up to 50 tags with a dashboard |

Supports update: yes

Discovery: supported
