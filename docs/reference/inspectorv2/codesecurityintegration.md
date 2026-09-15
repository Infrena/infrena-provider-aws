# aws.codesecurityintegration

**CloudFormation type:** `AWS::InspectorV2::CodeSecurityIntegration`

Inspector CodeSecurityIntegration resource schema

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::InspectorV2::CodeSecurityIntegration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Code Security Integration ARN |
| `AuthorizationUrl` | authorization_url | `string` | computed |  | Authorization URL for OAuth flow |
| `CreateIntegrationDetails` | create_integration_details | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Create Integration Details |
| `CreatedAt` | created_at | `string` | computed |  | Creation timestamp |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | Last update timestamp |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Code Security Integration name |
| `Status` |  | `string` | computed |  | Integration Status |
| `StatusReason` | status_reason | `string` | computed |  | Reason for the current status |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | Integration Type |
| `UpdateIntegrationDetails` | update_integration_details | `map` | optional, computed, provider-chosen, write-only |  | Update Integration Details |

Supports update: yes

Discovery: supported
