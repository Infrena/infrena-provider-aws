# aws.idnamespaceassociation

**CloudFormation type:** `AWS::CleanRooms::IdNamespaceAssociation`

Represents an association between an ID namespace and a collaboration

Region attribute: `region`

**Import ID:** `<region>/IdNamespaceAssociationIdentifier|MembershipIdentifier` (AWS::CleanRooms::IdNamespaceAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IdMappingConfig` | id_mapping_config | `map` | optional, computed, provider-chosen |  |  |
| `IdNamespaceAssociationIdentifier` | id_namespace_association_identifier | `string` | computed |  |  |
| `InputReferenceConfig` | input_reference_config | `map` | required, replaces on change |  |  |
| `InputReferenceProperties` | input_reference_properties | `map` | computed |  |  |
| `MembershipArn` | membership_arn | `string` | computed |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
