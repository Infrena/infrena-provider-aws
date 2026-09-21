# aws.changeset

**CloudFormation type:** `AWS::CloudFormation::ChangeSet`

Resource type definition for AWS::CloudFormation::ChangeSet

Region attribute: `region`

**Import ID:** `<region>/ChangeSetId` (AWS::CloudFormation::ChangeSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Capabilities` |  | `list` | optional, computed, provider-chosen, replaces on change |  | The capabilities that are allowed in the stack. |
| `ChangeSetId` | change_set_id | `string` | computed |  | The ARN of the change set. |
| `ChangeSetName` | change_set_name | `string` | required, replaces on change |  | The name of the change set. Must be unique among all change sets associated with the specified stack. |
| `ChangeSetType` | change_set_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The type of change set operation. |
| `CreationTime` | creation_time | `string` | computed |  | The time the change set was created. |
| `DeploymentMode` | deployment_mode | `string` | optional, computed, provider-chosen, replaces on change |  | Determines how CloudFormation handles configuration drift during deployment. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description to help you identify this change set. |
| `ImportExistingResources` | import_existing_resources | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates if the change set imports resources that already exist. |
| `IncludeNestedStacks` | include_nested_stacks | `boolean` | optional, computed, provider-chosen, replaces on change |  | Creates a change set for all nested stacks specified in the template. |
| `NotificationARNs` | notification_ar_ns | `list` | optional, computed, provider-chosen, replaces on change |  | The ARNs of Amazon SNS topics that CloudFormation associates with the stack. |
| `OnStackFailure` | on_stack_failure | `string` | optional, computed, provider-chosen, replaces on change |  | Determines what action will be taken if stack creation fails. |
| `RoleARN` | role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN of an IAM role that CloudFormation assumes when executing the change set. |
| `StackId` | stack_id | `string` | computed |  | The unique ID of the stack. |
| `StackName` | stack_name | `string` | required, replaces on change |  | The name or unique ID of the stack for which you are creating a change set. |
| `Tags` |  | `map` | replaces on change, tags map |  | Key-value pairs to associate with the change set. |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A structure that contains the body of the revised template. |
| `TemplateURL` | template_url | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The URL of the file that contains the revised template. |
| `UsePreviousTemplate` | use_previous_template | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Whether to reuse the template associated with the stack to create the change set. |

Supports update: no

Discovery: supported (parent resource required)
