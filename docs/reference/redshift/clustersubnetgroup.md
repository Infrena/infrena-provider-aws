# aws.clustersubnetgroup

**CloudFormation type:** `AWS::Redshift::ClusterSubnetGroup`

Resource Type definition for AWS::Redshift::ClusterSubnetGroup. Specifies an Amazon Redshift subnet group.

Region attribute: `region`

**Import ID:** `<region>/ClusterSubnetGroupName` (AWS::Redshift::ClusterSubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterSubnetGroupName` | cluster_subnet_group_name | `string` | computed |  | This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default". |
| `Description` |  | `string` | required |  | The description of the parameter group. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The list of VPC subnet IDs |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The list of tags for the cluster parameter group. |

Supports update: yes

Discovery: supported
