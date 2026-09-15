# aws.eventdatastore

**CloudFormation type:** `AWS::CloudTrail::EventDataStore`

A storage lake of event data against which you can run complex SQL-based queries. An event data store can include events that you have logged on your account from the last 7 to 2557 or 3653 days (about seven or ten years) depending on the selected BillingMode.

Region attribute: `region`

**Import ID:** `<region>/EventDataStoreArn` (AWS::CloudTrail::EventDataStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdvancedEventSelectors` | advanced_event_selectors | `list` | optional, computed, provider-chosen |  | The advanced event selectors that were used to select events for the data store. |
| `BillingMode` | billing_mode | `string` | optional, computed, provider-chosen |  | The mode that the event data store will use to charge for event storage. |
| `ContextKeySelectors` | context_key_selectors | `list` | optional, computed, provider-chosen |  | An array that enriches event records in an existing event data store by including additional information specified in individual ContexKeySelector entries. If you add ContextKeySelectors, you must set MaxEventSize to Large. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The timestamp of the event data store's creation. |
| `EventDataStoreArn` | event_data_store_arn | `string` | computed |  | The ARN of the event data store. |
| `FederationEnabled` | federation_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether federation is enabled on an event data store. |
| `FederationRoleArn` | federation_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The ARN of the role used for event data store federation. |
| `IngestionEnabled` | ingestion_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the event data store is ingesting events. |
| `InsightSelectors` | insight_selectors | `list` | optional, computed, provider-chosen |  | Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing event data store. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store. |
| `InsightsDestination` | insights_destination | `string` | optional, computed, provider-chosen |  | Specifies the ARN of the event data store that will collect Insights events. Both InsightSelectors and InsightsDestination need to have a value in order to enable Insights events on an event data store |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier. |
| `MaxEventSize` | max_event_size | `string` | optional, computed, provider-chosen |  | Specifies the maximum size allowed for the event. Valid values are Standard and Large. If you add ContextKeySelectors, this value must be set to Large. |
| `MultiRegionEnabled` | multi_region_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the event data store includes events from all regions, or only from the region in which it was created. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the event data store. |
| `OrganizationEnabled` | organization_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates that an event data store is collecting logged events for an organization. |
| `RetentionPeriod` | retention_period | `integer` | optional, computed, provider-chosen |  | The retention period, in days. |
| `Status` |  | `string` | computed |  | The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TerminationProtectionEnabled` | termination_protection_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the event data store is protected from termination. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  | The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp. |

Supports update: yes

Discovery: supported
