# aws.securitygroupvpcassociation

**CloudFormation type:** `AWS::EC2::SecurityGroupVpcAssociation`

Resource type definition for the AWS::EC2::SecurityGroupVpcAssociation resource

Region attribute: `region`

**Import ID:** `<region>/GroupId|VpcId` (AWS::EC2::SecurityGroupVpcAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupId` | group_id | `string` | required, replaces on change |  | The group ID of the specified security group. |
| `State` |  | `string` | computed |  | The state of the security group vpc association. |
| `StateReason` | state_reason | `string` | computed |  | The reason for the state of the security group vpc association. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC in the security group vpc association. |
| `VpcOwnerId` | vpc_owner_id | `string` | computed |  | The owner of the VPC in the security group vpc association. |

Supports update: no

Discovery: supported
