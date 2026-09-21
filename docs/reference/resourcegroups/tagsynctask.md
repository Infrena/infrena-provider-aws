# aws.tagsynctask

**CloudFormation type:** `AWS::ResourceGroups::TagSyncTask`

Schema for ResourceGroups::TagSyncTask

Region attribute: `region`

**Import ID:** `<region>/TaskArn` (AWS::ResourceGroups::TagSyncTask)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Group` |  | `string` | required, replaces on change |  | The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task |
| `GroupArn` | group_arn | `string` | computed |  | The Amazon resource name (ARN) of the ApplicationGroup for which the TagSyncTask is created |
| `GroupName` | group_name | `string` | computed |  | The Name of the application group for which the TagSyncTask is created |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf. |
| `Status` |  | `string` | computed |  | The status of the TagSyncTask |
| `TagKey` | tag_key | `string` | required, replaces on change |  | The tag key. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application. |
| `TagValue` | tag_value | `string` | required, replaces on change |  | The tag value. Resources tagged with this tag key-value pair will be added to the application. If a resource with this tag is later untagged, the tag-sync task removes the resource from the application. |
| `TaskArn` | task_arn | `string` | computed |  | The ARN of the TagSyncTask resource |

Supports update: no

Discovery: supported
