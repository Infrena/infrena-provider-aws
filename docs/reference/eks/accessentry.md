# aws.accessentry

**CloudFormation type:** `AWS::EKS::AccessEntry`

An object representing an Amazon EKS AccessEntry.

Region attribute: `region`

**Import ID:** `<region>/PrincipalArn|ClusterName` (AWS::EKS::AccessEntry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessEntryArn` | access_entry_arn | `string` | computed |  | The ARN of the access entry. |
| `AccessPolicies` | access_policies | `list` | optional, computed, provider-chosen |  | An array of access policies that are associated with the access entry. |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | The cluster that the access entry is created for. |
| `KubernetesGroups` | kubernetes_groups | `list` | optional, computed, provider-chosen |  | The Kubernetes groups that the access entry is associated with. |
| `PrincipalArn` | principal_arn | `string` | required, replaces on change |  | The principal ARN that the access entry is created for. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The node type to associate with the access entry. |
| `Username` |  | `string` | optional, computed, provider-chosen |  | The Kubernetes user that the access entry is associated with. |

Supports update: yes

Discovery: supported (parent resource required)
