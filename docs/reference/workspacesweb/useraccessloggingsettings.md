# aws.useraccessloggingsettings

**CloudFormation type:** `AWS::WorkSpacesWeb::UserAccessLoggingSettings`

Definition of AWS::WorkSpacesWeb::UserAccessLoggingSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/UserAccessLoggingSettingsArn` (AWS::WorkSpacesWeb::UserAccessLoggingSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `KinesisStreamArn` | kinesis_stream_arn | `string` | required |  | Kinesis stream ARN to which log events are published. |
| `Tags` |  | `map` | tags map |  |  |
| `UserAccessLoggingSettingsArn` | user_access_logging_settings_arn | `string` | computed |  |  |

Supports update: yes

Discovery: supported
