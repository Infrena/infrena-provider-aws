# aws.applicationinsights.application

**CloudFormation type:** `AWS::ApplicationInsights::Application`

Resource Type definition for AWS::ApplicationInsights::Application

Region attribute: `region`

**Import ID:** `<region>/ApplicationARN` (AWS::ApplicationInsights::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationARN` | application_arn | `string` | computed |  | The ARN of the ApplicationInsights application. |
| `AttachMissingPermission` | attach_missing_permission | `boolean` | optional, computed, provider-chosen, write-only |  | If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing |
| `AutoConfigurationEnabled` | auto_configuration_enabled | `boolean` | optional, computed, provider-chosen |  | If set to true, application will be configured with recommended monitoring configuration. |
| `CWEMonitorEnabled` | cwe_monitor_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether Application Insights can listen to CloudWatch events for the application resources. |
| `ComponentMonitoringSettings` | component_monitoring_settings | `list` | optional, computed, provider-chosen, write-only |  | The monitoring settings of the components. |
| `CustomComponents` | custom_components | `list` | optional, computed, provider-chosen, write-only |  | The custom grouped components. |
| `GroupingType` | grouping_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The grouping type of the application |
| `LogPatternSets` | log_pattern_sets | `list` | optional, computed, provider-chosen, write-only |  | The log pattern sets. |
| `OpsCenterEnabled` | ops_center_enabled | `boolean` | optional, computed, provider-chosen |  | When set to true, creates opsItems for any problems detected on an application. |
| `OpsItemSNSTopicArn` | ops_item_sns_topic_arn | `string` | optional, computed, provider-chosen, write-only |  | The SNS topic provided to Application Insights that is associated to the created opsItem. |
| `ResourceGroupName` | resource_group_name | `string` | required, replaces on change |  | The name of the resource group. |
| `SNSNotificationArn` | sns_notification_arn | `string` | optional, computed, provider-chosen, write-only |  | Application Insights sends notifications to this SNS topic whenever there is a problem update in the associated application. |
| `Tags` |  | `map` | tags map |  | The tags of Application Insights application. |

Supports update: yes

Discovery: supported
