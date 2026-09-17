# aws.portfolioproductassociation

**CloudFormation type:** `AWS::ServiceCatalog::PortfolioProductAssociation`

Resource Type definition for AWS::ServiceCatalog::PortfolioProductAssociation

Region attribute: `region`

**Import ID:** `<region>/PortfolioId|ProductId` (AWS::ServiceCatalog::PortfolioProductAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The language code. |
| `PortfolioId` | portfolio_id | `string` | optional, computed, provider-chosen, replaces on change | aws.portfolio.Id | The portfolio identifier. |
| `ProductId` | product_id | `string` | optional, computed, provider-chosen, replaces on change |  | The product identifier. |
| `SourcePortfolioId` | source_portfolio_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.portfolio.Id | The identifier of the source portfolio. The source portfolio must be a portfolio imported from a different account than the one creating the association. This account must have previously shared this portfolio with the account creating the association. |

Supports update: no

Discovery: supported (parent resource required)
