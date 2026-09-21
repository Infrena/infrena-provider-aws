# aws.tapepool

**CloudFormation type:** `AWS::StorageGateway::TapePool`

Creates a custom tape pool for archiving virtual tapes with optional retention lock.

Region attribute: `region`

**Import ID:** `<region>/PoolARN` (AWS::StorageGateway::TapePool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PoolARN` | pool_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the custom tape pool. |
| `PoolId` | pool_id | `string` | computed |  | The unique identifier of the custom tape pool, extracted from the ARN. |
| `PoolName` | pool_name | `string` | required, replaces on change |  | The name of the custom tape pool. |
| `RetentionLockTimeInDays` | retention_lock_time_in_days | `integer` | optional, computed, provider-chosen, replaces on change |  | Tape retention lock time in days (up to 36,500 days / 100 years). |
| `RetentionLockType` | retention_lock_type | `string` | optional, computed, provider-chosen, replaces on change |  | Tape retention lock type. Governance mode allows authorized removal; compliance mode prevents all removal. |
| `StorageClass` | storage_class | `string` | required, replaces on change |  | The storage class associated with the custom pool (S3 Glacier or S3 Glacier Deep Archive). |
| `Tags` |  | `map` | tags map |  | A list of up to 50 tags for the tape pool. |

Supports update: yes

Discovery: supported
