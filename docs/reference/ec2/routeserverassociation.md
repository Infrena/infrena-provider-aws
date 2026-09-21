# aws.routeserverassociation

**CloudFormation type:** `AWS::EC2::RouteServerAssociation`

VPC Route Server Association

Region attribute: `region`

**Import ID:** `<region>/RouteServerId|VpcId` (AWS::EC2::RouteServerAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RouteServerId` | route_server_id | `string` | required, replaces on change | aws.routeserver.Id | Route Server ID |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | VPC ID |

Supports update: no

Discovery: supported
