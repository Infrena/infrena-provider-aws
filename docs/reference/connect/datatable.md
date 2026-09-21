# aws.datatable

**CloudFormation type:** `AWS::Connect::DataTable`

Resource Type definition for AWS::Connect::DataTable

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|Arn` (AWS::Connect::DataTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The arn of the Data Table |
| `CreatedTime` | created_time | `float` | computed |  | A epoch time stamp field used for data table operations. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Data Table. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | Last modified region. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | A epoch time stamp field used for data table operations. |
| `LockVersion` | lock_version | `map` | computed |  | The lock version of the Data Table |
| `Name` |  | `string` | required |  | The name of the Data Table |
| `Status` |  | `string` | required, replaces on change |  | The status of the Data Table |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `TimeZone` | time_zone | `string` | required |  | The time zone of the Data Table |
| `ValueLockLevel` | value_lock_level | `string` | required |  | The value lock level of the Data Table |

Supports update: yes

Discovery: supported (parent resource required)
