# aws.formtype

**CloudFormation type:** `AWS::DataZone::FormType`

Create and manage form types in Amazon Datazone

Region attribute: `region`

**Import ID:** `<region>/DomainIdentifier|FormTypeIdentifier` (AWS::DataZone::FormType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when this Amazon DataZone metadata form type was created. |
| `CreatedBy` | created_by | `string` | computed |  | The user who created this Amazon DataZone metadata form type. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of this Amazon DataZone metadata form type. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the Amazon DataZone domain in which this metadata form type is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change |  | The ID of the Amazon DataZone domain in which this metadata form type is created. |
| `FormTypeIdentifier` | form_type_identifier | `string` | computed |  | The ID of this Amazon DataZone metadata form type. |
| `Model` |  | `map` | required |  | Indicates the smithy model of the API. |
| `Name` |  | `string` | required, replaces on change |  | The name of this Amazon DataZone metadata form type. |
| `OwningProjectId` | owning_project_id | `string` | computed |  | The ID of the project that owns this Amazon DataZone metadata form type. |
| `OwningProjectIdentifier` | owning_project_identifier | `string` | required |  | The ID of the Amazon DataZone project that owns this metadata form type. |
| `Revision` |  | `string` | computed |  | The revision of this Amazon DataZone metadata form type. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of this Amazon DataZone metadata form type. |

Supports update: yes

Discovery: not supported
