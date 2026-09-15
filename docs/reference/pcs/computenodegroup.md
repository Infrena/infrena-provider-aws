# aws.computenodegroup

**CloudFormation type:** `AWS::PCS::ComputeNodeGroup`

AWS::PCS::ComputeNodeGroup resource creates an AWS PCS compute node group.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::PCS::ComputeNodeGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmiId` | ami_id | `string` | optional, computed, provider-chosen |  | The ID of the Amazon Machine Image (AMI) that AWS PCS uses to launch instances. If not provided, AWS PCS uses the AMI ID specified in the custom launch template. |
| `Arn` |  | `string` | computed |  | The unique Amazon Resource Name (ARN) of the compute node group. |
| `ClusterId` | cluster_id | `string` | required, replaces on change | aws.pcs.cluster.Id | The ID of the cluster of the compute node group. |
| `CustomLaunchTemplate` | custom_launch_template | `map` | required |  | An Amazon EC2 launch template AWS PCS uses to launch compute nodes. |
| `ErrorInfo` | error_info | `list` | computed |  | The list of errors that occurred during compute node group provisioning. |
| `IamInstanceProfileArn` | iam_instance_profile_arn | `string` | required |  | The Amazon Resource Name (ARN) of the IAM instance profile used to pass an IAM role when launching EC2 instances. The role contained in your instance profile must have pcs:RegisterComputeNodeGroupInstance permissions attached to provision instances correctly. |
| `Id` |  | `string` | computed |  | The generated unique ID of the compute node group. |
| `InstanceConfigs` | instance_configs | `list` | required, replaces on change |  | A list of EC2 instance configurations that AWS PCS can provision in the compute node group. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name that identifies the compute node group. |
| `NodeLifecycleActions` | node_lifecycle_actions | `map` | optional, computed, provider-chosen |  | Custom scripts that run at defined points in a compute node's lifecycle. |
| `PurchaseOption` | purchase_option | `string` | optional, computed, provider-chosen |  | Specifies how EC2 instances are purchased on your behalf. AWS PCS supports On-Demand, Spot, Capacity Block, and Interruptible Capacity Reservation instances. For more information, see Instance purchasing options in the Amazon Elastic Compute Cloud User Guide. If you don't provide this option, it defaults to On-Demand. |
| `ScalingConfiguration` | scaling_configuration | `map` | required |  | Specifies the boundaries of the compute node group auto scaling. |
| `SlurmConfiguration` | slurm_configuration | `map` | optional, computed, provider-chosen |  | Additional options related to the Slurm scheduler. |
| `SpotOptions` | spot_options | `map` | optional, computed, provider-chosen |  | Additional configuration when you specify SPOT as the purchase option. |
| `Status` |  | `string` | computed |  | The provisioning status of the compute node group. The provisioning status doesn't indicate the overall health of the compute node group. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The list of subnet IDs where instances are provisioned by the compute node group. The subnets must be in the same VPC as the cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | 1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string. |

Supports update: yes

Discovery: supported (parent resource required)
