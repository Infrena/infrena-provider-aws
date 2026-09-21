# aws.channelnamespace

**CloudFormation type:** `AWS::AppSync::ChannelNamespace`

Resource schema for AppSync ChannelNamespace

Region attribute: `region`

**Import ID:** `<region>/ChannelNamespaceArn` (AWS::AppSync::ChannelNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.appsync.api.ApiId | AppSync Api Id that this Channel Namespace belongs to. |
| `ChannelNamespaceArn` | channel_namespace_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the Channel Namespace. |
| `CodeHandlers` | code_handlers | `string` | optional, computed, provider-chosen |  | String of APPSYNC_JS code to be used by the handlers. |
| `CodeS3Location` | code_s3_location | `string` | optional, computed, provider-chosen, write-only |  | The Amazon S3 endpoint where the code is located. |
| `HandlerConfigs` | handler_configs | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  | Namespace indentifier. |
| `PublishAuthModes` | publish_auth_modes | `list` | optional, computed, provider-chosen |  | List of AuthModes supported for Publish operations. |
| `SubscribeAuthModes` | subscribe_auth_modes | `list` | optional, computed, provider-chosen |  | List of AuthModes supported for Subscribe operations. |
| `Tags` |  | `map` | tags map |  | An arbitrary set of tags (key-value pairs) for this AppSync API. |

Supports update: yes

Discovery: supported (parent resource required)
