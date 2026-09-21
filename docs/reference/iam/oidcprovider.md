# aws.oidcprovider

**CloudFormation type:** `AWS::IAM::OIDCProvider`

Resource Type definition for AWS::IAM::OIDCProvider

Global type (no region attribute)

**Import ID:** `global/Arn` (AWS::IAM::OIDCProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the OIDC provider |
| `ClientIdList` | client_id_list | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `ThumbprintList` | thumbprint_list | `list` | optional, computed, provider-chosen |  |  |
| `Url` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
