# aws.flowsource

**CloudFormation type:** `AWS::MediaConnect::FlowSource`

Resource Type definition for AWS::MediaConnect::FlowSource

Region attribute: `region`

**Import ID:** `<region>/SourceArn` (AWS::MediaConnect::FlowSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Decryption` |  | `map` | optional, computed, provider-chosen |  | Information about the encryption of the flow. |
| `Description` |  | `string` | required |  | A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account. |
| `EntitlementArn` | entitlement_arn | `string` | optional, computed, provider-chosen |  | The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow. |
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.mediaconnect.flow.FlowArn | The ARN of the flow. |
| `GatewayBridgeSource` | gateway_bridge_source | `map` | optional, computed, provider-chosen |  | The source configuration for cloud flows receiving a stream from a bridge. |
| `IngestIp` | ingest_ip | `string` | computed |  | The IP address that the flow will be listening on for incoming content. |
| `IngestPort` | ingest_port | `integer` | optional, computed, provider-chosen |  | The port that the flow will be listening on for incoming content. |
| `MaxBitrate` | max_bitrate | `integer` | optional, computed, provider-chosen |  | The smoothing max bitrate for RIST, RTP, and RTP-FEC streams. |
| `MaxLatency` | max_latency | `integer` | optional, computed, provider-chosen |  | The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams. |
| `MinLatency` | min_latency | `integer` | optional, computed, provider-chosen |  | The minimum latency in milliseconds. |
| `Name` |  | `string` | required, replaces on change |  | The name of the source. |
| `Protocol` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The protocol that is used by the source. |
| `SenderControlPort` | sender_control_port | `integer` | optional, computed, provider-chosen |  | The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol. |
| `SenderIpAddress` | sender_ip_address | `string` | optional, computed, provider-chosen |  | The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol. |
| `SourceArn` | source_arn | `string` | computed |  | The ARN of the source. |
| `SourceIngestPort` | source_ingest_port | `string` | computed |  | The port that the flow will be listening on for incoming content.(ReadOnly) |
| `SourceListenerAddress` | source_listener_address | `string` | optional, computed, provider-chosen |  | Source IP or domain name for SRT-caller protocol. |
| `SourceListenerPort` | source_listener_port | `integer` | optional, computed, provider-chosen |  | Source port for SRT-caller protocol. |
| `StreamId` | stream_id | `string` | optional, computed, provider-chosen |  | The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to tag and organize this flow source. |
| `VpcInterfaceName` | vpc_interface_name | `string` | optional, computed, provider-chosen |  | The name of the VPC Interface this Source is configured with. |
| `WhitelistCidr` | whitelist_cidr | `string` | optional, computed, provider-chosen |  | The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. |

Supports update: yes

Discovery: supported (parent resource required)
