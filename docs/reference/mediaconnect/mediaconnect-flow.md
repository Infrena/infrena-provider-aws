# aws.mediaconnect.flow

**CloudFormation type:** `AWS::MediaConnect::Flow`

Resource Type definition for AWS::MediaConnect::Flow

Region attribute: `region`

**Import ID:** `<region>/FlowArn` (AWS::MediaConnect::Flow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS. |
| `EgressIp` | egress_ip | `string` | computed |  | The IP address from which video will be sent to output destinations. |
| `EncodingConfig` | encoding_config | `map` | optional, computed, provider-chosen |  | The encoding configuration to apply to the NDI source content when transcoding it to a transport stream (TS) for downstream distribution. You can choose between several predefined encoding profiles based on common use cases. |
| `FlowArn` | flow_arn | `string` | computed |  | The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow. |
| `FlowAvailabilityZone` | flow_availability_zone | `string` | computed |  | The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly) |
| `FlowNdiMachineName` | flow_ndi_machine_name | `string` | computed |  | A prefix for the names of the NDI sources that the flow creates.(ReadOnly) |
| `FlowSize` | flow_size | `string` | optional, computed, provider-chosen |  | Determines the processing capacity and feature set of the flow. Set this optional parameter to LARGE if you want to enable NDI sources or outputs on the flow. |
| `Maintenance` |  | `map` | optional, computed, provider-chosen |  | The maintenance setting of a flow. |
| `MediaStreams` | media_streams | `list` | optional, computed, provider-chosen |  | The media streams associated with the flow. You can associate any of these media streams with sources and outputs on the flow. |
| `Name` |  | `string` | required, replaces on change |  | The name of the flow. |
| `NdiConfig` | ndi_config | `map` | optional, computed, provider-chosen |  | Specifies the configuration settings for NDI sources and outputs. Required when the flow includes NDI sources or outputs. |
| `Source` |  | `map` | required |  | The settings for the source of the flow. |
| `SourceFailoverConfig` | source_failover_config | `map` | optional, computed, provider-chosen |  | The settings for source failover |
| `SourceMonitoringConfig` | source_monitoring_config | `map` | optional, computed, provider-chosen |  | The settings for source monitoring. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to tag this flow. |
| `VpcInterfaces` | vpc_interfaces | `list` | optional, computed, provider-chosen |  | The VPC interfaces that you added to this flow. |

Supports update: yes

Discovery: supported
