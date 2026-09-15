# aws.quotashare

**CloudFormation type:** `AWS::Batch::QuotaShare`

Creates an AWS Batch quota share. Each quota share operates as a virtual queue with a configured compute capacity, resource sharing strategy, and borrow limits.

Region attribute: `region`

**Import ID:** `<region>/QuotaShareArn` (AWS::Batch::QuotaShare)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityLimits` | capacity_limits | `list` | required |  | A list that specifies the quantity and type of compute capacity allocated to the quota share. |
| `JobQueue` | job_queue | `string` | required, replaces on change |  | The AWS Batch job queue associated with the quota share. This can be the job queue name or ARN. A job queue must be in the `VALID` state before you can associate it with a quota share. |
| `PreemptionConfiguration` | preemption_configuration | `map` | required |  | Specifies the preemption behavior for jobs in a quota share. |
| `QuotaShareArn` | quota_share_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the quota share. |
| `QuotaShareName` | quota_share_name | `string` | required, replaces on change |  | The name of the quota share. It can be up to 128 characters long. It can contain uppercase and lowercase letters, numbers, hyphens (-), and underscores (_). |
| `ResourceSharingConfiguration` | resource_sharing_configuration | `map` | required |  | Specifies whether a quota share reserves, lends, or both lends and borrows idle compute capacity. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the quota share. If the quota share is `ENABLED`, it is able to accept jobs. If the quota share is `DISABLED`, new jobs won't be accepted but jobs already submitted can finish. The default state is `ENABLED`. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags that you apply to the quota share to help you categorize and organize your resources. Each tag consists of a key and an optional value. |

Supports update: yes

Discovery: supported
