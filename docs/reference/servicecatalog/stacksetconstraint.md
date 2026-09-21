# aws.stacksetconstraint

**CloudFormation type:** `AWS::ServiceCatalog::StackSetConstraint`

Resource Type definition for AWS::ServiceCatalog::StackSetConstraint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::StackSetConstraint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code. |
| `AccountList` | account_list | `list` | required |  | One or more AWS accounts that will have access to the provisioned product. |
| `AdminRole` | admin_role | `string` | required |  | AdminRole ARN. |
| `Description` |  | `string` | required |  | The description of the constraint. |
| `ExecutionRole` | execution_role | `string` | required |  | ExecutionRole name. |
| `Id` |  | `string` | computed |  | Unique identifier for the constraint |
| `PortfolioId` | portfolio_id | `string` | required, replaces on change | aws.portfolio.Id | The portfolio identifier. |
| `ProductId` | product_id | `string` | required, replaces on change |  | The product identifier. |
| `RegionList` | region_list | `list` | required |  | One or more AWS Regions where the provisioned product will be available. |
| `StackInstanceControl` | stack_instance_control | `string` | required |  | Permission to create, update, and delete stack instances. Choose from ALLOWED and NOT_ALLOWED. |

Supports update: yes

Discovery: supported (parent resource required)
