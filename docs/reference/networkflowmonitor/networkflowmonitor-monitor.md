# aws.networkflowmonitor.monitor

**CloudFormation type:** `AWS::NetworkFlowMonitor::Monitor`

Creates a monitor for specific network flows between local and remote resources to monitor network performance for workloads.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::NetworkFlowMonitor::Monitor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the monitor. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time when the monitor was created. |
| `LocalResources` | local_resources | `list` | required |  | The local resources to monitor. |
| `ModifiedAt` | modified_at | `string` | computed |  | The date and time when the monitor was last modified. |
| `MonitorName` | monitor_name | `string` | required, replaces on change |  | The name of the monitor. |
| `MonitorStatus` | monitor_status | `string` | computed |  | The status of the monitor. |
| `RemoteResources` | remote_resources | `list` | optional, computed, provider-chosen |  | The remote resources to monitor. |
| `ScopeArn` | scope_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) of the scope for the monitor. |
| `Tags` |  | `map` | tags map |  | The tags for the monitor. |

Supports update: yes

Discovery: supported
