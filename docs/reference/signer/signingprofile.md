# aws.signingprofile

**CloudFormation type:** `AWS::Signer::SigningProfile`

A signing profile is a signing template that can be used to carry out a pre-defined signing job.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Signer::SigningProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified signing profile. |
| `PlatformId` | platform_id | `string` | required, replaces on change |  | The ID of the target signing platform. |
| `ProfileName` | profile_name | `string` | computed |  | A name for the signing profile. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. |
| `ProfileVersion` | profile_version | `string` | computed |  | A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name. |
| `ProfileVersionArn` | profile_version_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified signing profile version. |
| `SignatureValidityPeriod` | signature_validity_period | `map` | optional, computed, provider-chosen, replaces on change |  | Signature validity period of the profile. |
| `Tags` |  | `map` | tags map |  | A list of tags associated with the signing profile. |

Supports update: yes

Discovery: supported
