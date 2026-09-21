# aws.attributegroup

**CloudFormation type:** `AWS::ServiceCatalogAppRegistry::AttributeGroup`

Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalogAppRegistry::AttributeGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Attributes` |  | `map` | required |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the attribute group. |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required |  | The name of the attribute group. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
