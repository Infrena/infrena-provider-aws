# aws.trafficmirrorfilterrule

**CloudFormation type:** `AWS::EC2::TrafficMirrorFilterRule`

Resource Type definition for for AWS::EC2::TrafficMirrorFilterRule

Region attribute: `region`

**Import ID:** `<region>/TrafficMirrorFilterRuleId` (AWS::EC2::TrafficMirrorFilterRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Traffic Mirror Filter rule. |
| `DestinationCidrBlock` | destination_cidr_block | `string` | required |  | The destination CIDR block to assign to the Traffic Mirror rule. |
| `DestinationPortRange` | destination_port_range | `map` | optional, computed, provider-chosen |  | The destination port range. |
| `Protocol` |  | `integer` | optional, computed, provider-chosen |  | The number of protocol, for example 17 (UDP), to assign to the Traffic Mirror rule. |
| `RuleAction` | rule_action | `string` | required |  | The action to take on the filtered traffic. |
| `RuleNumber` | rule_number | `integer` | required |  | The number of the Traffic Mirror rule. |
| `SourceCidrBlock` | source_cidr_block | `string` | required |  | The source CIDR block to assign to the Traffic Mirror Filter rule. |
| `SourcePortRange` | source_port_range | `map` | optional, computed, provider-chosen |  | The source port range. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the Traffic Mirror Filter rule. |
| `TrafficDirection` | traffic_direction | `string` | required |  | The type of traffic. |
| `TrafficMirrorFilterId` | traffic_mirror_filter_id | `string` | required, replaces on change | aws.trafficmirrorfilter.Id | The ID of the filter that this rule is associated with. |
| `TrafficMirrorFilterRuleId` | traffic_mirror_filter_rule_id | `string` | computed |  | The ID of the Traffic Mirror Filter rule. |

Supports update: yes

Discovery: supported
