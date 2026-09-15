# aws.syncconfiguration

**CloudFormation type:** `AWS::CodeStarConnections::SyncConfiguration`

Schema for AWS::CodeStarConnections::SyncConfiguration resource which is used to enables an AWS resource to be synchronized from a source-provider.

Region attribute: `region`

**Import ID:** `<region>/ResourceName|SyncType` (AWS::CodeStarConnections::SyncConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Branch` |  | `string` | required |  | The name of the branch of the repository from which resources are to be synchronized, |
| `ConfigFile` | config_file | `string` | required |  | The source provider repository path of the sync configuration file of the respective SyncType. |
| `OwnerId` | owner_id | `string` | computed |  | the ID of the entity that owns the repository. |
| `ProviderType` | provider_type | `string` | computed |  | The name of the external provider where your third-party code repository is configured. |
| `PublishDeploymentStatus` | publish_deployment_status | `string` | optional, computed, provider-chosen |  | Whether to enable or disable publishing of deployment status to source providers. |
| `RepositoryLinkId` | repository_link_id | `string` | required | aws.repositorylink.RepositoryLinkId | A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with. |
| `RepositoryName` | repository_name | `string` | computed |  | The name of the repository that is being synced to. |
| `ResourceName` | resource_name | `string` | required, replaces on change |  | The name of the resource that is being synchronized to the repository. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository. |
| `SyncType` | sync_type | `string` | required, replaces on change |  | The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC. |
| `TriggerResourceUpdateOn` | trigger_resource_update_on | `string` | optional, computed, provider-chosen |  | When to trigger Git sync to begin the stack update. |

Supports update: yes

Discovery: supported
