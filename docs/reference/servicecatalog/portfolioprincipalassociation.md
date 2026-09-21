# aws.portfolioprincipalassociation

**CloudFormation type:** `AWS::ServiceCatalog::PortfolioPrincipalAssociation`

Resource Type definition for AWS::ServiceCatalog::PortfolioPrincipalAssociation

Region attribute: `region`

**Import ID:** `<region>/PortfolioId|PrincipalARN` (AWS::ServiceCatalog::PortfolioPrincipalAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The language code. |
| `PortfolioId` | portfolio_id | `string` | optional, computed, provider-chosen, replaces on change | aws.portfolio.Id | The portfolio identifier. |
| `PrincipalARN` | principal_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the principal (user, role, or group). |
| `PrincipalType` | principal_type | `string` | required, replaces on change |  | The principal type. The supported value is IAM if you use a fully defined Amazon Resource Name (ARN), or IAM_PATTERN if you use an ARN with no accountID, with or without wildcard characters. |

Supports update: no

Discovery: supported (parent resource required)
