# aws.qbusiness.index

**CloudFormation type:** `AWS::QBusiness::Index`

Definition of AWS::QBusiness::Index Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|IndexId` (AWS::QBusiness::Index)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `CapacityConfiguration` | capacity_configuration | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `DocumentAttributeConfigurations` | document_attribute_configurations | `list` | optional, computed, provider-chosen |  |  |
| `IndexArn` | index_arn | `string` | computed |  |  |
| `IndexId` | index_id | `string` | computed |  |  |
| `IndexStatistics` | index_statistics | `map` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
