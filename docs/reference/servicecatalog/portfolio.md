# aws.portfolio

**CloudFormation type:** `AWS::ServiceCatalog::Portfolio`

Resource type definition for AWS::ServiceCatalog::Portfolio

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::Portfolio)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `PortfolioName` | portfolio_name | `string` | computed |  |  |
| `ProviderName` | provider_name | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
