# aws.tracker

**CloudFormation type:** `AWS::Location::Tracker`

Definition of AWS::Location::Tracker Resource Type

Region attribute: `region`

**Import ID:** `<region>/TrackerName` (AWS::Location::Tracker)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EventBridgeEnabled` | event_bridge_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `KmsKeyEnableGeospatialQueries` | kms_key_enable_geospatial_queries | `boolean` | optional, computed, provider-chosen |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PositionFiltering` | position_filtering | `string` | optional, computed, provider-chosen |  |  |
| `PricingPlan` | pricing_plan | `string` | optional, computed, provider-chosen |  |  |
| `PricingPlanDataSource` | pricing_plan_data_source | `string` | optional, computed, provider-chosen |  | This shape is deprecated since 2022-02-01: Deprecated. No longer allowed. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TrackerArn` | tracker_arn | `string` | computed |  |  |
| `TrackerName` | tracker_name | `string` | required, replaces on change |  |  |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
