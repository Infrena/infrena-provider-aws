# aws.workspacesweb.portal

**CloudFormation type:** `AWS::WorkSpacesWeb::Portal`

Definition of AWS::WorkSpacesWeb::Portal Resource Type

Region attribute: `region`

**Import ID:** `<region>/PortalArn` (AWS::WorkSpacesWeb::Portal)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AuthenticationType` | authentication_type | `string` | optional, computed, provider-chosen |  |  |
| `BrowserSettingsArn` | browser_settings_arn | `string` | optional, computed, provider-chosen | aws.browsersettings.BrowserSettingsArn |  |
| `BrowserType` | browser_type | `string` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DataProtectionSettingsArn` | data_protection_settings_arn | `string` | optional, computed, provider-chosen | aws.dataprotectionsettings.DataProtectionSettingsArn |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `InstanceType` | instance_type | `string` | optional, computed, provider-chosen |  |  |
| `IpAccessSettingsArn` | ip_access_settings_arn | `string` | optional, computed, provider-chosen | aws.ipaccesssettings.IpAccessSettingsArn |  |
| `MaxConcurrentSessions` | max_concurrent_sessions | `float` | optional, computed, provider-chosen |  |  |
| `NetworkSettingsArn` | network_settings_arn | `string` | optional, computed, provider-chosen | aws.networksettings.NetworkSettingsArn |  |
| `PortalArn` | portal_arn | `string` | computed |  |  |
| `PortalCustomDomain` | portal_custom_domain | `string` | optional, computed, provider-chosen |  |  |
| `PortalEndpoint` | portal_endpoint | `string` | computed |  |  |
| `PortalStatus` | portal_status | `string` | computed |  |  |
| `RendererType` | renderer_type | `string` | computed |  |  |
| `ServiceProviderSamlMetadata` | service_provider_saml_metadata | `string` | computed |  |  |
| `SessionLoggerArn` | session_logger_arn | `string` | optional, computed, provider-chosen | aws.sessionlogger.SessionLoggerArn |  |
| `StatusReason` | status_reason | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TrustStoreArn` | trust_store_arn | `string` | optional, computed, provider-chosen | aws.workspacesweb.truststore.TrustStoreArn |  |
| `UserAccessLoggingSettingsArn` | user_access_logging_settings_arn | `string` | optional, computed, provider-chosen | aws.useraccessloggingsettings.UserAccessLoggingSettingsArn |  |
| `UserSettingsArn` | user_settings_arn | `string` | optional, computed, provider-chosen | aws.usersettings.UserSettingsArn |  |

Supports update: yes

Discovery: supported
