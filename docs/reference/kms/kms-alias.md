# aws.kms.alias

**CloudFormation type:** `AWS::KMS::Alias`

The ``AWS::KMS::Alias`` resource specifies a display name for a [KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys). You can use an alias to identify a KMS key in the KMS console, in the [DescribeKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html) operation, and in [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations), such as [Decrypt](https://docs.aws.amazon.com/kms/latest/APIReference/API_Decrypt.html) and [GenerateDataKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_GenerateDataKey.html).

Region attribute: `region`

**Import ID:** `<region>/AliasName` (AWS::KMS::Alias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasName` | alias_name | `string` | required, replaces on change |  | Specifies the alias name. This value must begin with ``alias/`` followed by a name, such as ``alias/ExampleAlias``. |
| `TargetKeyId` | target_key_id | `string` | required | aws.kms.key.KeyId | Associates the alias with the specified [](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk). The KMS key must be in the same AWS-account and Region. |

Supports update: yes

Discovery: supported
