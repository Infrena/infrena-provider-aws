# aws.addon

**CloudFormation type:** `AWS::EKS::Addon`

Resource Schema for AWS::EKS::Addon

Region attribute: `region`

**Import ID:** `<region>/ClusterName|AddonName` (AWS::EKS::Addon)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddonName` | addon_name | `string` | required, replaces on change |  | Name of Addon |
| `AddonVersion` | addon_version | `string` | optional, computed, provider-chosen |  | Version of Addon |
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the add-on |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | Name of Cluster |
| `ConfigurationValues` | configuration_values | `string` | optional, computed, provider-chosen |  | The configuration values to use with the add-on |
| `NamespaceConfig` | namespace_config | `map` | optional, computed, provider-chosen, replaces on change |  | The custom namespace configuration to use with the add-on |
| `PodIdentityAssociations` | pod_identity_associations | `list` | optional, computed, provider-chosen, write-only |  | An array of pod identities to apply to this add-on. |
| `PreserveOnDelete` | preserve_on_delete | `boolean` | optional, computed, provider-chosen, write-only |  | PreserveOnDelete parameter value |
| `ResolveConflicts` | resolve_conflicts | `string` | optional, computed, provider-chosen, write-only |  | Resolve parameter value conflicts |
| `ServiceAccountRoleArn` | service_account_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | IAM role to bind to the add-on's service account |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
