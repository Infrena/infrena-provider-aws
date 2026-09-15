# aws.kms.key

**CloudFormation type:** `AWS::KMS::Key`

The ``AWS::KMS::Key`` resource specifies an [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys) in KMSlong. You can use this resource to create symmetric encryption KMS keys, asymmetric KMS keys for encryption or signing, and symmetric HMAC KMS keys. You can use ``AWS::KMS::Key`` to create [multi-Region primary keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html#mrk-primary-key) of all supported types. To replicate a multi-Region key, use the ``AWS::KMS::ReplicaKey`` resource.

Region attribute: `region`

**Import ID:** `<region>/KeyId` (AWS::KMS::Key)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BypassPolicyLockoutSafetyCheck` | bypass_policy_lockout_safety_check | `boolean` | optional, computed, provider-chosen, write-only |  | Skips ("bypasses") the key policy lockout safety check. The default value is false. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the KMS key. Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use. |
| `EnableKeyRotation` | enable_key_rotation | `boolean` | optional, computed, provider-chosen |  | Enables automatic rotation of the key material for the specified KMS key. By default, automatic key rotation is not enabled. |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations. |
| `KeyId` | key_id | `string` | computed |  |  |
| `KeyPolicy` | key_policy | `string` | optional, computed, provider-chosen |  | The key policy to attach to the KMS key. |
| `KeySpec` | key_spec | `string` | optional, computed, provider-chosen |  | Specifies the type of KMS key to create. The default value, ``SYMMETRIC_DEFAULT``, creates a KMS key with a 256-bit symmetric key for encryption and decryption. In China Regions, ``SYMMETRIC_DEFAULT`` creates a 128-bit symmetric key that uses SM4 encryption. You can't change the ``KeySpec`` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *Developer Guide*. |
| `KeyUsage` | key_usage | `string` | optional, computed, provider-chosen |  | Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ``ENCRYPT_DECRYPT``. This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the ``KeyUsage`` value after the KMS key is created. |
| `MultiRegion` | multi_region | `boolean` | optional, computed, provider-chosen |  | Creates a multi-Region primary key that you can replicate in other AWS-Regions. You can't change the ``MultiRegion`` value after the KMS key is created. |
| `Origin` |  | `string` | optional, computed, provider-chosen |  | The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is ``AWS_KMS``, which means that KMS creates the key material. |
| `PendingWindowInDays` | pending_window_in_days | `integer` | optional, computed, provider-chosen, write-only |  | Specifies the number of days in the waiting period before KMS deletes a KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days. |
| `RotationPeriodInDays` | rotation_period_in_days | `integer` | optional, computed, provider-chosen, write-only |  | Specifies a custom period of time between each rotation date. If no value is specified, the default value is 365 days. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Assigns one or more tags to the replica key. |

Supports update: yes

Discovery: supported
