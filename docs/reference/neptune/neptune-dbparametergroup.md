# aws.neptune.dbparametergroup

**CloudFormation type:** `AWS::Neptune::DBParameterGroup`

AWS::Neptune::DBParameterGroup creates a new DB parameter group. This type can be declared in a template and referenced in the DBParameterGroupName parameter of AWS::Neptune::DBInstance

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Neptune::DBParameterGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | required, replaces on change |  | Provides the customer-specified description for this DB parameter group. |
| `Family` |  | `string` | required, replaces on change |  | Must be `neptune1` for engine versions prior to 1.2.0.0, or `neptune1.2` for engine version `1.2.0.0` and higher. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Provides the name of the DB parameter group. |
| `Parameters` |  | `map` | required |  | The parameters to set for this DB parameter group. |
| `Tags` |  | `map` | tags map |  | An optional array of key-value pairs to apply to this DB parameter group. |

Supports update: yes

Discovery: supported
