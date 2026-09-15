# aws.capacitymanagerdataexport

**CloudFormation type:** `AWS::EC2::CapacityManagerDataExport`

Resource Type definition for AWS::EC2::CapacityManagerDataExport

Region attribute: `region`

**Import ID:** `<region>/CapacityManagerDataExportId` (AWS::EC2::CapacityManagerDataExport)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityManagerDataExportId` | capacity_manager_data_export_id | `string` | computed |  | The unique identifier of the capacity manager data export. |
| `OutputFormat` | output_format | `string` | required, replaces on change |  | The format of the exported capacity manager data. Choose 'csv' for comma-separated values or 'parquet' for optimized columnar storage format. |
| `S3BucketName` | s3_bucket_name | `string` | required, replaces on change |  | The name of the Amazon S3 bucket where the capacity manager data export will be stored. The bucket must exist and be accessible by EC2 Capacity Manager service. |
| `S3BucketPrefix` | s3_bucket_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | The prefix for the S3 bucket location where exported files will be placed. If not specified, files will be placed in the root of the bucket. |
| `Schedule` |  | `string` | required, replaces on change |  | The schedule for the capacity manager data export. Currently supports hourly exports that provide periodic snapshots of capacity manager data. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to the capacity manager data export. |

Supports update: yes

Discovery: supported
