# aws.podidentityassociation

**CloudFormation type:** `AWS::EKS::PodIdentityAssociation`

An object representing an Amazon EKS PodIdentityAssociation.

Region attribute: `region`

**Import ID:** `<region>/AssociationArn` (AWS::EKS::PodIdentityAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationArn` | association_arn | `string` | computed |  | The ARN of the pod identity association. |
| `AssociationId` | association_id | `string` | computed |  | The ID of the pod identity association. |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | The cluster that the pod identity association is created for. |
| `DisableSessionTags` | disable_session_tags | `boolean` | optional, computed, provider-chosen |  | The Disable Session Tags of the pod identity association. |
| `ExternalId` | external_id | `string` | computed |  | The External Id of the pod identity association. |
| `Namespace` |  | `string` | required, replaces on change |  | The Kubernetes namespace that the pod identity association is created for. |
| `Policy` |  | `string` | optional, computed, provider-chosen |  | The policy of the pod identity association. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The IAM role ARN that the pod identity association is created for. |
| `ServiceAccount` | service_account | `string` | required, replaces on change |  | The Kubernetes service account that the pod identity association is created for. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetRoleArn` | target_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Target Role Arn of the pod identity association. |

Supports update: yes

Discovery: supported (parent resource required)
