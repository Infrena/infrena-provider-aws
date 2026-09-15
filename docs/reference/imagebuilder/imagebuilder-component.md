# aws.imagebuilder.component

**CloudFormation type:** `AWS::ImageBuilder::Component`

Resource Type definition for AWS::ImageBuilder::Component

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::Component)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the component. |
| `ChangeDescription` | change_description | `string` | optional, computed, provider-chosen, replaces on change |  | The change description of the component. |
| `Data` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The data of the component. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the component. |
| `Encrypted` |  | `boolean` | computed |  | The encryption status of the component. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The KMS key identifier used to encrypt the component. |
| `LatestVersion` | latest_version | `map` | computed |  | The latest version references of the component. |
| `Name` |  | `string` | required, replaces on change |  | The name of the component. |
| `Platform` |  | `string` | required, replaces on change |  | The platform of the component. |
| `SupportedOsVersions` | supported_os_versions | `list` | optional, computed, provider-chosen, replaces on change |  | The operating system (OS) version supported by the component. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the component. |
| `Type` | type_value | `string` | computed |  | The type of the component denotes whether the component is used to build the image or only to test it. |
| `Uri` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The uri of the component. |
| `Version` |  | `string` | required, replaces on change |  | The version of the component. |

Supports update: yes

Discovery: supported
