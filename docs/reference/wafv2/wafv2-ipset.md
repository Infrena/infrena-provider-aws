# aws.wafv2.ipset

**CloudFormation type:** `AWS::WAFv2::IPSet`

Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually

Region attribute: `region`

**Import ID:** `<region>/Name|Id|Scope` (AWS::WAFv2::IPSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Addresses` |  | `list` | required |  | List of IPAddresses. |
| `Arn` |  | `string` | computed |  | ARN of the WAF entity. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the entity. |
| `IPAddressVersion` | ip_address_version | `string` | required |  | Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address. |
| `Id` |  | `string` | computed |  | Id of the IPSet |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the IPSet. |
| `Scope` |  | `string` | required, replaces on change |  | Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
