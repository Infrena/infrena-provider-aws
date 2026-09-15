# aws.multiplexprogram

**CloudFormation type:** `AWS::MediaLive::Multiplexprogram`

Resource schema for AWS::MediaLive::Multiplexprogram

Region attribute: `region`

**Import ID:** `<region>/ProgramName|MultiplexId` (AWS::MediaLive::Multiplexprogram)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelId` | channel_id | `string` | computed |  | The MediaLive channel associated with the program. |
| `MultiplexId` | multiplex_id | `string` | optional, computed, provider-chosen, replaces on change | aws.multiplex.Id | The ID of the multiplex that the program belongs to. |
| `MultiplexProgramSettings` | multiplex_program_settings | `map` | optional, computed, provider-chosen |  | Multiplex Program settings configuration. |
| `PacketIdentifiersMap` | packet_identifiers_map | `map` | optional, computed, provider-chosen |  | Packet identifiers map for a given Multiplex program. |
| `PipelineDetails` | pipeline_details | `list` | optional, computed, provider-chosen |  | Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time. |
| `PreferredChannelPipeline` | preferred_channel_pipeline | `string` | optional, computed, provider-chosen, write-only |  | Indicates which pipeline is preferred by the multiplex for program ingest. |
| `ProgramName` | program_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the multiplex program. |

Supports update: yes

Discovery: supported (parent resource required)
