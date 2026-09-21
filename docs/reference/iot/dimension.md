# aws.dimension

**CloudFormation type:** `AWS::IoT::Dimension`

A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoT::Dimension)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN (Amazon resource name) of the created dimension. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A unique identifier for the dimension. |
| `StringValues` | string_values | `list` | required |  | Specifies the value or list of values for the dimension. |
| `Tags` |  | `map` | tags map |  | Metadata that can be used to manage the dimension. |
| `Type` | type_value | `string` | required, replaces on change |  | Specifies the type of the dimension. |

Supports update: yes

Discovery: supported
