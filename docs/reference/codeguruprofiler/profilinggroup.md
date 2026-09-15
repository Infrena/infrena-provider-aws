# aws.profilinggroup

**CloudFormation type:** `AWS::CodeGuruProfiler::ProfilingGroup`

This resource schema represents the Profiling Group resource in the Amazon CodeGuru Profiler service.

Region attribute: `region`

**Import ID:** `<region>/ProfilingGroupName` (AWS::CodeGuruProfiler::ProfilingGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentPermissions` | agent_permissions | `map` | optional, computed, provider-chosen |  | The agent permissions attached to this profiling group. |
| `AnomalyDetectionNotificationConfiguration` | anomaly_detection_notification_configuration | `list` | optional, computed, provider-chosen |  | Configuration for Notification Channels for Anomaly Detection feature in CodeGuru Profiler which enables customers to detect anomalies in the application profile for those methods that represent the highest proportion of CPU time or latency |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified profiling group. |
| `ComputePlatform` | compute_platform | `string` | optional, computed, provider-chosen, replaces on change |  | The compute platform of the profiling group. |
| `ProfilingGroupName` | profiling_group_name | `string` | required, replaces on change |  | The name of the profiling group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with a profiling group. |

Supports update: yes

Discovery: supported
