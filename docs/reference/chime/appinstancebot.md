# aws.appinstancebot

**CloudFormation type:** `AWS::Chime::AppInstanceBot`

Resource Type definition for AWS::Chime::AppInstanceBot

Region attribute: `region`

**Import ID:** `<region>/AppInstanceBotArn` (AWS::Chime::AppInstanceBot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppInstanceArn` | app_instance_arn | `string` | required, replaces on change | aws.appinstance.AppInstanceArn | The ARN of the AppInstance. |
| `AppInstanceBotArn` | app_instance_bot_arn | `string` | computed |  | The ARN of the AppInstanceBot. |
| `Configuration` |  | `map` | required |  | A structure that contains configuration data. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The time at which the AppInstanceBot was created, as an ISO 8601 timestamp. |
| `LastUpdatedTimestamp` | last_updated_timestamp | `string` | computed |  | The time at which the AppInstanceBot was last updated, as an ISO 8601 timestamp. |
| `Metadata` |  | `string` | optional, computed, provider-chosen |  | The metadata of the AppInstanceBot. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the AppInstanceBot. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the AppInstanceBot. |

Supports update: yes

Discovery: supported (parent resource required)
