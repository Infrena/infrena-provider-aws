# aws.nodegroup

**CloudFormation type:** `AWS::EKS::Nodegroup`

Resource schema for AWS::EKS::Nodegroup

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EKS::Nodegroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmiType` | ami_type | `string` | optional, computed, provider-chosen, replaces on change |  | The AMI type for your node group. |
| `Arn` |  | `string` | computed |  |  |
| `CapacityType` | capacity_type | `string` | optional, computed, provider-chosen, replaces on change |  | The capacity type of your managed node group. |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | Name of the cluster to create the node group in. |
| `DiskSize` | disk_size | `integer` | optional, computed, provider-chosen, replaces on change |  | The root device disk size (in GiB) for your node group instances. |
| `ForceUpdateEnabled` | force_update_enabled | `boolean` | optional, computed, provider-chosen, write-only |  | Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue. |
| `Id` |  | `string` | computed |  |  |
| `InstanceTypes` | instance_types | `list` | optional, computed, provider-chosen, replaces on change |  | Specify the instance types for a node group. |
| `Labels` |  | `map` | optional, computed, provider-chosen |  | The Kubernetes labels to be applied to the nodes in the node group when they are created. |
| `LaunchTemplate` | launch_template | `map` | optional, computed, provider-chosen |  | An object representing a launch template specification for AWS EKS Nodegroup. |
| `NodeRepairConfig` | node_repair_config | `map` | optional, computed, provider-chosen |  | The node auto repair configuration for node group. |
| `NodeRole` | node_role | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the IAM role to associate with your node group. |
| `NodegroupName` | nodegroup_name | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name to give your node group. |
| `ReleaseVersion` | release_version | `string` | optional, computed, provider-chosen |  | The AMI version of the Amazon EKS-optimized AMI to use with your node group. |
| `RemoteAccess` | remote_access | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing a remote access configuration specification for AWS EKS Nodegroup. |
| `ScalingConfig` | scaling_config | `map` | optional, computed, provider-chosen |  | An object representing a auto scaling group specification for AWS EKS Nodegroup. |
| `Subnets` |  | `list` | required, replaces on change |  | The subnets to use for the Auto Scaling group that is created for your node group. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency. |
| `Taints` |  | `list` | optional, computed, provider-chosen |  | The Kubernetes taints to be applied to the nodes in the node group when they are created. |
| `UpdateConfig` | update_config | `map` | optional, computed, provider-chosen |  | The node group update configuration. |
| `Version` |  | `string` | optional, computed, provider-chosen |  | The Kubernetes version to use for your managed nodes. |
| `WarmPoolConfig` | warm_pool_config | `map` | optional, computed, provider-chosen |  | The warm pool configuration for the node group. |

Supports update: yes

Discovery: supported (parent resource required)
