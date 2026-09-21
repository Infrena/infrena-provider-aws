# aws.browsersettings

**CloudFormation type:** `AWS::WorkSpacesWeb::BrowserSettings`

Definition of AWS::WorkSpacesWeb::BrowserSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/BrowserSettingsArn` (AWS::WorkSpacesWeb::BrowserSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `BrowserPolicy` | browser_policy | `string` | optional, computed, provider-chosen |  |  |
| `BrowserSettingsArn` | browser_settings_arn | `string` | computed |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `WebContentFilteringPolicy` | web_content_filtering_policy | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
