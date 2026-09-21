# aws.hostkey

**CloudFormation type:** `AWS::Transfer::HostKey`

Resource type definition for AWS::Transfer::HostKey

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Transfer::HostKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The unique Amazon Resource Name (ARN) for the host key. |
| `DateImported` | date_imported | `string` | computed |  | The date on which the host key was added to the server. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The text description for this host key. |
| `HostKeyBody` | host_key_body | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The private key portion of an SSH key pair. Transfer Family accepts RSA, ECDSA, and ED25519 keys. |
| `HostKeyFingerprint` | host_key_fingerprint | `string` | computed |  | The public key fingerprint, which is a short sequence of bytes used to identify the longer public key. |
| `HostKeyId` | host_key_id | `string` | computed |  | A unique identifier for the host key. |
| `ServerId` | server_id | `string` | required, replaces on change | aws.server.ServerId | The identifier of the server that contains the host key. |
| `Tags` |  | `map` | tags map |  | Key-value pairs that can be used to group and search for host keys. |
| `Type` | type_value | `string` | computed |  | The encryption algorithm that is used for the host key. |

Supports update: yes

Discovery: supported (parent resource required)
