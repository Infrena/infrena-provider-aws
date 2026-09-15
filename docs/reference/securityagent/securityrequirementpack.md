# aws.securityrequirementpack

**CloudFormation type:** `AWS::SecurityAgent::SecurityRequirementPack`

Resource Type definition for AWS::SecurityAgent::SecurityRequirementPack

Region attribute: `region`

**Import ID:** `<region>/PackId` (AWS::SecurityAgent::SecurityRequirementPack)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the pack |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | KMS key for client-side encryption of pack contents |
| `Name` |  | `string` | required |  | Name of the security requirement pack |
| `PackId` | pack_id | `string` | computed |  | Unique identifier of the security requirement pack |
| `SecurityRequirements` | security_requirements | `list` | optional, computed, provider-chosen |  | Security requirements within this pack |
| `Status` |  | `string` | optional, computed, provider-chosen |  | Whether the pack is enabled or disabled |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the security requirement pack |

Supports update: yes

Discovery: supported
