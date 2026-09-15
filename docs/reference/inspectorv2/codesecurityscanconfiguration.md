# aws.codesecurityscanconfiguration

**CloudFormation type:** `AWS::InspectorV2::CodeSecurityScanConfiguration`

Inspector CodeSecurityScanConfiguration resource schema

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::InspectorV2::CodeSecurityScanConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Code Security Scan Configuration ARN |
| `Configuration` |  | `map` | optional, computed, provider-chosen |  | Code Security Scan Configuration |
| `Level` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Configuration Level |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Code Security Scan Configuration name |
| `ScopeSettings` | scope_settings | `map` | optional, computed, provider-chosen, replaces on change |  | Scope Settings |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
