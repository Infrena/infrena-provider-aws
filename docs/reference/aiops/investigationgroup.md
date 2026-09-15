# aws.investigationgroup

**CloudFormation type:** `AWS::AIOps::InvestigationGroup`

Definition of AWS::AIOps::InvestigationGroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AIOps::InvestigationGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Investigation Group's ARN. |
| `ChatbotNotificationChannels` | chatbot_notification_channels | `list` | optional, computed, provider-chosen |  | An array of key-value pairs of notification channels to apply to this resource. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp value. |
| `CreatedBy` | created_by | `string` | computed |  | User friendly name for resources. |
| `CrossAccountConfigurations` | cross_account_configurations | `list` | optional, computed, provider-chosen |  | An array of cross account configurations. |
| `EncryptionConfig` | encryption_config | `map` | optional, computed, provider-chosen |  |  |
| `InvestigationGroupPolicy` | investigation_group_policy | `string` | optional, computed, provider-chosen |  | Investigation Group policy |
| `IsCloudTrailEventHistoryEnabled` | is_cloud_trail_event_history_enabled | `boolean` | optional, computed, provider-chosen |  | Flag to enable cloud trail history |
| `LastModifiedAt` | last_modified_at | `string` | computed |  | User friendly name for resources. |
| `LastModifiedBy` | last_modified_by | `string` | computed |  | User friendly name for resources. |
| `Name` |  | `string` | required, replaces on change |  | User friendly name for resources. |
| `RetentionInDays` | retention_in_days | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of days to retain the investigation group |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Investigation Role's ARN. |
| `TagKeyBoundaries` | tag_key_boundaries | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
