# aws.datatransformationprofile

**CloudFormation type:** `AWS::HealthLake::DataTransformationProfile`

Creates a Data Transformation Profile in AWS HealthLake that converts healthcare data from a source format (such as C-CDA or CSV) into FHIR R4. A profile is immutable once created; to change its template content, replace the resource. Only its tags can be updated in place.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::HealthLake::DataTransformationProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the data transformation profile. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier (key ID or ARN) of a customer-managed KMS key used to encrypt the profile's template content at rest. If omitted, an AWS owned key is used. |
| `ProfileDescription` | profile_description | `string` | optional, computed, provider-chosen, replaces on change |  | A human-readable description of the profile's purpose. |
| `ProfileId` | profile_id | `string` | computed |  | The unique, server-generated identifier of the profile (32-character lowercase hexadecimal). |
| `ProfileName` | profile_name | `string` | required, replaces on change |  | The human-readable name of the profile. |
| `Source` |  | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The source from which to create the profile's initial template content. Exactly one of the members must be specified. Use StarterProfile (C-CDA only), ProfileMapping (C-CDA or CSV), or ExistingVersionedProfileId to clone an existing profile. Each produces a published profile. |
| `SourceFormat` | source_format | `string` | required, replaces on change |  | The source format that this profile converts from. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this profile. |
| `TargetFormat` | target_format | `string` | computed |  | The target format that this profile converts to. Always FHIR_R4. |

Supports update: yes

Discovery: supported (parent resource required)
