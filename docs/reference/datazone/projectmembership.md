# aws.projectmembership

**CloudFormation type:** `AWS::DataZone::ProjectMembership`

Definition of AWS::DataZone::ProjectMembership Resource Type

Region attribute: `region`

**Import ID:** `<region>/DomainIdentifier|MemberIdentifier|MemberIdentifierType|ProjectIdentifier` (AWS::DataZone::ProjectMembership)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Designation` |  | `string` | required, write-only |  |  |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change |  |  |
| `Member` |  | `string` | required, replaces on change, write-only |  |  |
| `MemberIdentifier` | member_identifier | `string` | computed |  |  |
| `MemberIdentifierType` | member_identifier_type | `string` | computed |  |  |
| `ProjectIdentifier` | project_identifier | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
