# aws.internetmonitor.monitor

**CloudFormation type:** `AWS::InternetMonitor::Monitor`

Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application

Region attribute: `region`

**Import ID:** `<region>/MonitorName` (AWS::InternetMonitor::Monitor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ) |
| `HealthEventsConfig` | health_events_config | `map` | optional, computed, provider-chosen |  |  |
| `IncludeLinkedAccounts` | include_linked_accounts | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `InternetMeasurementsLogDelivery` | internet_measurements_log_delivery | `map` | optional, computed, provider-chosen |  |  |
| `LinkedAccountId` | linked_account_id | `string` | optional, computed, provider-chosen, write-only |  |  |
| `MaxCityNetworksToMonitor` | max_city_networks_to_monitor | `integer` | optional, computed, provider-chosen |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  | The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ) |
| `MonitorArn` | monitor_arn | `string` | computed |  |  |
| `MonitorName` | monitor_name | `string` | required, replaces on change |  |  |
| `ProcessingStatus` | processing_status | `string` | computed |  |  |
| `ProcessingStatusInfo` | processing_status_info | `string` | computed |  |  |
| `Resources` |  | `list` | optional, computed, provider-chosen |  |  |
| `ResourcesToAdd` | resources_to_add | `list` | optional, computed, provider-chosen, write-only |  |  |
| `ResourcesToRemove` | resources_to_remove | `list` | optional, computed, provider-chosen, write-only |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TrafficPercentageToMonitor` | traffic_percentage_to_monitor | `integer` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
