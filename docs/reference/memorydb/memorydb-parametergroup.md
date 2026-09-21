# aws.memorydb.parametergroup

**CloudFormation type:** `AWS::MemoryDB::ParameterGroup`

The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.

Region attribute: `region`

**Import ID:** `<region>/ParameterGroupName` (AWS::MemoryDB::ParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARN` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the parameter group. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the parameter group. |
| `Family` |  | `string` | required, replaces on change |  | The name of the parameter group family that this parameter group is compatible with. |
| `ParameterGroupName` | parameter_group_name | `string` | required, replaces on change |  | The name of the parameter group. |
| `Parameters` |  | `map` | optional, computed, provider-chosen, write-only |  | An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this parameter group. |

Supports update: yes

Discovery: supported
