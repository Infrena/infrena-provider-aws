# aws.optiongroup

**CloudFormation type:** `AWS::RDS::OptionGroup`

The ``AWS::RDS::OptionGroup`` resource creates or updates an option group, to enable and configure features that are specific to a particular DB engine.

Region attribute: `region`

**Import ID:** `<region>/OptionGroupName` (AWS::RDS::OptionGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EngineName` | engine_name | `string` | required, replaces on change |  | Specifies the name of the engine that this option group should be associated with. |
| `MajorEngineVersion` | major_engine_version | `string` | required, replaces on change |  | Specifies the major version of the engine that this option group should be associated with. |
| `OptionConfigurations` | option_configurations | `list` | optional, computed, provider-chosen |  | A list of all available options for an option group. |
| `OptionGroupDescription` | option_group_description | `string` | required, replaces on change |  | The description of the option group. |
| `OptionGroupName` | option_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the option group to be created. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the option group. |

Supports update: yes

Discovery: supported
