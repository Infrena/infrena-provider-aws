# aws.spendinglimit

**CloudFormation type:** `AWS::Braket::SpendingLimit`

Creates a spending limit for a specified quantum device. Spending limits help you control costs by setting maximum amounts that can be spent on quantum computing tasks within a specified time period.

Region attribute: `region`

**Import ID:** `<region>/SpendingLimitArn` (AWS::Braket::SpendingLimit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The date and time when the spending limit was created, in ISO 8601 format. |
| `DeviceArn` | device_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the quantum device to apply the spending limit to. |
| `QueuedSpend` | queued_spend | `string` | computed |  | The amount currently queued for spending on the device, in USD. |
| `SpendingLimit` | spending_limit | `string` | required |  | The maximum amount that can be spent on the specified device, in USD. |
| `SpendingLimitArn` | spending_limit_arn | `string` | computed |  | The Amazon Resource Name (ARN) that uniquely identifies the spending limit. |
| `Tags` |  | `map` | tags map |  | The tags to apply to the spending limit. |
| `TimePeriod` | time_period | `map` | optional, computed, provider-chosen |  | Defines a time range for spending limits, specifying when the limit is active. |
| `TotalSpend` | total_spend | `string` | computed |  | The total amount spent on the device so far during the current time period, in USD. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time when the spending limit was last modified, in ISO 8601 format. |

Supports update: yes

Discovery: supported
