# aws.intermediatetable

**CloudFormation type:** `AWS::CleanRooms::IntermediateTable`

Represents an intermediate table that stores cached query results within a collaboration

Region attribute: `region`

**Import ID:** `<region>/IntermediateTableIdentifier|MembershipIdentifier` (AWS::CleanRooms::IntermediateTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalysisRules` | analysis_rules | `list` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IntermediateTableIdentifier` | intermediate_table_identifier | `string` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  |  |
| `MembershipArn` | membership_arn | `string` | computed |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `PopulationAnalysisConfiguration` | population_analysis_configuration | `map` | required, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
