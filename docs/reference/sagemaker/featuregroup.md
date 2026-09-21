# aws.featuregroup

**CloudFormation type:** `AWS::SageMaker::FeatureGroup`

Resource Type definition for AWS::SageMaker::FeatureGroup

Region attribute: `region`

**Import ID:** `<region>/FeatureGroupName` (AWS::SageMaker::FeatureGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | A timestamp of FeatureGroup creation time. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description about the FeatureGroup. |
| `EventTimeFeatureName` | event_time_feature_name | `string` | required, replaces on change |  | The Event Time Feature Name. |
| `FeatureDefinitions` | feature_definitions | `list` | required |  | An Array of Feature Definition |
| `FeatureGroupName` | feature_group_name | `string` | required, replaces on change |  | The Name of the FeatureGroup. |
| `FeatureGroupStatus` | feature_group_status | `string` | computed |  | The status of the feature group. |
| `OfflineStoreConfig` | offline_store_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `OnlineStoreConfig` | online_store_config | `map` | optional, computed, provider-chosen |  |  |
| `RecordIdentifierFeatureName` | record_identifier_feature_name | `string` | required, replaces on change |  | The Record Identifier Feature Name. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | Role Arn |
| `Tags` |  | `map` | replaces on change, tags map |  | An array of key-value pair to apply to this resource. |
| `ThroughputConfig` | throughput_config | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
