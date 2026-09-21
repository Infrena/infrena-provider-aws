# aws.attributegroupassociation

**CloudFormation type:** `AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation`

Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn|AttributeGroupArn` (AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Application` |  | `string` | required, replaces on change |  | The name or the Id of the Application. |
| `ApplicationArn` | application_arn | `string` | computed |  |  |
| `AttributeGroup` | attribute_group | `string` | required, replaces on change |  | The name or the Id of the AttributeGroup. |
| `AttributeGroupArn` | attribute_group_arn | `string` | computed |  |  |

Supports update: no

Discovery: supported (parent resource required)
