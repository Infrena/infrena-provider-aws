# aws.repositoryassociation

**CloudFormation type:** `AWS::CodeGuruReviewer::RepositoryAssociation`

This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.

Region attribute: `region`

**Import ID:** `<region>/AssociationArn` (AWS::CodeGuruReviewer::RepositoryAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationArn` | association_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the repository association. |
| `BucketName` | bucket_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`. |
| `ConnectionArn` | connection_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection. |
| `Name` |  | `string` | required, replaces on change |  | Name of the repository to be associated. |
| `Owner` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | The tags associated with a repository association. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of repository to be associated. |

Supports update: no

Discovery: supported
