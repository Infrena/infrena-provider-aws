# aws.tagoption

**CloudFormation type:** `AWS::ServiceCatalog::TagOption`

Resource type definition for AWS::ServiceCatalog::TagOption

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::TagOption)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Active` |  | `boolean` | optional, computed, provider-chosen |  | The TagOption active state. |
| `Id` |  | `string` | computed |  | The TagOption identifier. |
| `Key` |  | `string` | required, replaces on change |  | The TagOption key. |
| `Value` |  | `string` | required, replaces on change |  | The TagOption value. |

Supports update: yes

Discovery: supported
