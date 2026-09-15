# KMS

AWS Key Management Service (KMS) provides encryption key management and cryptographic operations. KMS keys can encrypt and decrypt data, and aliases provide friendly names for keys. Keys are regional resources — each region has its own set of keys.

## Key

A KMS key is a cryptographic key used for encryption and decryption operations. You can create symmetric keys (the default) for general encryption, asymmetric keys for encryption or signing, or HMAC keys. Keys can have automatic rotation enabled, and you can attach a key policy to control access.

**Type:** `aws.kms.key`

**Import ID:** `<region>/KeyId`

**Settable attributes:**
- `region`: the AWS region for the key (inherited from provider defaults if not specified)
- `description`: a description of the key to distinguish it from others in your account
- `enable_key_rotation`: whether to enable automatic key rotation (boolean, default false)
- `enabled`: whether the key is enabled for cryptographic operations (boolean, default true)
- `key_policy`: the key policy as a JSON document to control who can use and manage the key
- `key_spec`: the type of key to create (default `SYMMETRIC_DEFAULT`; options include `SYMMETRIC_DEFAULT`, `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, etc.)
- `key_usage`: determines the cryptographic operations (default `ENCRYPT_DECRYPT`; alternatives like `SIGN_VERIFY` for asymmetric keys)
- `multi_region`: whether this is a multi-Region primary key (boolean, default false)
- `origin`: the source of the key material (default `AWS_KMS`)
- `tags`: a map of tags

Example:

```yaml
  app_encryption_key:
    type: aws.kms.key
    region: us-east-1
    description: Encryption key for application data
    enable_key_rotation: true
    key_policy: |
      {
        "Version": "2012-10-17",
        "Statement": [
          {
            "Sid": "Enable IAM User Permissions",
            "Effect": "Allow",
            "Principal": {
              "AWS": "arn:aws:iam::123456789012:root"
            },
            "Action": "kms:*",
            "Resource": "*"
          },
          {
            "Sid": "Allow services to use the key",
            "Effect": "Allow",
            "Principal": {
              "Service": "s3.amazonaws.com"
            },
            "Action": ["kms:Decrypt", "kms:GenerateDataKey"],
            "Resource": "*"
          }
        ]
      }
    tags:
      Environment: production
      Purpose: application-encryption
```

## Alias

An alias is a friendly name for a KMS key. Instead of using a key ID (a long UUID), you can reference a key by its alias (for example, `alias/my-app-key`). Aliases make it easier to reference keys in your application code and policies. The alias name must begin with `alias/`.

**Type:** `aws.kms.alias`

**Import ID:** `<region>/AliasName`

**Settable attributes:**
- `alias_name` (required): the alias name, must begin with `alias/` (for example, `alias/my-app-key`)
- `target_key_id` (required): the key ID of the KMS key to associate with the alias
- `region`: the AWS region (inherited from provider defaults if not specified)

Example:

```yaml
  app_key_alias:
    type: aws.kms.alias
    alias_name: alias/app-encryption-key
    target_key_id: ${app_encryption_key.KeyId}
    region: us-east-1
```

## Common patterns and pitfalls

- **Key policies are required for new keys.** If you do not provide a key policy, KMS creates a default policy that grants permissions to the AWS account root. It is recommended to specify a key policy explicitly to follow the principle of least privilege.
- **Key policies are JSON strings.** Key policies are written as JSON text in YAML block strings (`|` or `|-`), just like IAM policies.
- **Aliases must start with `alias/`.** When creating an alias, the `alias_name` must begin with `alias/` followed by a name, such as `alias/my-app-key`.
- **Aliases are region-specific.** An alias is associated with a key in a specific region. If you create keys in multiple regions, you need separate aliases in each region.
- **Key material cannot be changed.** Once a key is created, the `key_spec` (the key type) cannot be changed. Choose the correct key type when creating the key.
- **Automatic key rotation.** Enable `enable_key_rotation` to have KMS automatically rotate the key material once per year. This does not change the key ID; new data is encrypted with rotated material, but existing encrypted data remains readable.
- **Regional resource.** KMS keys are regional. Each region has its own independent set of keys. The key's region is specified in the resource configuration or inherited from the provider's default region.
- **Keys cannot be disabled during deletion.** If you attempt to schedule a key for deletion, AWS prevents immediate deletion to avoid data loss. Keys have a configurable waiting period (7–30 days) before deletion is finalized.

## Reference pages

See the generated reference pages for the complete attribute list and behavior:

- [aws.kms.key](../reference/kms/kms-key.md)
- [aws.kms.alias](../reference/kms/kms-alias.md)
