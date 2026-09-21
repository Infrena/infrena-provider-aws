# aws.plugin

**CloudFormation type:** `AWS::QBusiness::Plugin`

Definition of AWS::QBusiness::Plugin Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|PluginId` (AWS::QBusiness::Plugin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | optional, computed, provider-chosen, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `AuthConfiguration` | auth_configuration | `string` | required |  |  |
| `BuildStatus` | build_status | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CustomPluginConfiguration` | custom_plugin_configuration | `map` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `PluginArn` | plugin_arn | `string` | computed |  |  |
| `PluginId` | plugin_id | `string` | computed |  |  |
| `ServerUrl` | server_url | `string` | optional, computed, provider-chosen |  |  |
| `State` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
