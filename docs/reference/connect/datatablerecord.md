# aws.datatablerecord

**CloudFormation type:** `AWS::Connect::DataTableRecord`

Resource Type definition for AWS::Connect::DataTableRecord

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|DataTableArn|RecordId` (AWS::Connect::DataTableRecord)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataTableArn` | data_table_arn | `string` | required, replaces on change | aws.datatable.Arn |  |
| `DataTableRecord` | data_table_record | `map` | required |  |  |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn |  |
| `RecordId` | record_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
