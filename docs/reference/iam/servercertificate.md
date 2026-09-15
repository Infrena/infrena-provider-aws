# aws.servercertificate

**CloudFormation type:** `AWS::IAM::ServerCertificate`

Resource Type definition for AWS::IAM::ServerCertificate

Global type (no region attribute)

**Import ID:** `global/ServerCertificateName` (AWS::IAM::ServerCertificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the server certificate |
| `CertificateBody` | certificate_body | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `CertificateChain` | certificate_chain | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Path` |  | `string` | optional, computed, provider-chosen |  |  |
| `PrivateKey` | private_key | `string` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  |  |
| `ServerCertificateName` | server_certificate_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
