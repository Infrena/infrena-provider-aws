# aws.rds.dbparametergroup

**CloudFormation type:** `AWS::RDS::DBParameterGroup`

The ``AWS::RDS::DBParameterGroup`` resource creates a custom parameter group for an RDS database family.

Region attribute: `region`

**Import ID:** `<region>/DBParameterGroupName` (AWS::RDS::DBParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBParameterGroupArn` | db_parameter_group_arn | `string` | computed |  |  |
| `DBParameterGroupName` | db_parameter_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the DB parameter group. |
| `Description` |  | `string` | required, replaces on change |  | Provides the customer-specified description for this DB parameter group. |
| `Family` |  | `string` | required, replaces on change |  | The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | A mapping of parameter names and values for the parameter update. You must specify at least one parameter name and value. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the DB parameter group. |

Supports update: yes

Discovery: supported
