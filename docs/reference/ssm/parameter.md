# aws.parameter

**CloudFormation type:** `AWS::SSM::Parameter`

The ``AWS::SSM::Parameter`` resource creates an SSM parameter in SYSlong Parameter Store.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::SSM::Parameter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedPattern` | allowed_pattern | `string` | optional, computed, provider-chosen, write-only |  | A regular expression used to validate the parameter value. For example, for ``String`` types with values restricted to numbers, you can specify the following: ``AllowedPattern=^\d+$`` |
| `Arn` |  | `string` | computed |  |  |
| `DataType` | data_type | `string` | optional, computed, provider-chosen |  | The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``. |
| `Description` |  | `string` | optional, computed, provider-chosen, write-only |  | Information about the parameter. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the parameter. |
| `Policies` |  | `string` | optional, computed, provider-chosen, write-only |  | Information about the policies assigned to a parameter. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter. |
| `Tier` |  | `string` | optional, computed, provider-chosen, write-only |  | The parameter tier. |
| `Type` | type_value | `string` | required |  | The type of parameter. |
| `Value` |  | `string` | required |  | The parameter value. |

Supports update: yes

Discovery: supported
