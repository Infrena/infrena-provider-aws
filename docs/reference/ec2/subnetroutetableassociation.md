# aws.subnetroutetableassociation

**CloudFormation type:** `AWS::EC2::SubnetRouteTableAssociation`

Associates a subnet with a route table. The subnet and route table must be in the same VPC. This association causes traffic originating from the subnet to be routed according to the routes in the route table. A route table can be associated with multiple subnets. To create a route table, see [AWS::EC2::RouteTable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routetable.html).

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SubnetRouteTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `RouteTableId` | route_table_id | `string` | required, replaces on change | aws.routetable.RouteTableId | The ID of the route table. |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet. |

Supports update: no

Discovery: supported
