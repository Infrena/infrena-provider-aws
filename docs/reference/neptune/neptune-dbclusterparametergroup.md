# aws.neptune.dbclusterparametergroup

**CloudFormation type:** `AWS::Neptune::DBClusterParameterGroup`

The AWS::Neptune::DBClusterParameterGroup resource creates a new Amazon Neptune DB cluster parameter group

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Neptune::DBClusterParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | required, replaces on change |  | Provides the customer-specified description for this DB cluster parameter group. |
| `Family` |  | `string` | required, replaces on change |  | Must be neptune1 for engine versions prior to 1.2.0.0, or neptune1.2 for engine version 1.2.0.0 and higher. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Provides the name of the DB cluster parameter group. |
| `Parameters` |  | `map` | required |  | An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The list of tags for the cluster parameter group. |

Supports update: yes

Discovery: supported
