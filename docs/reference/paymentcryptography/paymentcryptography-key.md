# aws.paymentcryptography.key

**CloudFormation type:** `AWS::PaymentCryptography::Key`

Definition of AWS::PaymentCryptography::Key Resource Type

Region attribute: `region`

**Import ID:** `<region>/KeyIdentifier` (AWS::PaymentCryptography::Key)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeriveKeyUsage` | derive_key_usage | `string` | optional, computed, provider-chosen |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `Exportable` |  | `boolean` | required |  |  |
| `KeyAttributes` | key_attributes | `map` | required |  |  |
| `KeyCheckValueAlgorithm` | key_check_value_algorithm | `string` | optional, computed, provider-chosen |  |  |
| `KeyIdentifier` | key_identifier | `string` | computed |  |  |
| `KeyOrigin` | key_origin | `string` | computed |  | Defines the source of a key |
| `KeyState` | key_state | `string` | computed |  | Defines the state of a key |
| `Policy` |  | `string` | optional, computed, provider-chosen |  | The resource-based policy attached to the key, in JSON format. |
| `ReplicationRegions` | replication_regions | `list` | optional, computed, provider-chosen, write-only |  |  |
| `ReplicationStatus` | replication_status | `map` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
