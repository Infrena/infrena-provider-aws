# aws.projectprofile

**CloudFormation type:** `AWS::DataZone::ProjectProfile`

Definition of AWS::DataZone::ProjectProfile Resource Type

Region attribute: `region`

**Import ID:** `<region>/DomainIdentifier|Identifier` (AWS::DataZone::ProjectProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowCustomProjectResourceTags` | allow_custom_project_resource_tags | `boolean` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CreatedBy` | created_by | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DomainId` | domain_id | `string` | computed |  |  |
| `DomainIdentifier` | domain_identifier | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainUnitId` | domain_unit_id | `string` | computed |  |  |
| `DomainUnitIdentifier` | domain_unit_identifier | `string` | optional, computed, provider-chosen, write-only |  |  |
| `EnvironmentConfigurations` | environment_configurations | `list` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Identifier` |  | `string` | computed |  |  |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `ProjectResourceTags` | project_resource_tags | `list` | optional, computed, provider-chosen |  |  |
| `ProjectResourceTagsDescription` | project_resource_tags_description | `string` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `UseDefaultConfigurations` | use_default_configurations | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  |  |

Supports update: yes

Discovery: supported
