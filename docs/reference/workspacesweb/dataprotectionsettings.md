# aws.dataprotectionsettings

**CloudFormation type:** `AWS::WorkSpacesWeb::DataProtectionSettings`

Definition of AWS::WorkSpacesWeb::DataProtectionSettings Resource Type

Region attribute: `region`

**Import ID:** `<region>/DataProtectionSettingsArn` (AWS::WorkSpacesWeb::DataProtectionSettings)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DataProtectionSettingsArn` | data_protection_settings_arn | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `InlineRedactionConfiguration` | inline_redaction_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
