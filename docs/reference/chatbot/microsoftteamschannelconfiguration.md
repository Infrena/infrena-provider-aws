# aws.microsoftteamschannelconfiguration

**CloudFormation type:** `AWS::Chatbot::MicrosoftTeamsChannelConfiguration`

Resource schema for AWS::Chatbot::MicrosoftTeamsChannelConfiguration.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Chatbot::MicrosoftTeamsChannelConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the configuration |
| `ConfigurationName` | configuration_name | `string` | required, replaces on change |  | The name of the configuration |
| `CustomizationResourceArns` | customization_resource_arns | `list` | optional, computed, provider-chosen |  | ARNs of Custom Actions to associate with notifications in the provided chat channel. |
| `GuardrailPolicies` | guardrail_policies | `list` | optional, computed, provider-chosen |  | The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set. |
| `IamRoleArn` | iam_role_arn | `string` | required | aws.role.Arn | The ARN of the IAM role that defines the permissions for AWS Chatbot |
| `LoggingLevel` | logging_level | `string` | optional, computed, provider-chosen |  | Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs |
| `SnsTopicArns` | sns_topic_arns | `list` | optional, computed, provider-chosen |  | ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications. |
| `Tags` |  | `map` | tags map |  | The tags to add to the configuration |
| `TeamId` | team_id | `string` | required, replaces on change |  | The id of the Microsoft Teams team |
| `TeamsChannelId` | teams_channel_id | `string` | required |  | The id of the Microsoft Teams channel |
| `TeamsChannelName` | teams_channel_name | `string` | optional, computed, provider-chosen |  | The name of the Microsoft Teams channel |
| `TeamsTenantId` | teams_tenant_id | `string` | required, replaces on change |  | The id of the Microsoft Teams tenant |
| `UserRoleRequired` | user_role_required | `boolean` | optional, computed, provider-chosen |  | Enables use of a user role requirement in your chat configuration |

Supports update: yes

Discovery: supported
