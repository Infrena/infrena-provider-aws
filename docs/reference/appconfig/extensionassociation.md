# aws.extensionassociation

**CloudFormation type:** `AWS::AppConfig::ExtensionAssociation`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AppConfig::ExtensionAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ExtensionArn` | extension_arn | `string` | computed |  |  |
| `ExtensionIdentifier` | extension_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ExtensionVersionNumber` | extension_version_number | `integer` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  |  |
| `ResourceArn` | resource_arn | `string` | computed |  |  |
| `ResourceIdentifier` | resource_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
