# aws.keysigningkey

**CloudFormation type:** `AWS::Route53::KeySigningKey`

Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.

Global type (no region attribute)

**Import ID:** `global/HostedZoneId|Name` (AWS::Route53::KeySigningKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HostedZoneId` | hosted_zone_id | `string` | required, replaces on change | aws.hostedzone.Id | The unique string (ID) used to identify a hosted zone. |
| `KeyManagementServiceArn` | key_management_service_arn | `string` | required, replaces on change |  | The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone. |
| `Name` |  | `string` | required, replaces on change |  | An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone. |
| `Status` |  | `string` | required |  | A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE. |

Supports update: yes

Discovery: supported
