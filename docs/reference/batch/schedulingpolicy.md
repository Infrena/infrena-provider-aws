# aws.schedulingpolicy

**CloudFormation type:** `AWS::Batch::SchedulingPolicy`

Resource Type definition for AWS::Batch::SchedulingPolicy

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Batch::SchedulingPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the Scheduling Policy. |
| `FairsharePolicy` | fairshare_policy | `map` | optional, computed, provider-chosen |  | Fair Share Policy for the Job Queue. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of Scheduling Policy. |
| `QuotaSharePolicy` | quota_share_policy | `map` | optional, computed, provider-chosen |  | Quota Share Policy for the Job Queue. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
