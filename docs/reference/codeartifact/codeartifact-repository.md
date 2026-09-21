# aws.codeartifact.repository

**CloudFormation type:** `AWS::CodeArtifact::Repository`

The resource schema to create a CodeArtifact repository.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeArtifact::Repository)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the repository. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A text description of the repository. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the domain that contains the repository. |
| `DomainOwner` | domain_owner | `string` | computed |  | The 12-digit account ID of the AWS account that owns the domain. |
| `ExternalConnections` | external_connections | `list` | optional, computed, provider-chosen |  | A list of external connections associated with the repository. |
| `Name` |  | `string` | computed |  | The name of the repository. This is used for GetAtt |
| `PermissionsPolicyDocument` | permissions_policy_document | `map` | optional, computed, provider-chosen |  | The access control resource policy on the provided repository. |
| `RepositoryName` | repository_name | `string` | required, replaces on change |  | The name of the repository. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Upstreams` |  | `list` | optional, computed, provider-chosen |  | A list of upstream repositories associated with the repository. |

Supports update: yes

Discovery: supported
