# aws.ippoolroutetableassociation

**CloudFormation type:** `AWS::EC2::IpPoolRouteTableAssociation`

Resource Type definition for AWS::EC2::IpPoolRouteTableAssociation

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::EC2::IpPoolRouteTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | The route table association ID. |
| `PublicIpv4Pool` | public_ipv4_pool | `string` | required, replaces on change |  | The ID of the public IPv4 pool. |
| `RouteTableId` | route_table_id | `string` | required, replaces on change | aws.routetable.RouteTableId | The ID of the route table. |

Supports update: no

Discovery: supported
