# aws.launchroleconstraint

**CloudFormation type:** `AWS::ServiceCatalog::LaunchRoleConstraint`

Resource Type definition for AWS::ServiceCatalog::LaunchRoleConstraint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::LaunchRoleConstraint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code for the constraint. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the launch role constraint. |
| `Id` |  | `string` | computed |  | The unique identifier for the launch role constraint. |
| `LocalRoleName` | local_role_name | `string` | optional, computed, provider-chosen |  | The local IAM role name to use in the launch constraint. |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The ID of the portfolio to which this launch role constraint applies. |
| `ProductId` | product_id | `string` | required, replaces on change |  | The ID of the product to which this launch role constraint applies. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the IAM role used for the launch constraint. |

Supports update: yes

Discovery: supported (parent resource required)
