# aws.dhcpoptions

**CloudFormation type:** `AWS::EC2::DHCPOptions`

Resource Type definition for AWS::EC2::DHCPOptions

Region attribute: `region`

**Import ID:** `<region>/DhcpOptionsId` (AWS::EC2::DHCPOptions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DhcpOptionsId` | dhcp_options_id | `string` | computed |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen, replaces on change |  | This value is used to complete unqualified DNS hostnames. |
| `DomainNameServers` | domain_name_servers | `list` | optional, computed, provider-chosen, replaces on change |  | The IPv4 addresses of up to four domain name servers, or AmazonProvidedDNS. |
| `Ipv6AddressPreferredLeaseTime` | ipv6_address_preferred_lease_time | `integer` | optional, computed, provider-chosen, replaces on change |  | The preferred Lease Time for ipV6 address in seconds. |
| `NetbiosNameServers` | netbios_name_servers | `list` | optional, computed, provider-chosen, replaces on change |  | The IPv4 addresses of up to four NetBIOS name servers. |
| `NetbiosNodeType` | netbios_node_type | `integer` | optional, computed, provider-chosen, replaces on change |  | The NetBIOS node type (1, 2, 4, or 8). |
| `NtpServers` | ntp_servers | `list` | optional, computed, provider-chosen, replaces on change |  | The IPv4 addresses of up to four Network Time Protocol (NTP) servers. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the DHCP options set. |

Supports update: yes

Discovery: supported
