# aws.trafficmirrorsession

**CloudFormation type:** `AWS::EC2::TrafficMirrorSession`

Resource schema for AWS::EC2::TrafficMirrorSession

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TrafficMirrorSession)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Traffic Mirror session. |
| `Id` |  | `string` | computed |  | The ID of a Traffic Mirror session. |
| `NetworkInterfaceId` | network_interface_id | `string` | required | aws.networkinterface.Id | The ID of the source network interface. |
| `OwnerId` | owner_id | `string` | optional, computed, provider-chosen |  | The ID of the account that owns the Traffic Mirror session. |
| `PacketLength` | packet_length | `integer` | optional, computed, provider-chosen |  | The number of bytes in each packet to mirror. |
| `SessionNumber` | session_number | `integer` | required |  | The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets. |
| `Tags` |  | `map` | tags map |  | The tags assigned to the Traffic Mirror session. |
| `TrafficMirrorFilterId` | traffic_mirror_filter_id | `string` | required | aws.trafficmirrorfilter.Id | The ID of a Traffic Mirror filter. |
| `TrafficMirrorTargetId` | traffic_mirror_target_id | `string` | required | aws.trafficmirrortarget.Id | The ID of a Traffic Mirror target. |
| `VirtualNetworkId` | virtual_network_id | `integer` | optional, computed, provider-chosen |  | The VXLAN ID for the Traffic Mirror session. |

Supports update: yes

Discovery: supported
