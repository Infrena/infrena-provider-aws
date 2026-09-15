# aws.vpcpeeringconnection

**CloudFormation type:** `AWS::EC2::VPCPeeringConnection`

Resource Type definition for AWS::EC2::VPCPeeringConnection

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::VPCPeeringConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssumeRoleRegion` | assume_role_region | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Region code to use when calling Security Token Service (STS) to assume the PeerRoleArn, if provided. |
| `Id` |  | `string` | computed |  |  |
| `PeerOwnerId` | peer_owner_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS account ID of the owner of the accepter VPC. |
| `PeerRegion` | peer_region | `string` | optional, computed, provider-chosen, replaces on change |  | The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request. |
| `PeerRoleArn` | peer_role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account. |
| `PeerVpcId` | peer_vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC with which you are creating the VPC peering connection. You must specify this parameter in the request. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |

Supports update: yes

Discovery: supported
