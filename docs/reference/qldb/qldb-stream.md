# aws.qldb.stream

**CloudFormation type:** `AWS::QLDB::Stream`

Resource schema for AWS::QLDB::Stream.

Region attribute: `region`

**Import ID:** `<region>/LedgerName|Id` (AWS::QLDB::Stream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ExclusiveEndTime` | exclusive_end_time | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `InclusiveStartTime` | inclusive_start_time | `string` | required, replaces on change |  |  |
| `KinesisConfiguration` | kinesis_configuration | `map` | required, replaces on change |  |  |
| `LedgerName` | ledger_name | `string` | required, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn |  |
| `StreamName` | stream_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
