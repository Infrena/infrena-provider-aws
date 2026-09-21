# aws.routeserverpropagation

**CloudFormation type:** `AWS::EC2::RouteServerPropagation`

VPC Route Server Propagation

Region attribute: `region`

**Import ID:** `<region>/RouteServerId|RouteTableId` (AWS::EC2::RouteServerPropagation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RouteServerId` | route_server_id | `string` | required, replaces on change | aws.routeserver.Id | Route Server ID |
| `RouteTableId` | route_table_id | `string` | required, replaces on change | aws.routetable.RouteTableId | Route Table ID |

Supports update: no

Discovery: supported
