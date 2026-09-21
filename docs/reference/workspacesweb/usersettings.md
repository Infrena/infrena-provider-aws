# aws.usersettings

**CloudFormation type:** `AWS::WorkSpacesWeb::UserSettings`

Definition of AWS::WorkSpacesWeb::UserSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/UserSettingsArn` (AWS::WorkSpacesWeb::UserSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen |  |  |
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `BrandingConfiguration` | branding_configuration | `map` | optional, computed, provider-chosen |  |  |
| `CookieSynchronizationConfiguration` | cookie_synchronization_configuration | `map` | optional, computed, provider-chosen |  |  |
| `CopyAllowed` | copy_allowed | `string` | required |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen |  |  |
| `DeepLinkAllowed` | deep_link_allowed | `string` | optional, computed, provider-chosen |  |  |
| `DisconnectTimeoutInMinutes` | disconnect_timeout_in_minutes | `float` | optional, computed, provider-chosen |  |  |
| `DownloadAllowed` | download_allowed | `string` | required |  |  |
| `IdleDisconnectTimeoutInMinutes` | idle_disconnect_timeout_in_minutes | `float` | optional, computed, provider-chosen |  |  |
| `PasteAllowed` | paste_allowed | `string` | required |  |  |
| `PrintAllowed` | print_allowed | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `ToolbarConfiguration` | toolbar_configuration | `map` | optional, computed, provider-chosen |  |  |
| `UploadAllowed` | upload_allowed | `string` | required |  |  |
| `UserSettingsArn` | user_settings_arn | `string` | computed |  |  |
| `WebAuthnAllowed` | web_authn_allowed | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
