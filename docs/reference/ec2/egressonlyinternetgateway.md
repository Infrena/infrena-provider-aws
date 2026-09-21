# aws.egressonlyinternetgateway

**CloudFormation type:** `AWS::EC2::EgressOnlyInternetGateway`

Resource Type definition for AWS::EC2::EgressOnlyInternetGateway

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::EgressOnlyInternetGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | Service Generated ID of the EgressOnlyInternetGateway |
| `Tags` |  | `map` | tags map |  | Any tags assigned to the egress only internet gateway. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC for which to create the egress-only internet gateway. |

Supports update: yes

Discovery: supported
