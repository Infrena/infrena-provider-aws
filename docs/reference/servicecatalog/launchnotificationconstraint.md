# aws.launchnotificationconstraint

**CloudFormation type:** `AWS::ServiceCatalog::LaunchNotificationConstraint`

Resource Type definition for AWS::ServiceCatalog::LaunchNotificationConstraint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::LaunchNotificationConstraint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  | Unique identifier for the constraint |
| `NotificationArns` | notification_arns | `list` | required |  |  |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id |  |
| `ProductId` | product_id | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
