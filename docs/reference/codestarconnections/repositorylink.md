# aws.repositorylink

**CloudFormation type:** `AWS::CodeStarConnections::RepositoryLink`

Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.

Region attribute: `region`

**Import ID:** `<region>/RepositoryLinkArn` (AWS::CodeStarConnections::RepositoryLink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionArn` | connection_arn | `string` | required | aws.codestarconnections.connection.ConnectionArn | The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services. |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen |  | The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used. |
| `OwnerId` | owner_id | `string` | required, replaces on change |  | the ID of the entity that owns the repository. |
| `ProviderType` | provider_type | `string` | computed |  | The name of the external provider where your third-party code repository is configured. |
| `RepositoryLinkArn` | repository_link_arn | `string` | computed |  | A unique Amazon Resource Name (ARN) to designate the repository link. |
| `RepositoryLinkId` | repository_link_id | `string` | computed |  | A UUID that uniquely identifies the RepositoryLink. |
| `RepositoryName` | repository_name | `string` | required, replaces on change |  | The repository for which the link is being created. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Specifies the tags applied to a RepositoryLink. |

Supports update: yes

Discovery: supported
