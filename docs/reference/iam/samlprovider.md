# aws.samlprovider

**CloudFormation type:** `AWS::IAM::SAMLProvider`

Resource Type definition for AWS::IAM::SAMLProvider

Global type (no region attribute)

**Import ID:** `global/Arn` (AWS::IAM::SAMLProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddPrivateKey` | add_private_key | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  | The private key from your external identity provider |
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the SAML provider |
| `AssertionEncryptionMode` | assertion_encryption_mode | `string` | optional, computed, provider-chosen |  | The encryption setting for the SAML provider |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PrivateKeyList` | private_key_list | `list` | optional, computed, provider-chosen |  |  |
| `RemovePrivateKey` | remove_private_key | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Key ID of the private key to remove |
| `SamlMetadataDocument` | saml_metadata_document | `string` | optional, computed, provider-chosen |  |  |
| `SamlProviderUUID` | saml_provider_uuid | `string` | computed |  | The unique identifier assigned to the SAML provider |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
