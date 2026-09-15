# aws.natgateway

**CloudFormation type:** `AWS::EC2::NatGateway`

Specifies a network address translation (NAT) gateway in the specified subnet. You can create either a public NAT gateway or a private NAT gateway. The default is a public NAT gateway. If you create a public NAT gateway, you must specify an elastic IP address.

Region attribute: `region`

**Import ID:** `<region>/NatGatewayId` (AWS::EC2::NatGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllocationId` | allocation_id | `string` | optional, computed, provider-chosen, replaces on change |  | [Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway. |
| `AutoProvisionZones` | auto_provision_zones | `string` | computed |  |  |
| `AutoScalingIps` | auto_scaling_ips | `string` | computed |  |  |
| `AvailabilityMode` | availability_mode | `string` | optional, computed, provider-chosen, replaces on change |  | Indicates whether this is a zonal (single-AZ) or regional (multi-AZ) NAT gateway. |
| `AvailabilityZoneAddresses` | availability_zone_addresses | `list` | optional, computed, provider-chosen |  | For regional NAT gateways only: Specifies which Availability Zones you want the NAT gateway to support and the Elastic IP addresses (EIPs) to use in each AZ. The regional NAT gateway uses these EIPs to handle outbound NAT traffic from their respective AZs. If not specified, the NAT gateway will automatically expand to new AZs and associate EIPs upon detection of an elastic network interface. If you specify this parameter, auto-expansion is disabled and you must manually manage AZ coverage. |
| `ConnectivityType` | connectivity_type | `string` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity. |
| `EniId` | eni_id | `string` | computed |  |  |
| `MaxDrainDurationSeconds` | max_drain_duration_seconds | `integer` | optional, computed, provider-chosen, write-only |  | The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds. |
| `NatGatewayId` | nat_gateway_id | `string` | computed |  |  |
| `PrivateIpAddress` | private_ip_address | `string` | optional, computed, provider-chosen, replaces on change |  | The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned. |
| `RouteTableId` | route_table_id | `string` | computed |  |  |
| `SecondaryAllocationIds` | secondary_allocation_ids | `list` | optional, computed, provider-chosen |  | Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*. |
| `SecondaryPrivateIpAddressCount` | secondary_private_ip_address_count | `integer` | optional, computed, provider-chosen |  | [Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*. |
| `SecondaryPrivateIpAddresses` | secondary_private_ip_addresses | `list` | optional, computed, provider-chosen |  | Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*. |
| `SubnetId` | subnet_id | `string` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | The ID of the subnet in which the NAT gateway is located. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the NAT gateway. |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId | The ID of the VPC in which the NAT gateway is located. |

Supports update: yes

Discovery: supported
