# aws.cloudhsm.cluster

**CloudFormation type:** `AWS::CloudHSM::Cluster`

Creates and manages an AWS CloudHSM cluster.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudHSM::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the cluster. |
| `BackupPolicy` | backup_policy | `string` | computed |  | The cluster's backup policy. |
| `BackupRetentionPolicy` | backup_retention_policy | `map` | optional, computed, provider-chosen |  | A policy that defines how the service retains backups. |
| `ClusterId` | cluster_id | `string` | computed |  | The cluster's identifier (ID). |
| `HsmType` | hsm_type | `string` | required, replaces on change |  | The type of HSM to use in the cluster. |
| `Mode` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The mode to use in the cluster. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The NetworkType to create a cluster with. |
| `SecurityGroup` | security_group | `string` | computed |  | The identifier (ID) of the cluster's security group. |
| `State` |  | `string` | computed |  | The cluster's state. |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, replaces on change, write-only | aws.subnet.SubnetId | The identifiers (IDs) of the subnets where the cluster is created. You must specify at least one subnet. |
| `SubnetMapping` | subnet_mapping | `map` | computed |  | A map from availability zone to the cluster's subnet in that availability zone. |
| `Tags` |  | `map` | tags map |  | Tags to apply to the CloudHSM cluster. |
| `VpcId` | vpc_id | `string` | computed |  | The identifier (ID) of the virtual private cloud (VPC) that contains the cluster. |

Supports update: yes

Discovery: supported
