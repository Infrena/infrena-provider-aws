# aws.eks.cluster

**CloudFormation type:** `AWS::EKS::Cluster`

An object representing an Amazon EKS cluster.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::EKS::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessConfig` | access_config | `map` | optional, computed, provider-chosen |  | An object representing the Access Config to use for the cluster. |
| `ActiveCertificateAuthorityId` | active_certificate_authority_id | `string` | optional, computed, provider-chosen, write-only | aws.eks.certificateauthority.Id | The ID of the certificate authority to activate as the cluster's signing CA. Setting or changing this value activates the specified CA (the previously active CA becomes trusted). |
| `Arn` |  | `string` | computed |  | The ARN of the cluster, such as arn:aws:eks:us-west-2:666666666666:cluster/prod. |
| `BootstrapSelfManagedAddons` | bootstrap_self_managed_addons | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Set this value to false to avoid creating the default networking add-ons when the cluster is created. |
| `CertificateAuthority` | certificate_authority | `map` | optional, computed, provider-chosen |  | The certificate authority information for the cluster, including the trust bundle and the currently active signing certificate authority. |
| `CertificateAuthorityData` | certificate_authority_data | `string` | computed |  | The certificate-authority-data for your cluster. |
| `ClusterSecurityGroupId` | cluster_security_group_id | `string` | computed |  | The cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control plane to data plane communication. |
| `ComputeConfig` | compute_config | `map` | optional, computed, provider-chosen |  | Todo: add description |
| `ControlPlaneScalingConfig` | control_plane_scaling_config | `map` | optional, computed, provider-chosen |  | Configuration for provisioned control plane scaling. |
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | Set this value to true to enable deletion protection for the cluster. |
| `EncryptionConfig` | encryption_config | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `EncryptionConfigKeyArn` | encryption_config_key_arn | `string` | computed |  | Amazon Resource Name (ARN) or alias of the customer master key (CMK). |
| `Endpoint` |  | `string` | computed |  | The endpoint for your Kubernetes API server, such as https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com. |
| `Force` |  | `boolean` | optional, computed, provider-chosen, write-only |  | Force cluster version update |
| `Id` |  | `string` | computed |  | The unique ID given to your cluster. |
| `KubeApiServerConfig` | kube_api_server_config | `map` | optional, computed, provider-chosen |  | The configuration for the Kubernetes API server on an Amazon EKS cluster. |
| `KubeControllerManagerConfig` | kube_controller_manager_config | `map` | optional, computed, provider-chosen |  | The configuration for the Kubernetes controller manager on an Amazon EKS cluster. |
| `KubeSchedulerConfig` | kube_scheduler_config | `map` | optional, computed, provider-chosen |  | The configuration for the Kubernetes scheduler on an Amazon EKS cluster. |
| `KubernetesNetworkConfig` | kubernetes_network_config | `map` | optional, computed, provider-chosen |  | The Kubernetes network configuration for the cluster. |
| `Logging` |  | `map` | optional, computed, provider-chosen |  | Enable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs based on log types. By default, cluster control plane logs aren't exported to CloudWatch Logs. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name to give to your cluster. |
| `OpenIdConnectIssuerUrl` | open_id_connect_issuer_url | `string` | computed |  | The issuer URL for the cluster's OIDC identity provider, such as https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E. If you need to remove https:// from this output value, you can include the following code in your template. |
| `OutpostConfig` | outpost_config | `map` | optional, computed, provider-chosen, replaces on change |  | An object representing the Outpost configuration to use for AWS EKS outpost cluster. |
| `RemoteNetworkConfig` | remote_network_config | `map` | optional, computed, provider-chosen |  | Configuration fields for specifying on-premises node and pod CIDRs that are external to the VPC passed during cluster creation. |
| `ResourcesVpcConfig` | resources_vpc_config | `map` | required |  | An object representing the VPC configuration to use for an Amazon EKS cluster. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. |
| `RollbackConfig` | rollback_config | `map` | optional, computed, provider-chosen, write-only |  | The rollback configuration to use for the cluster version rollback. |
| `StorageConfig` | storage_config | `map` | optional, computed, provider-chosen |  | Todo: add description |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpgradePolicy` | upgrade_policy | `map` | optional, computed, provider-chosen |  | An object representing the Upgrade Policy to use for the cluster. |
| `Version` |  | `string` | optional, computed, provider-chosen |  | The desired Kubernetes version for your cluster. If you don't specify a value here, the latest version available in Amazon EKS is used. |
| `ZonalShiftConfig` | zonal_shift_config | `map` | optional, computed, provider-chosen |  | The current zonal shift configuration to use for the cluster. |

Supports update: yes

Discovery: supported
