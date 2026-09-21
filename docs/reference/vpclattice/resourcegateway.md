# aws.resourcegateway

**CloudFormation type:** `AWS::VpcLattice::ResourceGateway`

Creates a resource gateway for a service.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ResourceGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Ipv4AddressesPerEni` | ipv4_addresses_per_eni | `integer` | optional, computed, provider-chosen |  | The number of IPv4 addresses to allocate per ENI for the resource gateway |
| `Name` |  | `string` | required, replaces on change |  |  |
| `ResourceConfigDnsResolution` | resource_config_dns_resolution | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | The ID of one or more security groups to associate with the endpoint network interface. |
| `SubnetIds` | subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId | The ID of one or more subnets in which to create an endpoint network interface. |
| `Tags` |  | `map` | tags map |  |  |
| `VpcIdentifier` | vpc_identifier | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
