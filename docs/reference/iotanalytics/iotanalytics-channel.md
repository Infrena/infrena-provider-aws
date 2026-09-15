# aws.iotanalytics.channel

**CloudFormation type:** `AWS::IoTAnalytics::Channel`

Resource Type definition for AWS::IoTAnalytics::Channel

Region attribute: `region`

**Import ID:** `<region>/ChannelName` (AWS::IoTAnalytics::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelName` | channel_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ChannelStorage` | channel_storage | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `RetentionPeriod` | retention_period | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
