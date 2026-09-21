# aws.configurationbundle

**CloudFormation type:** `AWS::BedrockAgentCore::ConfigurationBundle`

Definition of AWS::BedrockAgentCore::ConfigurationBundle Resource Type

Region attribute: `region`

**Import ID:** `<region>/BundleArn` (AWS::BedrockAgentCore::ConfigurationBundle)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BranchName` | branch_name | `string` | optional, computed, provider-chosen, write-only |  | The branch name for version tracking. |
| `BundleArn` | bundle_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the configuration bundle. |
| `BundleId` | bundle_id | `string` | computed |  | The unique identifier of the configuration bundle. |
| `BundleName` | bundle_name | `string` | required, replaces on change |  | The name for the configuration bundle. Names must be unique within your account. |
| `CommitMessage` | commit_message | `string` | optional, computed, provider-chosen, write-only |  | A commit message describing the version of the configuration bundle. |
| `Components` |  | `map` | required |  | A map of component identifiers to their configurations. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the configuration bundle was created. |
| `CreatedBy` | created_by | `map` | optional, computed, provider-chosen, write-only |  | The source that created a configuration bundle version. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description for the configuration bundle. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The ARN of the KMS key used to encrypt component configurations. |
| `LineageMetadata` | lineage_metadata | `map` | computed |  | The version lineage metadata that tracks parent versions and creation source. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the configuration bundle. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the configuration bundle was last updated. |
| `VersionId` | version_id | `string` | computed |  | The version identifier of the configuration bundle. |

Supports update: yes

Discovery: supported
