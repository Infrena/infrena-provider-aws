# aws.msk.vpcconnection

**CloudFormation type:** `AWS::MSK::VpcConnection`

Resource Type definition for AWS::MSK::VpcConnection

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MSK::VpcConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Authentication` |  | `string` | required, replaces on change |  | The type of private link authentication |
| `ClientSubnets` | client_subnets | `list` | required, replaces on change |  |  |
| `SecurityGroups` | security_groups | `list` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `TargetClusterArn` | target_cluster_arn | `string` | required, replaces on change | aws.msk.cluster.Arn | The Amazon Resource Name (ARN) of the target cluster |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
