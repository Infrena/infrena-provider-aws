# aws.refreshschedule

**CloudFormation type:** `AWS::QuickSight::RefreshSchedule`

Definition of the AWS::QuickSight::RefreshSchedule Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|DataSetId|Schedule/ScheduleId` (AWS::QuickSight::RefreshSchedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the data source.</p> |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DataSetId` | data_set_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.dataset.DataSetId |  |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
