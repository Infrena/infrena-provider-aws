# aws.portfolioshare

**CloudFormation type:** `AWS::ServiceCatalog::PortfolioShare`

Resource Type definition for AWS::ServiceCatalog::PortfolioShare

Region attribute: `region`

**Import ID:** `<region>/PortfolioId|AccountId` (AWS::ServiceCatalog::PortfolioShare)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code. |
| `AccountId` | account_id | `string` | required, replaces on change |  | The AWS account ID. |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The portfolio identifier. |
| `ShareTagOptions` | share_tag_options | `boolean` | optional, computed, provider-chosen |  | Enables or disables TagOptions sharing when creating the portfolio share. |

Supports update: yes

Discovery: supported (parent resource required)
