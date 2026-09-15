# aws.mediapackagev2.channel

**CloudFormation type:** `AWS::MediaPackageV2::Channel`

<p>Represents an entry point into AWS Elemental MediaPackage for an ABR video content stream sent from an upstream encoder such as AWS Elemental MediaLive. The channel continuously analyzes the content that it receives and prepares it to be distributed to consumers via one or more origin endpoints.</p>

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaPackageV2::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) associated with the resource.</p> |
| `ChannelGroupName` | channel_group_name | `string` | required, replaces on change |  |  |
| `ChannelName` | channel_name | `string` | required, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  | <p>The date and time the channel was created.</p> |
| `Description` |  | `string` | optional, computed, provider-chosen |  | <p>Enter any descriptive text that helps you to identify the channel.</p> |
| `IngestEndpointUrls` | ingest_endpoint_urls | `list` | computed |  |  |
| `IngestEndpoints` | ingest_endpoints | `list` | computed |  | <p>The list of ingest endpoints.</p> |
| `InputSwitchConfiguration` | input_switch_configuration | `map` | optional, computed, provider-chosen |  | <p>The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.</p> |
| `InputType` | input_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  | <p>The date and time the channel was modified.</p> |
| `OutputHeaderConfiguration` | output_header_configuration | `map` | optional, computed, provider-chosen |  | <p>The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.</p> |
| `OutputLockingMode` | output_locking_mode | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
