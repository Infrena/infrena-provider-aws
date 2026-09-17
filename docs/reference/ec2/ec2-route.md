# aws.ec2.route

**CloudFormation type:** `AWS::EC2::Route`

Specifies a route in a route table. For more information, see [Routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#route-table-routes) in the *Amazon VPC User Guide*.

Region attribute: `region`

**Import ID:** `<region>/RouteTableId|CidrBlock` (AWS::EC2::Route)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CarrierGatewayId` | carrier_gateway_id | `string` | optional, computed, provider-chosen | aws.carriergateway.CarrierGatewayId | The ID of the carrier gateway. |
| `CidrBlock` | cidr_block | `string` | computed |  |  |
| `CoreNetworkArn` | core_network_arn | `string` | optional, computed, provider-chosen | aws.corenetwork.CoreNetworkArn | The Amazon Resource Name (ARN) of the core network. |
| `DestinationCidrBlock` | destination_cidr, destination_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv4 CIDR address block used for the destination match. Routing decisions are based on the most specific match. We modify the specified CIDR block to its canonical form; for example, if you specify ``100.68.0.18/18``, we modify it to ``100.68.0.0/18``. |
| `DestinationIpv6CidrBlock` | destination_ipv6_cidr_block | `string` | optional, computed, provider-chosen, replaces on change |  | The IPv6 CIDR block used for the destination match. Routing decisions are based on the most specific match. |
| `DestinationPrefixListId` | destination_prefix_list_id | `string` | optional, computed, provider-chosen, replaces on change | aws.prefixlist.PrefixListId | The ID of a prefix list used for the destination match. |
| `EgressOnlyInternetGatewayId` | egress_only_internet_gateway_id | `string` | optional, computed, provider-chosen | aws.egressonlyinternetgateway.Id | [IPv6 traffic only] The ID of an egress-only internet gateway. |
| `GatewayId` | gateway_id | `string` | optional, computed, provider-chosen |  | The ID of an internet gateway or virtual private gateway attached to your VPC. |
| `InstanceId` | instance_id | `string` | optional, computed, provider-chosen | aws.ec2.instance.InstanceId | The ID of a NAT instance in your VPC. The operation fails if you specify an instance ID unless exactly one network interface is attached. |
| `LocalGatewayId` | local_gateway_id | `string` | optional, computed, provider-chosen |  | The ID of the local gateway. |
| `NatGatewayId` | nat_gateway_id | `string` | optional, computed, provider-chosen | aws.natgateway.NatGatewayId | [IPv4 traffic only] The ID of a NAT gateway. |
| `NetworkInterfaceId` | network_interface_id | `string` | optional, computed, provider-chosen | aws.networkinterface.Id | The ID of a network interface. |
| `OdbNetworkArn` | odb_network_arn | `string` | optional, computed, provider-chosen | aws.odbnetwork.OdbNetworkArn | The Amazon Resource Name (ARN) of the ODB network. |
| `RouteTableId` | route_table_id | `string` | required, replaces on change | aws.routetable.RouteTableId | The ID of the route table for the route. |
| `TransitGatewayId` | transit_gateway_id | `string` | optional, computed, provider-chosen | aws.transitgateway.Id | The ID of a transit gateway. |
| `VpcEndpointId` | vpc_endpoint_id | `string` | optional, computed, provider-chosen | aws.ec2.vpcendpoint.Id | The ID of a VPC endpoint. Supported for Gateway Load Balancer endpoints only. |
| `VpcPeeringConnectionId` | vpc_peering_connection_id | `string` | optional, computed, provider-chosen | aws.vpcpeeringconnection.Id | The ID of a VPC peering connection. |

Supports update: yes

Discovery: supported (parent resource required)
