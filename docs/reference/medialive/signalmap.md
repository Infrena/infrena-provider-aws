# aws.signalmap

**CloudFormation type:** `AWS::MediaLive::SignalMap`

Definition of AWS::MediaLive::SignalMap Resource Type

Region attribute: `region`

**Import ID:** `<region>/Identifier` (AWS::MediaLive::SignalMap)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | A signal map's ARN (Amazon Resource Name) |
| `CloudWatchAlarmTemplateGroupIdentifiers` | cloud_watch_alarm_template_group_identifiers | `list` | optional, computed, provider-chosen, write-only |  |  |
| `CloudWatchAlarmTemplateGroupIds` | cloud_watch_alarm_template_group_ids | `list` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A resource's optional description. |
| `DiscoveryEntryPointArn` | discovery_entry_point_arn | `string` | required |  | A top-level supported AWS resource ARN to discovery a signal map from. |
| `ErrorMessage` | error_message | `string` | computed |  | Error message associated with a failed creation or failed update attempt of a signal map. |
| `EventBridgeRuleTemplateGroupIdentifiers` | event_bridge_rule_template_group_identifiers | `list` | optional, computed, provider-chosen, write-only |  |  |
| `EventBridgeRuleTemplateGroupIds` | event_bridge_rule_template_group_ids | `list` | computed |  |  |
| `FailedMediaResourceMap` | failed_media_resource_map | `map` | computed |  | A map representing an incomplete AWS media workflow as a graph. |
| `ForceRediscovery` | force_rediscovery | `boolean` | optional, computed, provider-chosen, write-only |  | If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided. |
| `Id` |  | `string` | computed |  | A signal map's id. |
| `Identifier` |  | `string` | computed |  |  |
| `LastDiscoveredAt` | last_discovered_at | `string` | computed |  |  |
| `LastSuccessfulMonitorDeployment` | last_successful_monitor_deployment | `map` | computed |  | Represents the latest successful monitor deployment of a signal map. |
| `MediaResourceMap` | media_resource_map | `map` | computed |  | A map representing an AWS media workflow as a graph. |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `MonitorChangesPendingDeployment` | monitor_changes_pending_deployment | `boolean` | computed |  | If true, there are pending monitor changes for this signal map that can be deployed. |
| `MonitorDeployment` | monitor_deployment | `map` | computed |  | Represents the latest monitor deployment of a signal map. |
| `Name` |  | `string` | required |  | A resource's name. Names must be unique within the scope of a resource type in a specific region. |
| `Status` |  | `string` | computed |  | A signal map's current status which is dependent on its lifecycle actions or associated jobs. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Represents the tags associated with a resource. |

Supports update: yes

Discovery: supported
