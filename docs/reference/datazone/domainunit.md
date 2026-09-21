# aws.domainunit

**CloudFormation type:** `AWS::DataZone::DomainUnit`

A domain unit enables you to easily organize your assets and other domain entities under specific business units and teams.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::DomainUnit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp at which the domain unit was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the domain unit. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the domain where the domain unit was created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change |  | The ID of the domain where you want to create a domain unit. |
| `Id` |  | `string` | computed |  | The ID of the domain unit. |
| `Identifier` |  | `string` | computed |  | The identifier of the domain unit that you want to get. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp at which the domain unit was last updated. |
| `Name` |  | `string` | required |  | The name of the domain unit. |
| `ParentDomainUnitId` | parent_domain_unit_id | `string` | computed |  | The ID of the parent domain unit. |
| `ParentDomainUnitIdentifier` | parent_domain_unit_identifier | `string` | required, replaces on change |  | The ID of the parent domain unit. |

Supports update: yes

Discovery: supported (parent resource required)
