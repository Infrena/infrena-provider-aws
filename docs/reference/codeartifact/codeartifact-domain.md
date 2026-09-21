# aws.codeartifact.domain

**CloudFormation type:** `AWS::CodeArtifact::Domain`

The resource schema to create a CodeArtifact domain.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeArtifact::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the domain. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the domain. |
| `EncryptionKey` | encryption_key | `string` | computed |  | The ARN of an AWS Key Management Service (AWS KMS) key associated with a domain. |
| `Name` |  | `string` | computed |  | The name of the domain. This field is used for GetAtt |
| `Owner` |  | `string` | computed |  | The 12-digit account ID of the AWS account that owns the domain. This field is used for GetAtt |
| `PermissionsPolicyDocument` | permissions_policy_document | `map` | optional, computed, provider-chosen |  | The access control resource policy on the provided domain. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
