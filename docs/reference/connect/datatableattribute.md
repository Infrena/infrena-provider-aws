# aws.datatableattribute

**CloudFormation type:** `AWS::Connect::DataTableAttribute`

Resource Type definition for AWS::Connect::DataTableAttribute

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|DataTableArn|AttributeId` (AWS::Connect::DataTableAttribute)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttributeId` | attribute_id | `string` | computed |  |  |
| `DataTableArn` | data_table_arn | `string` | required, replaces on change | aws.datatable.Arn |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn |  |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `float` | computed |  |  |
| `LockVersion` | lock_version | `map` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Primary` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `Validation` |  | `map` | optional, computed, provider-chosen |  |  |
| `ValueType` | value_type | `string` | required |  |  |

Supports update: yes

Discovery: supported (parent resource required)
