# aws.eks.capability

**CloudFormation type:** `AWS::EKS::Capability`

Resource Type definition for EKS Capability.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::EKS::Capability)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the capability. |
| `CapabilityName` | capability_name | `string` | required, replaces on change |  | A unique name for the capability. The name must be unique within your cluster and can contain alphanumeric characters, hyphens, and underscores. |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | The name of the EKS cluster where you want to create the capability. |
| `Configuration` |  | `map` | optional, computed, provider-chosen |  | Configuration settings for a capability. The structure of this object varies depending on the capability type. |
| `CreatedAt` | created_at | `string` | computed |  | The Unix epoch timestamp in seconds for when the capability was created. |
| `DeletePropagationPolicy` | delete_propagation_policy | `string` | required |  | Specifies how Kubernetes resources managed by the capability should be handled when the capability is deleted. Currently, the only supported value is RETAIN which retains all Kubernetes resources managed by the capability when the capability is deleted. |
| `ModifiedAt` | modified_at | `string` | computed |  | The Unix epoch timestamp in seconds for when the capability was last modified. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services. This role must have a trust policy that allows the EKS service principal to assume it, and it must have the necessary permissions for the capability type you're creating. |
| `Status` |  | `string` | computed |  | The current status of the capability. Valid values include: CREATING (the capability is being created), ACTIVE (the capability is running and available), UPDATING (the capability is being updated), DELETING (the capability is being deleted), CREATE_FAILED (the capability creation failed), UPDATE_FAILED (the capability update failed), or DELETE_FAILED (the capability deletion failed). |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of capability to create. Valid values are: ACK (AWS Controllers for Kubernetes, which lets you manage AWS resources directly from Kubernetes), ARGOCD (Argo CD for GitOps-based continuous delivery), or KRO (Kube Resource Orchestrator for composing and managing custom Kubernetes resources). |
| `Version` |  | `string` | computed |  | The version of the capability software that is currently running. |

Supports update: yes

Discovery: supported (parent resource required)
