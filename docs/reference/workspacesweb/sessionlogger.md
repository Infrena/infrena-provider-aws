# aws.sessionlogger

**CloudFormation type:** `AWS::WorkSpacesWeb::SessionLogger`

Definition of AWS::WorkSpacesWeb::SessionLogger Resource Type

Region attribute: `region`

**Import ID:** `<region>/SessionLoggerArn` (AWS::WorkSpacesWeb::SessionLogger)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AssociatedPortalArns` | associated_portal_arns | `list` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  |  |
| `CustomerManagedKey` | customer_managed_key | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  |  |
| `EventFilter` | event_filter | `string` | required |  |  |
| `LogConfiguration` | log_configuration | `map` | required |  |  |
| `SessionLoggerArn` | session_logger_arn | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
