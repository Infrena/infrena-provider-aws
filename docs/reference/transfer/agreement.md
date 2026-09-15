# aws.agreement

**CloudFormation type:** `AWS::Transfer::Agreement`

Resource Type definition for AWS::Transfer::Agreement

Region attribute: `region`

**Import ID:** `<region>/AgreementId|ServerId` (AWS::Transfer::Agreement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessRole` | access_role | `string` | required |  | Specifies the access role for the agreement. |
| `AgreementId` | agreement_id | `string` | computed |  | A unique identifier for the agreement. |
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the agreement. |
| `BaseDirectory` | base_directory | `string` | optional, computed, provider-chosen |  | Specifies the base directory for the agreement. |
| `CustomDirectories` | custom_directories | `map` | optional, computed, provider-chosen |  | Specifies a separate directory for each type of file to store for an AS2 message. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A textual description for the agreement. |
| `EnforceMessageSigning` | enforce_message_signing | `string` | optional, computed, provider-chosen |  | Specifies whether to enforce an AS2 message is signed for this agreement. |
| `LocalProfileId` | local_profile_id | `string` | required | aws.transfer.profile.ProfileId | A unique identifier for the local profile. |
| `PartnerProfileId` | partner_profile_id | `string` | required | aws.transfer.profile.ProfileId | A unique identifier for the partner profile. |
| `PreserveFilename` | preserve_filename | `string` | optional, computed, provider-chosen |  | Specifies whether to preserve the filename received for this agreement. |
| `ServerId` | server_id | `string` | required, replaces on change | aws.server.ServerId | A unique identifier for the server. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | Specifies the status of the agreement. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose. |

Supports update: yes

Discovery: supported (parent resource required)
