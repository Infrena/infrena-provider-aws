# aws.ivs.publickey

**CloudFormation type:** `AWS::IVS::PublicKey`

Resource Type definition for AWS::IVS::PublicKey

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::PublicKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Key-pair identifier. |
| `Fingerprint` |  | `string` | computed |  | Key-pair identifier. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the public key to be imported. The value does not need to be unique. |
| `PublicKeyMaterial` | public_key_material | `string` | optional, computed, provider-chosen, replaces on change |  | The public portion of a customer-generated key pair. This field is required to create the AWS::IVS::PublicKey resource. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |

Supports update: yes

Discovery: supported
