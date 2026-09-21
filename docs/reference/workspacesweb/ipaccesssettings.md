# aws.ipaccesssettings

**CloudFormation type:** `AWS::WorkSpacesWeb::IpAccessSettings`

Definition of AWS::WorkSpacesWeb::IpAccessSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/IpAccessSettingsArn` (AWS::WorkSpacesWeb::IpAccessSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `IpAccessSettingsArn` | ip_access_settings_arn | `string` | computed |  |  |
| `IpRules` | ip_rules | `list` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
