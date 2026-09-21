# aws.securitygroupegress

**CloudFormation type:** `AWS::EC2::SecurityGroupEgress`

Adds the specified outbound (egress) rule to a security group.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SecurityGroupEgress)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CidrIp` | cidr_ip | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 address range, in CIDR format. |
| `CidrIpv6` | cidr_ipv6 | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv6 address range, in CIDR format. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of an egress (outbound) security group rule. |
| `DestinationPrefixListId` | destination_prefix_list_id | `string` | optional, computed, provider-chosen, replaces on change | aws.prefixlist.PrefixListId | The prefix list IDs for an AWS service. This is the AWS service to access through a VPC endpoint from instances associated with the security group. |
| `DestinationSecurityGroupId` | destination_security_group_id | `string` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The ID of the security group. |
| `FromPort` | from_port | `integer` | optional, computed, provider-chosen, replaces on change |  | If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types). |
| `GroupId` | group_id | `string` | required, replaces on change |  | The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID. |
| `Id` |  | `string` | computed |  |  |
| `IpProtocol` | ip_protocol | `string` | required, replaces on change |  | The IP protocol name (``tcp``, ``udp``, ``icmp``, ``icmpv6``) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)). |
| `ToPort` | to_port | `integer` | optional, computed, provider-chosen, replaces on change |  | If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes). |

Supports update: yes

Discovery: supported
