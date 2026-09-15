# aws.transfer.profile

**CloudFormation type:** `AWS::Transfer::Profile`

Resource Type definition for AWS::Transfer::Profile

Region attribute: `region`

**Import ID:** `<region>/ProfileId` (AWS::Transfer::Profile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the profile. |
| `As2Id` | as2_id | `string` | required |  | AS2 identifier agreed with a trading partner. |
| `CertificateIds` | certificate_ids | `list` | optional, computed, provider-chosen | aws.transfer.certificate.CertificateId | List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages. |
| `ProfileId` | profile_id | `string` | computed |  | A unique identifier for the profile |
| `ProfileType` | profile_type | `string` | required, replaces on change |  | Enum specifying whether the profile is local or associated with a trading partner. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
