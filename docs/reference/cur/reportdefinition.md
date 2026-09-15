# aws.reportdefinition

**CloudFormation type:** `AWS::CUR::ReportDefinition`

The AWS::CUR::ReportDefinition resource creates a Cost & Usage Report with user-defined settings. You can use this resource to define settings like time granularity (hourly, daily, monthly), file format (Parquet, CSV), and S3 bucket for delivery of these reports.

Region attribute: `region`

**Import ID:** `<region>/ReportName` (AWS::CUR::ReportDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalArtifacts` | additional_artifacts | `list` | optional, computed, provider-chosen |  | A list of manifests that you want Amazon Web Services to create for this report. |
| `AdditionalSchemaElements` | additional_schema_elements | `list` | optional, computed, provider-chosen, replaces on change |  | A list of strings that indicate additional content that Amazon Web Services includes in the report, such as individual resource IDs. |
| `BillingViewArn` | billing_view_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon resource name of the billing view. You can get this value by using the billing view service public APIs. |
| `Compression` |  | `string` | required |  | The compression format that AWS uses for the report. |
| `Format` |  | `string` | required |  | The format that AWS saves the report in. |
| `RefreshClosedReports` | refresh_closed_reports | `boolean` | required |  | Whether you want Amazon Web Services to update your reports after they have been finalized if Amazon Web Services detects charges related to previous months. These charges can include refunds, credits, or support fees. |
| `ReportName` | report_name | `string` | required, replaces on change |  | The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces. |
| `ReportVersioning` | report_versioning | `string` | required, replaces on change |  | Whether you want Amazon Web Services to overwrite the previous version of each report or to deliver the report in addition to the previous versions. |
| `S3Bucket` | s3_bucket | `string` | required |  | The S3 bucket where AWS delivers the report. |
| `S3Prefix` | s3_prefix | `string` | required |  | The prefix that AWS adds to the report name when AWS delivers the report. Your prefix can't include spaces. |
| `S3Region` | s3_region | `string` | required |  | The region of the S3 bucket that AWS delivers the report into. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TimeUnit` | time_unit | `string` | required, replaces on change |  | The granularity of the line items in the report. |

Supports update: yes

Discovery: supported
