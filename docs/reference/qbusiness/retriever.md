# aws.retriever

**CloudFormation type:** `AWS::QBusiness::Retriever`

Definition of AWS::QBusiness::Retriever Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|RetrieverId` (AWS::QBusiness::Retriever)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `Configuration` |  | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `RetrieverArn` | retriever_arn | `string` | computed |  |  |
| `RetrieverId` | retriever_id | `string` | computed |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
