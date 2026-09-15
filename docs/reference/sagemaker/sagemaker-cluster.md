# aws.sagemaker.cluster

**CloudFormation type:** `AWS::SageMaker::Cluster`

Resource Type definition for AWS::SageMaker::Cluster

Region attribute: `region`

**Import ID:** `<region>/ClusterArn` (AWS::SageMaker::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScaling` | auto_scaling | `map` | optional, computed, provider-chosen |  | Configuration for cluster auto-scaling |
| `ClusterArn` | cluster_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the HyperPod Cluster. |
| `ClusterName` | cluster_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the HyperPod Cluster. |
| `ClusterRole` | cluster_role | `string` | optional, computed, provider-chosen |  | The cluster role for the autoscaler to assume. |
| `ClusterStatus` | cluster_status | `string` | computed |  | The status of the HyperPod Cluster. |
| `CreationTime` | creation_time | `string` | computed |  | The time at which the HyperPod cluster was created. |
| `FailureMessage` | failure_message | `string` | computed |  | The failure message of the HyperPod Cluster. |
| `InstanceGroups` | instance_groups | `list` | optional, computed, provider-chosen |  | The instance groups of the SageMaker HyperPod cluster. |
| `NodeProvisioningMode` | node_provisioning_mode | `string` | optional, computed, provider-chosen |  | Determines the scaling strategy for the SageMaker HyperPod cluster. When set to 'Continuous', enables continuous scaling which dynamically manages node provisioning. If the parameter is omitted, uses the standard scaling approach in previous release. |
| `NodeRecovery` | node_recovery | `string` | optional, computed, provider-chosen |  | If node auto-recovery is set to true, faulty nodes will be replaced or rebooted when a failure is detected. If set to false, nodes will be labelled when a fault is detected. |
| `Orchestrator` |  | `map` | optional, computed, provider-chosen |  | Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster or Slurm configuration. |
| `RestrictedInstanceGroups` | restricted_instance_groups | `list` | optional, computed, provider-chosen |  | The restricted instance groups of the SageMaker HyperPod cluster. |
| `RestrictedInstanceGroupsConfig` | restricted_instance_groups_config | `map` | optional, computed, provider-chosen |  | The cluster-level configuration for restricted instance groups, including shared environment settings for inter-RIG communication and FSx Lustre sharing. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Custom tags for managing the SageMaker HyperPod cluster as an AWS resource. You can add tags to your cluster in the same way you add them in other AWS services that support tagging. |
| `TieredStorageConfig` | tiered_storage_config | `map` | optional, computed, provider-chosen |  | Configuration for tiered storage in the SageMaker HyperPod cluster. |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. |

Supports update: yes

Discovery: supported
