# aws.deadline.fleet

**CloudFormation type:** `AWS::Deadline::Fleet`

Resource Type definition for AWS::Deadline::Fleet

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Fleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Capabilities` |  | `map` | computed |  |  |
| `Configuration` |  | `string` | required |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `FleetId` | fleet_id | `string` | computed |  |  |
| `HostConfiguration` | host_configuration | `map` | optional, computed, provider-chosen |  |  |
| `MaxWorkerCount` | max_worker_count | `integer` | required |  |  |
| `MinWorkerCount` | min_worker_count | `integer` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `WorkerCount` | worker_count | `integer` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
