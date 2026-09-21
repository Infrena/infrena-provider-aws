# aws.routetable

**CloudFormation type:** `AWS::EC2::RouteTable`

Specifies a route table for the specified VPC. After you create a route table, you can add routes and associate the table with a subnet.

Region attribute: `region`

**Import ID:** `<region>/RouteTableId` (AWS::EC2::RouteTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RouteTableId` | route_table_id | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | Any tags assigned to the route table. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: yes

Discovery: supported
