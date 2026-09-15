# aws.reportplan

**CloudFormation type:** `AWS::Backup::ReportPlan`

Contains detailed information about a report plan in AWS Backup Audit Manager.

Region attribute: `region`

**Import ID:** `<region>/ReportPlanArn` (AWS::Backup::ReportPlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ReportDeliveryChannel` | report_delivery_channel | `map` | required |  | A structure that contains information about where and how to deliver your reports, specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your reports. |
| `ReportPlanArn` | report_plan_arn | `string` | computed |  | An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of the ARN depends on the resource type. |
| `ReportPlanDescription` | report_plan_description | `string` | optional, computed, provider-chosen |  | An optional description of the report plan with a maximum of 1,024 characters. |
| `ReportPlanName` | report_plan_name | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name of the report plan. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_). |
| `ReportPlanTags` | report_plan_tags | `map` | optional, computed, provider-chosen, tags map |  | Metadata that you can assign to help organize the report plans that you create. Each tag is a key-value pair. |
| `ReportSetting` | report_setting | `map` | required |  | Identifies the report template for the report. Reports are built using a report template. |

Supports update: yes

Discovery: supported
