# aws.supportapp.slackchannelconfiguration

**CloudFormation type:** `AWS::SupportApp::SlackChannelConfiguration`

An AWS Support App resource that creates, updates, lists and deletes Slack channel configurations.

Region attribute: `region`

**Import ID:** `<region>/TeamId|ChannelId` (AWS::SupportApp::SlackChannelConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelId` | channel_id | `string` | required, replaces on change |  | The channel ID in Slack, which identifies a channel within a workspace. |
| `ChannelName` | channel_name | `string` | optional, computed, provider-chosen |  | The channel name in Slack. |
| `ChannelRoleArn` | channel_role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services. |
| `NotifyOnAddCorrespondenceToCase` | notify_on_add_correspondence_to_case | `boolean` | optional, computed, provider-chosen |  | Whether to notify when a correspondence is added to a case. |
| `NotifyOnCaseSeverity` | notify_on_case_severity | `string` | required |  | The severity level of a support case that a customer wants to get notified for. |
| `NotifyOnCreateOrReopenCase` | notify_on_create_or_reopen_case | `boolean` | optional, computed, provider-chosen |  | Whether to notify when a case is created or reopened. |
| `NotifyOnResolveCase` | notify_on_resolve_case | `boolean` | optional, computed, provider-chosen |  | Whether to notify when a case is resolved. |
| `TeamId` | team_id | `string` | required, replaces on change |  | The team ID in Slack, which uniquely identifies a workspace. |

Supports update: yes

Discovery: supported
