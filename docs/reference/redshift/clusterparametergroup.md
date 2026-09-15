# aws.clusterparametergroup

**CloudFormation type:** `AWS::Redshift::ClusterParameterGroup`

Resource Type definition for AWS::Redshift::ClusterParameterGroup

Region attribute: `region`

**Import ID:** `<region>/ParameterGroupName` (AWS::Redshift::ClusterParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | required, replaces on change |  | A description of the parameter group. |
| `ParameterGroupFamily` | parameter_group_family | `string` | required, replaces on change |  | The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters. |
| `ParameterGroupName` | parameter_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the cluster parameter group. |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  | An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
