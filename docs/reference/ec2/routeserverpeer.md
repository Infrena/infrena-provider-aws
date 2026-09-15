# aws.routeserverpeer

**CloudFormation type:** `AWS::EC2::RouteServerPeer`

VPC Route Server Peer

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::RouteServerPeer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Route Server Peer. |
| `BgpOptions` | bgp_options | `map` | required, replaces on change |  | BGP Options |
| `EndpointEniAddress` | endpoint_eni_address | `string` | computed |  | Elastic Network Interface IP address owned by the Route Server Endpoint |
| `EndpointEniId` | endpoint_eni_id | `string` | computed |  | Elastic Network Interface ID owned by the Route Server Endpoint |
| `Id` |  | `string` | computed |  | The ID of the Route Server Peer. |
| `PeerAddress` | peer_address | `string` | required, replaces on change |  | IP address of the Route Server Peer |
| `RouteServerEndpointId` | route_server_endpoint_id | `string` | required, replaces on change | aws.routeserverendpoint.Id | Route Server Endpoint ID |
| `RouteServerId` | route_server_id | `string` | computed |  | Route Server ID |
| `SubnetId` | subnet_id | `string` | computed |  | Subnet ID |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcId` | vpc_id | `string` | computed |  | VPC ID |

Supports update: yes

Discovery: supported
