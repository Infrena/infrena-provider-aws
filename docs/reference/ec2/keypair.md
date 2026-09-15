# aws.keypair

**CloudFormation type:** `AWS::EC2::KeyPair`

Specifies a key pair for use with an EC2long instance as follows:

Region attribute: `region`

**Import ID:** `<region>/KeyName` (AWS::EC2::KeyPair)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `KeyFingerprint` | key_fingerprint | `string` | computed |  |  |
| `KeyFormat` | key_format | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The format of the key pair. |
| `KeyName` | key_name | `string` | required, replaces on change |  | A unique name for the key pair. |
| `KeyPairId` | key_pair_id | `string` | computed |  |  |
| `KeyType` | key_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of key pair. Note that ED25519 keys are not supported for Windows instances. |
| `PublicKeyMaterial` | public_key_material | `string` | optional, computed, provider-chosen, replaces on change |  | The public key material. The ``PublicKeyMaterial`` property is used to import a key pair. If this property is not specified, then a new key pair will be created. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | The tags to apply to the key pair. |

Supports update: no

Discovery: supported
