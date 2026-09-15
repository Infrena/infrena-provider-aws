# aws.securityconfig

**CloudFormation type:** `AWS::OpenSearchServerless::SecurityConfig`

Resource Type definition for AWS::OpenSearchServerless::SecurityConfig

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::OpenSearchServerless::SecurityConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Security config description |
| `IamFederationOptions` | iam_federation_options | `map` | optional, computed, provider-chosen |  | Describe IAM federation options in form of key value map |
| `IamIdentityCenterOptions` | iam_identity_center_options | `map` | optional, computed, provider-chosen |  | Describes IAM Identity Center options for an OpenSearch Serverless security configuration in the form of a key-value map |
| `Id` |  | `string` | computed |  | The identifier of the security config |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The friendly name of the security config |
| `SamlOptions` | saml_options | `map` | optional, computed, provider-chosen |  | Describes saml options in form of key value map |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | Config type for security config |

Supports update: yes

Discovery: supported (parent resource required)
