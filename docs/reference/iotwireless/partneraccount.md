# aws.partneraccount

**CloudFormation type:** `AWS::IoTWireless::PartnerAccount`

Create and manage partner account

Region attribute: `region`

**Import ID:** `<region>/PartnerAccountId` (AWS::IoTWireless::PartnerAccount)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountLinked` | account_linked | `boolean` | optional, computed, provider-chosen |  | Whether the partner account is linked to the AWS account. |
| `Arn` |  | `string` | computed |  | PartnerAccount arn. Returned after successful create. |
| `Fingerprint` |  | `string` | computed |  | The fingerprint of the Sidewalk application server private key. |
| `PartnerAccountId` | partner_account_id | `string` | optional, computed, provider-chosen, replaces on change | aws.partneraccount.PartnerAccountId | The partner account ID to disassociate from the AWS account |
| `PartnerType` | partner_type | `string` | optional, computed, provider-chosen |  | The partner type |
| `Sidewalk` |  | `map` | optional, computed, provider-chosen, write-only |  | The Sidewalk account credentials. |
| `SidewalkResponse` | sidewalk_response | `map` | optional, computed, provider-chosen |  | The Sidewalk account credentials. |
| `SidewalkUpdate` | sidewalk_update | `map` | optional, computed, provider-chosen, write-only |  | The Sidewalk account credentials. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the destination. |

Supports update: yes

Discovery: supported
