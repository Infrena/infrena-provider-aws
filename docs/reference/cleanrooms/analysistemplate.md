# aws.analysistemplate

**CloudFormation type:** `AWS::CleanRooms::AnalysisTemplate`

Represents a stored analysis within a collaboration

Region attribute: `region`

**Import ID:** `<region>/AnalysisTemplateIdentifier|MembershipIdentifier` (AWS::CleanRooms::AnalysisTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AnalysisParameters` | analysis_parameters | `list` | optional, computed, provider-chosen, replaces on change |  | The member who can query can provide this placeholder for a literal data value in an analysis template |
| `AnalysisTemplateIdentifier` | analysis_template_identifier | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ErrorMessageConfiguration` | error_message_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Format` |  | `string` | required, replaces on change |  |  |
| `MembershipArn` | membership_arn | `string` | computed |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Schema` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Source` |  | `string` | required, replaces on change |  |  |
| `SourceMetadata` | source_metadata | `string` | optional, computed, provider-chosen |  |  |
| `SyntheticDataParameters` | synthetic_data_parameters | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms analysis template. |

Supports update: yes

Discovery: supported (parent resource required)
