# aws.networkacl

**CloudFormation type:** `AWS::EC2::NetworkAcl`

Specifies a network ACL for your VPC.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::NetworkAcl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the network ACL. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC for the network ACL. |

Supports update: yes

Discovery: supported
