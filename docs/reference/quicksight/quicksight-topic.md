# aws.quicksight.topic

**CloudFormation type:** `AWS::QuickSight::Topic`

Definition of the AWS::QuickSight::Topic Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|TopicId` (AWS::QuickSight::Topic)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ConfigOptions` | config_options | `map` | optional, computed, provider-chosen |  |  |
| `CustomInstructions` | custom_instructions | `map` | optional, computed, provider-chosen |  | <p>Instructions that provide additional guidance and context for response generation.</p> |
| `DataSets` | data_sets | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, replaces on change, write-only | aws.folder.Arn |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | replaces on change, tags map |  |  |
| `TopicId` | topic_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.topic.TopicId |  |
| `UserExperienceVersion` | user_experience_version | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
