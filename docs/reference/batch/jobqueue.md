# aws.jobqueue

**CloudFormation type:** `AWS::Batch::JobQueue`

Resource Type definition for AWS::Batch::JobQueue

Region attribute: `region`

**Import ID:** `<region>/JobQueueArn` (AWS::Batch::JobQueue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputeEnvironmentOrder` | compute_environment_order | `list` | optional, computed, provider-chosen |  |  |
| `JobQueueArn` | job_queue_arn | `string` | computed |  |  |
| `JobQueueName` | job_queue_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `JobQueueType` | job_queue_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `JobStateTimeLimitActions` | job_state_time_limit_actions | `list` | optional, computed, provider-chosen |  |  |
| `Priority` |  | `integer` | required |  |  |
| `SchedulingPolicyArn` | scheduling_policy_arn | `string` | optional, computed, provider-chosen | aws.schedulingpolicy.Arn |  |
| `ServiceEnvironmentOrder` | service_environment_order | `list` | optional, computed, provider-chosen |  |  |
| `State` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
