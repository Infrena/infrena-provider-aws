# aws.playbackkeypair

**CloudFormation type:** `AWS::IVS::PlaybackKeyPair`

Resource Type definition for AWS::IVS::PlaybackKeyPair

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::PlaybackKeyPair)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Key-pair identifier. |
| `Fingerprint` |  | `string` | computed |  | Key-pair identifier. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique. |
| `PublicKeyMaterial` | public_key_material | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The public portion of a customer-generated key pair. This field is required to create the AWS::IVS::PlaybackKeyPair resource. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the asset model. |

Supports update: yes

Discovery: supported
