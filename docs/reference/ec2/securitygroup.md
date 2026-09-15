# aws.securitygroup

**CloudFormation type:** `AWS::EC2::SecurityGroup`

Resource Type definition for AWS::EC2::SecurityGroup

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::SecurityGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GroupDescription` | description, group_description | `string` | required, replaces on change |  | A description for the security group. |
| `GroupId` | group_id | `string` | computed |  | The group ID of the specified security group. |
| `GroupName` | name, group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the security group. |
| `Id` |  | `string` | computed |  | The group name or group ID depending on whether the SG is created in default or specific VPC |
| `SecurityGroupEgress` | egress, security_group_egress | `list` | optional, computed, provider-chosen |  | [VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. |
| `SecurityGroupIngress` | ingress, security_group_ingress | `list` | optional, computed, provider-chosen |  | The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the security group. |
| `VpcId` | vpc, vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId | The ID of the VPC for the security group. |

Supports update: yes

Discovery: supported
