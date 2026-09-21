# aws.launchtemplateconstraint

**CloudFormation type:** `AWS::ServiceCatalog::LaunchTemplateConstraint`

Resource Type definition for AWS::ServiceCatalog::LaunchTemplateConstraint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::LaunchTemplateConstraint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the constraint. |
| `Id` |  | `string` | computed |  | Unique identifier for the constraint |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The portfolio identifier. |
| `ProductId` | product_id | `string` | required, replaces on change |  | The product identifier. |
| `Rules` |  | `string` | required |  | A json encoded string of the template constraint rules |

Supports update: yes

Discovery: supported (parent resource required)
