# aws.appstream.stack

**CloudFormation type:** `AWS::AppStream::Stack`

Resource Type definition for AWS::AppStream::Stack

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::AppStream::Stack)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessEndpoints` | access_endpoints | `list` | optional, computed, provider-chosen |  | The list of virtual private cloud (VPC) interface endpoint objects. Users of the stack can connect to AppStream 2.0 only through the specified endpoints. |
| `AgentAccessConfig` | agent_access_config | `map` | optional, computed, provider-chosen |  | The configuration for agent access on a stack. Agent access enables AI agents to interact with desktop applications during streaming sessions. |
| `ApplicationSettings` | application_settings | `map` | optional, computed, provider-chosen |  | The persistent application settings for users of a stack. |
| `AttributesToDelete` | attributes_to_delete | `list` | optional, computed, provider-chosen, write-only |  | The stack attributes to delete. |
| `ContentRedirection` | content_redirection | `map` | optional, computed, provider-chosen |  | The content redirection settings for the stack. |
| `DeleteStorageConnectors` | delete_storage_connectors | `boolean` | optional, computed, provider-chosen, write-only |  | This parameter has been deprecated. Deletes the storage connectors currently enabled for the stack. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description to display. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The stack name to display. |
| `EmbedHostDomains` | embed_host_domains | `list` | optional, computed, provider-chosen |  | The domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions. |
| `FeedbackURL` | feedback_url | `string` | optional, computed, provider-chosen |  | The URL that users are redirected to after they click the Send Feedback link. If no URL is specified, no Send Feedback link is displayed. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the stack. |
| `RedirectURL` | redirect_url | `string` | optional, computed, provider-chosen |  | The URL that users are redirected to after their streaming session ends. |
| `StorageConnectors` | storage_connectors | `list` | optional, computed, provider-chosen |  | The storage connectors to enable. |
| `StreamingExperienceSettings` | streaming_experience_settings | `map` | optional, computed, provider-chosen |  | The streaming protocol that you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs. |
| `UserSettings` | user_settings | `list` | optional, computed, provider-chosen |  | The actions that are enabled or disabled for users during their streaming sessions. By default, these actions are enabled. |

Supports update: yes

Discovery: supported
