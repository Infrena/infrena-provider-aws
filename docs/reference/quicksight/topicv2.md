# aws.topicv2

**CloudFormation type:** `AWS::QuickSight::TopicV2`

Definition of the AWS::QuickSight::TopicV2 Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|TopicId` (AWS::QuickSight::TopicV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CustomInstructions` | custom_instructions | `map` | optional, computed, provider-chosen |  |  |
| `DataSetRelations` | data_set_relations | `list` | optional, computed, provider-chosen |  |  |
| `DataSets` | data_sets | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, replaces on change, write-only | aws.folder.Arn |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | replaces on change, tags map |  |  |
| `TopicId` | topic_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.topic.TopicId |  |

Supports update: yes

Discovery: supported
