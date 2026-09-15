# aws.configuredtableassociation

**CloudFormation type:** `AWS::CleanRooms::ConfiguredTableAssociation`

Represents a table that can be queried within a collaboration

Region attribute: `region`

**Import ID:** `<region>/ConfiguredTableAssociationIdentifier|MembershipIdentifier` (AWS::CleanRooms::ConfiguredTableAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ConfiguredTableAssociationAnalysisRules` | configured_table_association_analysis_rules | `list` | optional, computed, provider-chosen |  |  |
| `ConfiguredTableAssociationIdentifier` | configured_table_association_identifier | `string` | computed |  |  |
| `ConfiguredTableIdentifier` | configured_table_identifier | `string` | required, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration. |

Supports update: yes

Discovery: supported (parent resource required)
