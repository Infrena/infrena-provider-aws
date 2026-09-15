# aws.replicakey

**CloudFormation type:** `AWS::KMS::ReplicaKey`

The AWS::KMS::ReplicaKey resource specifies a multi-region replica AWS KMS key in AWS Key Management Service (AWS KMS).

Region attribute: `region`

**Import ID:** `<region>/KeyId` (AWS::KMS::ReplicaKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use. |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations. |
| `KeyId` | key_id | `string` | computed |  |  |
| `KeyPolicy` | key_policy | `string` | required |  | The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules. |
| `PendingWindowInDays` | pending_window_in_days | `integer` | optional, computed, provider-chosen, write-only |  | Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days. |
| `PrimaryKeyArn` | primary_key_arn | `string` | required, replaces on change | aws.kms.key.Arn | Identifies the primary AWS KMS key to create a replica of. Specify the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
