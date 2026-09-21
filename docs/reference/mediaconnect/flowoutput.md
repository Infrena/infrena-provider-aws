# aws.flowoutput

**CloudFormation type:** `AWS::MediaConnect::FlowOutput`

Resource schema for AWS::MediaConnect::FlowOutput

Region attribute: `region`

**Import ID:** `<region>/OutputArn` (AWS::MediaConnect::FlowOutput)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CidrAllowList` | cidr_allow_list | `list` | optional, computed, provider-chosen |  | The range of IP addresses that should be allowed to initiate output requests to this flow. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the output. |
| `Destination` |  | `string` | optional, computed, provider-chosen |  | The address where you want to send the output. |
| `Encryption` |  | `map` | optional, computed, provider-chosen |  | Information about the encryption of the flow. |
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.mediaconnect.flow.FlowArn | The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow. |
| `MaxLatency` | max_latency | `integer` | optional, computed, provider-chosen |  | The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams. |
| `MediaStreamOutputConfigurations` | media_stream_output_configurations | `list` | optional, computed, provider-chosen |  | The definition for each media stream that is associated with the output. |
| `MinLatency` | min_latency | `integer` | optional, computed, provider-chosen |  | The minimum latency in milliseconds. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the output. This value must be unique within the current flow. |
| `NdiOutputTimecodeSource` | ndi_output_timecode_source | `string` | optional, computed, provider-chosen |  | The timecode source for the NDI output. |
| `NdiProgramName` | ndi_program_name | `string` | optional, computed, provider-chosen |  | A suffix for the names of the NDI sources that the flow creates. If a custom name isn't specified, MediaConnect uses the output name. |
| `NdiSpeedHqQuality` | ndi_speed_hq_quality | `integer` | optional, computed, provider-chosen |  | A quality setting for the NDI Speed HQ encoder. |
| `OutputArn` | output_arn | `string` | computed |  | The ARN of the output. |
| `OutputStatus` | output_status | `string` | optional, computed, provider-chosen |  | An indication of whether the output should transmit data or not. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port to use when content is distributed to this output. |
| `Protocol` |  | `string` | optional, computed, provider-chosen |  | The protocol that is used by the source or output. |
| `RemoteId` | remote_id | `string` | optional, computed, provider-chosen |  | The remote ID for the Zixi-pull stream. |
| `RouterIntegrationState` | router_integration_state | `string` | optional, computed, provider-chosen |  |  |
| `RouterIntegrationTransitEncryption` | router_integration_transit_encryption | `map` | optional, computed, provider-chosen |  | The configuration that defines how content is encrypted during transit between the MediaConnect router and a MediaConnect flow. |
| `SmoothingLatency` | smoothing_latency | `integer` | optional, computed, provider-chosen |  | The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams. |
| `StreamId` | stream_id | `string` | optional, computed, provider-chosen |  | The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to tag and organize this flow output. |
| `VpcInterfaceAttachment` | vpc_interface_attachment | `map` | optional, computed, provider-chosen |  | The settings for attaching a VPC interface to an output. |

Supports update: yes

Discovery: supported (parent resource required)
