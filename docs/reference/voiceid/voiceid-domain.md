# aws.voiceid.domain

**CloudFormation type:** `AWS::VoiceID::Domain`

The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.

Region attribute: `region`

**Import ID:** `<region>/DomainId` (AWS::VoiceID::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DomainId` | domain_id | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `ServerSideEncryptionConfiguration` | server_side_encryption_configuration | `map` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
