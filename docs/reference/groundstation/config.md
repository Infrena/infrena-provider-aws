# aws.config

**CloudFormation type:** `AWS::GroundStation::Config`

AWS Ground Station config resource type for CloudFormation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::GroundStation::Config)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ConfigData` | config_data | `map` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | computed |  |  |

Supports update: yes

Discovery: supported
