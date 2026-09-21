# aws.snapshotschedule

**CloudFormation type:** `AWS::Redshift::SnapshotSchedule`

Creates a snapshot schedule that lets you set up automatic snapshots of your Amazon Redshift cluster at regular intervals.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Redshift::SnapshotSchedule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the snapshot schedule. |
| `AssociatedClusterCount` | associated_cluster_count | `integer` | computed |  | The number of clusters associated with the schedule. |
| `ScheduleDefinitions` | schedule_definitions | `list` | required |  | The definition of the snapshot schedule. The definition is made up of schedule expressions, for example "cron(30 12 *)" or "rate(12 hours)". |
| `ScheduleDescription` | schedule_description | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the snapshot schedule. |
| `ScheduleIdentifier` | schedule_identifier | `string` | required, replaces on change |  | A unique identifier for the snapshot schedule. Only alphanumeric characters are allowed. |
| `Tags` |  | `map` | tags map |  | An optional set of tags for the snapshot schedule. |

Supports update: yes

Discovery: supported
