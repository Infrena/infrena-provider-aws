# aws.anycastiplist

**CloudFormation type:** `AWS::CloudFront::AnycastIpList`

An Anycast static IP list. For more information, see [Request Anycast static IPs to use for allowlisting](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/request-static-ips.html) in the *Amazon CloudFront Developer Guide*.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::AnycastIpList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnycastIpList` | anycast_ip_list | `map` | computed |  | An Anycast static IP list. For more information, see [Request Anycast static IPs to use for allowlisting](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/request-static-ips.html) in the *Amazon CloudFront Developer Guide*. |
| `ETag` | e_tag | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | The IP address type for the Anycast static IP list. |
| `IpCount` | ip_count | `integer` | required, replaces on change |  | The number of IP addresses in the Anycast static IP list. |
| `IpamCidrConfigResults` | ipam_cidr_config_results | `list` | computed |  |  |
| `IpamCidrConfigs` | ipam_cidr_configs | `list` | optional, computed, provider-chosen, write-only |  | A list of IPAM CIDR configurations that define the IP address ranges, IPAM pools, and associated Anycast IP addresses. |
| `Name` |  | `string` | required, replaces on change |  | The name of the Anycast static IP list. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
