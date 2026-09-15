# aws.hooktypeconfig

**CloudFormation type:** `AWS::CloudFormation::HookTypeConfig`

Specifies the configuration data for a registered hook in CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/ConfigurationArn` (AWS::CloudFormation::HookTypeConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Configuration` |  | `string` | optional, computed, provider-chosen |  | The configuration data for the extension, in this account and region. |
| `ConfigurationAlias` | configuration_alias | `string` | optional, computed, provider-chosen, replaces on change |  | An alias by which to refer to this extension configuration data. |
| `ConfigurationArn` | configuration_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the configuration data, in this account and region. |
| `TypeArn` | type_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the type without version number. |
| `TypeName` | type_name | `string` | optional, computed, provider-chosen |  | The name of the type being registered. |

Supports update: yes

Discovery: supported
