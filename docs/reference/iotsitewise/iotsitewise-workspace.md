# aws.iotsitewise.workspace

**CloudFormation type:** `AWS::IoTSiteWise::Workspace`

Represents an AWS IoT SiteWise workspace that provides logical isolation for tasks and pipelines.

Region attribute: `region`

**Import ID:** `<region>/WorkspaceArn` (AWS::IoTSiteWise::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time the workspace was created. |
| `EncryptionConfiguration` | encryption_configuration | `map` | required, replaces on change |  | The encryption configuration for the workspace. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the AWS KMS key used for KMS_BASED_ENCRYPTION. Required when EncryptionConfiguration.EncryptionType is KMS_BASED_ENCRYPTION. |
| `Status` |  | `string` | computed |  | The current state of the workspace. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The time the workspace was last updated. |
| `WorkspaceArn` | workspace_arn | `string` | computed |  | The ARN of the workspace. |
| `WorkspaceDescription` | workspace_description | `string` | optional, computed, provider-chosen |  | A description of the workspace. |
| `WorkspaceName` | workspace_name | `string` | required, replaces on change |  | The name of the workspace. |

Supports update: yes

Discovery: supported
