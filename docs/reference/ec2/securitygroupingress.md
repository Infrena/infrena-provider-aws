# aws.securitygroupingress

**CloudFormation type:** `AWS::EC2::SecurityGroupIngress`

Resource Type definition for AWS::EC2::SecurityGroupIngress

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SecurityGroupIngress)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CidrIp` | cidr_ip | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 ranges |
| `CidrIpv6` | cidr_ipv6 | `string` | optional, computed, provider-chosen, replaces on change |  | [VPC only] The IPv6 ranges |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Updates the description of an ingress (inbound) security group rule. You can replace an existing description, or add a description to a rule that did not have one previously |
| `FromPort` | from_port | `integer` | optional, computed, provider-chosen, replaces on change |  | The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes. |
| `GroupId` | group_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID. |
| `GroupName` | group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the security group. |
| `Id` |  | `string` | computed |  | The Security Group Rule Id |
| `IpProtocol` | ip_protocol | `string` | required, replaces on change |  | The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers). |
| `SourcePrefixListId` | source_prefix_list_id | `string` | optional, computed, provider-chosen, replaces on change | aws.prefixlist.PrefixListId | [EC2-VPC only] The ID of a prefix list. |
| `SourceSecurityGroupId` | source_security_group_id | `string` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The ID of the security group. You must specify either the security group ID or the security group name. For security groups in a nondefault VPC, you must specify the security group ID. |
| `SourceSecurityGroupName` | source_security_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | [EC2-Classic, default VPC] The name of the source security group. |
| `SourceSecurityGroupOwnerId` | source_security_group_owner_id | `string` | optional, computed, provider-chosen, replaces on change |  | [nondefault VPC] The AWS account ID that owns the source security group. You can't specify this property with an IP address range. |
| `ToPort` | to_port | `integer` | optional, computed, provider-chosen, replaces on change |  | The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you specify all ICMP/ICMPv6 types, you must specify all codes. |

Supports update: yes

Discovery: supported
