# aws.docdb.dbclusterparametergroup

**CloudFormation type:** `AWS::DocDB::DBClusterParameterGroup`

Resource Type definition for AWS::DocDB::DBClusterParameterGroup

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DocDB::DBClusterParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | required, replaces on change |  | The description for the DB cluster parameter group. |
| `Family` |  | `string` | required, replaces on change |  | The DB cluster parameter group family name (e.g. docdb5.0). |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the DB cluster parameter group. If omitted, CloudFormation generates a unique name. The name is stored as lowercase. |
| `Parameters` |  | `map` | required |  | An object containing key-value pairs of parameters to set for the DB cluster parameter group. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
