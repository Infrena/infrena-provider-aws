# aws.routeserverendpoint

**CloudFormation type:** `AWS::EC2::RouteServerEndpoint`

VPC Route Server Endpoint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::RouteServerEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Route Server Endpoint. |
| `EniAddress` | eni_address | `string` | computed |  | Elastic Network Interface IP address owned by the Route Server Endpoint |
| `EniId` | eni_id | `string` | computed |  | Elastic Network Interface ID owned by the Route Server Endpoint |
| `Id` |  | `string` | computed |  | The ID of the Route Server Endpoint. |
| `RouteServerId` | route_server_id | `string` | required, replaces on change | aws.routeserver.Id | Route Server ID |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | Subnet ID |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcId` | vpc_id | `string` | computed |  | VPC ID |

Supports update: yes

Discovery: supported
