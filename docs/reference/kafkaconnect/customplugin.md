# aws.customplugin

**CloudFormation type:** `AWS::KafkaConnect::CustomPlugin`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/CustomPluginArn` (AWS::KafkaConnect::CustomPlugin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContentType` | content_type | `string` | required, replaces on change |  | The type of the plugin file. |
| `CustomPluginArn` | custom_plugin_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the custom plugin to use. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A summary description of the custom plugin. |
| `FileDescription` | file_description | `map` | computed |  | Details about the custom plugin file. |
| `Location` |  | `map` | required, replaces on change |  | Information about the location of a custom plugin. |
| `Name` |  | `string` | required, replaces on change |  | The name of the custom plugin. |
| `Revision` |  | `integer` | computed |  | The revision of the custom plugin. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
