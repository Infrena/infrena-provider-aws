# aws.eventtracker

**CloudFormation type:** `AWS::Personalize::EventTracker`

Resource schema for AWS::Personalize::EventTracker

Region attribute: `region`

**Import ID:** `<region>/EventTrackerArn` (AWS::Personalize::EventTracker)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetGroupArn` | dataset_group_arn | `string` | required, replaces on change | aws.personalize.datasetgroup.DatasetGroupArn | The Amazon Resource Name (ARN) of the dataset group that receives the event data. |
| `EventTrackerArn` | event_tracker_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the event tracker. |
| `Name` |  | `string` | required, replaces on change |  | The name for the event tracker. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the event tracker. |
| `TrackingId` | tracking_id | `string` | computed |  | The ID of the event tracker. Include this ID in requests to the PutEvents API. |

Supports update: yes

Discovery: supported
