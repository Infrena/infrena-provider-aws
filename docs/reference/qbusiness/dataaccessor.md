# aws.dataaccessor

**CloudFormation type:** `AWS::QBusiness::DataAccessor`

Definition of AWS::QBusiness::DataAccessor Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|DataAccessorId` (AWS::QBusiness::DataAccessor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionConfigurations` | action_configurations | `list` | required |  |  |
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `AuthenticationDetail` | authentication_detail | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DataAccessorArn` | data_accessor_arn | `string` | computed |  |  |
| `DataAccessorId` | data_accessor_id | `string` | computed |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `IdcApplicationArn` | idc_application_arn | `string` | computed |  |  |
| `Principal` |  | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
