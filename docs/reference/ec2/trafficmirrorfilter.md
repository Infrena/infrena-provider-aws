# aws.trafficmirrorfilter

**CloudFormation type:** `AWS::EC2::TrafficMirrorFilter`

Resource schema for AWS::EC2::TrafficMirrorFilter

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TrafficMirrorFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of a traffic mirror filter. |
| `Id` |  | `string` | computed |  | The ID of a traffic mirror filter. |
| `NetworkServices` | network_services | `list` | optional, computed, provider-chosen |  | The network service that is associated with the traffic mirror filter. |
| `Tags` |  | `map` | tags map |  | The tags for a traffic mirror filter. |

Supports update: yes

Discovery: supported
