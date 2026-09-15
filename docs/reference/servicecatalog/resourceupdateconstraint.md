# aws.resourceupdateconstraint

**CloudFormation type:** `AWS::ServiceCatalog::ResourceUpdateConstraint`

Resource type definition for AWS::ServiceCatalog::ResourceUpdateConstraint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::ResourceUpdateConstraint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the constraint |
| `Id` |  | `string` | computed |  | Unique identifier for the constraint |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The portfolio identifier |
| `ProductId` | product_id | `string` | required, replaces on change |  | The product identifier |
| `TagUpdateOnProvisionedProduct` | tag_update_on_provisioned_product | `string` | required |  | ALLOWED or NOT_ALLOWED, to permit or prevent changes to the tags on provisioned instances of the specified portfolio / product combination |

Supports update: yes

Discovery: supported (parent resource required)
