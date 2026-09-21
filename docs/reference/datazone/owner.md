# aws.owner

**CloudFormation type:** `AWS::DataZone::Owner`

A owner can set up authorization permissions on their resources.

Region attribute: `region`

**Import ID:** `<region>/DomainIdentifier|EntityType|EntityIdentifier|OwnerType|OwnerIdentifier` (AWS::DataZone::Owner)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change |  | The ID of the domain in which you want to add the entity owner. |
| `EntityIdentifier` | entity_identifier | `string` | required, replaces on change |  | The ID of the entity to which you want to add an owner. |
| `EntityType` | entity_type | `string` | required, replaces on change |  | The type of an entity. |
| `Owner` |  | `map` | required, replaces on change, write-only |  | The properties of a domain unit's owner. |
| `OwnerIdentifier` | owner_identifier | `string` | computed |  |  |
| `OwnerType` | owner_type | `string` | computed |  |  |

Supports update: no

Discovery: supported (parent resource required)
