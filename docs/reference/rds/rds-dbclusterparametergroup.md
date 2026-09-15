# aws.rds.dbclusterparametergroup

**CloudFormation type:** `AWS::RDS::DBClusterParameterGroup`

The ``AWS::RDS::DBClusterParameterGroup`` resource creates a new Amazon RDS DB cluster parameter group.

Region attribute: `region`

**Import ID:** `<region>/DBClusterParameterGroupName` (AWS::RDS::DBClusterParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBClusterParameterGroupName` | db_cluster_parameter_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the DB cluster parameter group. |
| `Description` |  | `string` | required, replaces on change |  | The description for the DB cluster parameter group. |
| `Family` |  | `string` | required, replaces on change |  | The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family. |
| `Parameters` |  | `map` | required |  | Provides a list of parameters for the DB cluster parameter group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the DB cluster parameter group. |

Supports update: yes

Discovery: supported
