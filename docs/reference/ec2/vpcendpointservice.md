# aws.vpcendpointservice

**CloudFormation type:** `AWS::EC2::VPCEndpointService`

Resource Type definition for AWS::EC2::VPCEndpointService

Region attribute: `region`

**Import ID:** `<region>/ServiceId` (AWS::EC2::VPCEndpointService)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptanceRequired` | acceptance_required | `boolean` | optional, computed, provider-chosen |  |  |
| `ContributorInsightsEnabled` | contributor_insights_enabled | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `GatewayLoadBalancerArns` | gateway_load_balancer_arns | `list` | optional, computed, provider-chosen |  |  |
| `NetworkLoadBalancerArns` | network_load_balancer_arns | `list` | optional, computed, provider-chosen |  |  |
| `PayerResponsibility` | payer_responsibility | `string` | optional, computed, provider-chosen |  |  |
| `PrivateDnsName` | private_dns_name | `string` | optional, computed, provider-chosen |  |  |
| `PrivateDnsNameConfiguration` | private_dns_name_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ServiceId` | service_id | `string` | computed |  |  |
| `SupportedIpAddressTypes` | supported_ip_address_types | `list` | optional, computed, provider-chosen |  | Specify which Ip Address types are supported for VPC endpoint service. |
| `SupportedRegions` | supported_regions | `list` | optional, computed, provider-chosen |  | The Regions from which service consumers can access the service. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to add to the VPC endpoint service. |

Supports update: yes

Discovery: supported
