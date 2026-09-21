# aws.logging

**CloudFormation type:** `AWS::IoT::Logging`

Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::IoT::Logging)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | required, replaces on change |  | Your 12-digit account ID (used as the primary identifier for the CloudFormation resource). |
| `DefaultLogLevel` | default_log_level | `string` | required |  | The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED. |
| `EventConfigurations` | event_configurations | `list` | optional, computed, provider-chosen |  | Configurations for event-based logging that specifies which event types to log and their logging settings. Overrides account-level logging for the specified event |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the role that allows IoT to write to Cloudwatch logs. |

Supports update: yes

Discovery: supported
