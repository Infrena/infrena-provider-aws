# aws.qbusiness.permission

**CloudFormation type:** `AWS::QBusiness::Permission`

Definition of AWS::QBusiness::Permission Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|StatementId` (AWS::QBusiness::Permission)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required, replaces on change |  |  |
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `Conditions` |  | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Principal` |  | `string` | required, replaces on change |  |  |
| `StatementId` | statement_id | `string` | required, replaces on change |  |  |

Supports update: no

Discovery: supported (parent resource required)
