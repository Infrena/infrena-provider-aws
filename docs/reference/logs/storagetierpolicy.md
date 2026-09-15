# aws.storagetierpolicy

**CloudFormation type:** `AWS::Logs::StorageTierPolicy`

Resource Type definition for AWS::Logs::StorageTierPolicy. Manages the storage tier policy for a CloudWatch Logs account. When created, enables Intelligent-Tiering which automatically moves infrequently accessed log data to lower-cost storage tiers. Deleting this resource reverts the account to standard (single-tier) storage.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::Logs::StorageTierPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The AWS account ID that owns this storage tier policy. |
| `LastUpdatedTime` | last_updated_time | `float` | computed |  | Timestamp (milliseconds after Jan 1, 1970 00:00:00 UTC) when the storage tier policy was last updated. |
| `StorageTier` | storage_tier | `string` | required |  | The storage tier to apply. Only INTELLIGENT_TIERING is accepted for creation. |

Supports update: yes

Discovery: supported
