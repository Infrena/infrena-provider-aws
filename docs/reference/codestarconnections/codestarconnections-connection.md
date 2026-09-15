# aws.codestarconnections.connection

**CloudFormation type:** `AWS::CodeStarConnections::Connection`

Schema for AWS::CodeStarConnections::Connection resource which can be used to connect external source providers with AWS CodePipeline

Region attribute: `region`

**Import ID:** `<region>/ConnectionArn` (AWS::CodeStarConnections::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionArn` | connection_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services. |
| `ConnectionName` | connection_name | `string` | required, replaces on change |  | The name of the connection. Connection names must be unique in an AWS user account. |
| `ConnectionStatus` | connection_status | `string` | computed |  | The current status of the connection. |
| `HostArn` | host_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn. |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository. |
| `ProviderType` | provider_type | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Specifies the tags applied to a connection. |

Supports update: yes

Discovery: supported
