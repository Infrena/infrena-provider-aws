# aws.collaboration

**CloudFormation type:** `AWS::CleanRooms::Collaboration`

Represents a collaboration between AWS accounts that allows for secure data collaboration

Region attribute: `region`

**Import ID:** `<region>/CollaborationIdentifier` (AWS::CleanRooms::Collaboration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedResultRegions` | allowed_result_regions | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `AnalyticsEngine` | analytics_engine | `string` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AutoApprovedChangeTypes` | auto_approved_change_types | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `CreatorDisplayName` | creator_display_name | `string` | required, replaces on change |  |  |
| `CreatorMLMemberAbilities` | creator_ml_member_abilities | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatorMemberAbilities` | creator_member_abilities | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatorPaymentConfiguration` | creator_payment_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `DataEncryptionMetadata` | data_encryption_metadata | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Description` |  | `string` | required |  |  |
| `IsMetricsEnabled` | is_metrics_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `JobLogStatus` | job_log_status | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Members` |  | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required |  |  |
| `QueryLogStatus` | query_log_status | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration. |

Supports update: yes

Discovery: supported
