# aws.datazone.project

**CloudFormation type:** `AWS::DataZone::Project`

Amazon DataZone projects are business use case–based groupings of people, assets (data), and tools used to simplify access to the AWS analytics.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the project was created. |
| `CreatedBy` | created_by | `string` | computed |  | The Amazon DataZone user who created the project. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Amazon DataZone project. |
| `DomainId` | domain_id | `string` | computed |  | The identifier of the Amazon DataZone domain in which the project was created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The ID of the Amazon DataZone domain in which this project is created. |
| `DomainUnitId` | domain_unit_id | `string` | optional, computed, provider-chosen | aws.domainunit.Id | The ID of the domain unit. |
| `GlossaryTerms` | glossary_terms | `list` | optional, computed, provider-chosen |  | The glossary terms that can be used in this Amazon DataZone project. |
| `Id` |  | `string` | computed |  | The ID of the Amazon DataZone project. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the project was last updated. |
| `MembershipAssignments` | membership_assignments | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | The project membership assignments. |
| `Name` |  | `string` | required |  | The name of the Amazon DataZone project. |
| `ProjectCategory` | project_category | `string` | optional, computed, provider-chosen, replaces on change |  | The project category. |
| `ProjectExecutionRole` | project_execution_role | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The project execution role ARN. |
| `ProjectProfileId` | project_profile_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.projectprofile.Id | The project profile ID. |
| `ProjectProfileVersion` | project_profile_version | `string` | optional, computed, provider-chosen, write-only |  | The project profile version to which the project should be updated. You can only specify the following string for this parameter: latest. |
| `ProjectStatus` | project_status | `string` | computed |  | The status of the project. |
| `ResourceTags` | resource_tags | `list` | optional, computed, provider-chosen |  | The resource tags of the project. |
| `UserParameters` | user_parameters | `list` | optional, computed, provider-chosen, write-only |  | The user parameters of the project. |

Supports update: yes

Discovery: supported (parent resource required)
