# aws.voiceconnector

**CloudFormation type:** `AWS::Chime::VoiceConnector`

An Amazon Chime SDK Voice Connector configuration, including outbound host name and encryption settings.

Region attribute: `region`

**Import ID:** `<region>/VoiceConnectorArn` (AWS::Chime::VoiceConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsRegion` | aws_region | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Region in which the Voice Connector is created. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The Voice Connector creation timestamp, in ISO 8601 format. |
| `Name` |  | `string` | required |  | The name of the Voice Connector. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of network for the Voice Connector. |
| `OutboundHostName` | outbound_host_name | `string` | computed |  | The outbound host name for the Voice Connector. |
| `RequireEncryption` | require_encryption | `boolean` | required |  | Enables or disables encryption for the Voice Connector. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the Voice Connector. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  | The Voice Connector updated timestamp, in ISO 8601 format. |
| `VoiceConnectorArn` | voice_connector_arn | `string` | computed |  | The ARN of the Voice Connector. |
| `VoiceConnectorId` | voice_connector_id | `string` | computed |  | The Voice Connector ID. |

Supports update: yes

Discovery: supported
