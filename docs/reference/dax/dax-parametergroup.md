# aws.dax.parametergroup

**CloudFormation type:** `AWS::DAX::ParameterGroup`

Resource Type definition for AWS::DAX::ParameterGroup

Region attribute: `region`

**Import ID:** `<region>/ParameterGroupName` (AWS::DAX::ParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the parameter group. |
| `ParameterGroupName` | parameter_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the parameter group. |
| `ParameterNameValues` | parameter_name_values | `map` | optional, computed, provider-chosen |  | An array of name-value pairs for the parameters in the group. Each element in the array represents a single parameter. |

Supports update: yes

Discovery: supported
