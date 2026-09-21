# aws.idmappingtable

**CloudFormation type:** `AWS::CleanRooms::IdMappingTable`

Represents an association between an ID mapping workflow and a collaboration

Region attribute: `region`

**Import ID:** `<region>/IdMappingTableIdentifier|MembershipIdentifier` (AWS::CleanRooms::IdMappingTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IdMappingTableIdentifier` | id_mapping_table_identifier | `string` | computed |  |  |
| `InputReferenceConfig` | input_reference_config | `map` | required, replaces on change |  |  |
| `InputReferenceProperties` | input_reference_properties | `map` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  |  |
| `MembershipArn` | membership_arn | `string` | computed |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
