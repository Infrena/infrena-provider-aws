# aws.acceptedportfolioshare

**CloudFormation type:** `AWS::ServiceCatalog::AcceptedPortfolioShare`

Resource Type definition for AWS::ServiceCatalog::AcceptedPortfolioShare

Region attribute: `region`

**Import ID:** `<region>/PortfolioId` (AWS::ServiceCatalog::AcceptedPortfolioShare)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The language code. |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The portfolio identifier. |

Supports update: no

Discovery: supported
