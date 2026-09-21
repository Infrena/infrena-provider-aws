# aws.codecommit.repository

**CloudFormation type:** `AWS::CodeCommit::Repository`

Resource Type definition for AWS::CodeCommit::Repository

Region attribute: `region`

**Import ID:** `<region>/RepositoryId` (AWS::CodeCommit::Repository)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the repository. |
| `CloneUrlHttp` | clone_url_http | `string` | computed |  | The URL to use for cloning the repository over HTTPS. |
| `CloneUrlSsh` | clone_url_ssh | `string` | computed |  | The URL to use for cloning the repository over SSH. |
| `Code` |  | `map` | optional, computed, provider-chosen, write-only |  | Information about code to be committed to a repository after it is created in an AWS CloudFormation stack. Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The ID of the AWS Key Management Service encryption key used to encrypt and decrypt the repository. |
| `Name` |  | `string` | computed |  | The repository's name. |
| `RepositoryDescription` | repository_description | `string` | optional, computed, provider-chosen |  | A comment or description about the new repository. |
| `RepositoryId` | repository_id | `string` | computed |  | The ID of the repository. |
| `RepositoryName` | repository_name | `string` | required |  | The name of the new repository to be created. |
| `Tags` |  | `map` | tags map |  | One or more tag key-value pairs to use when tagging this repository. |
| `Triggers` |  | `list` | optional, computed, provider-chosen |  | Information about a trigger for a repository. |

Supports update: yes

Discovery: supported
