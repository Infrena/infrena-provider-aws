# aws.logs.integration

**CloudFormation type:** `AWS::Logs::Integration`

Resource Schema for Logs Integration Resource

Region attribute: `region`

**Import ID:** `<region>/IntegrationName` (AWS::Logs::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IntegrationName` | integration_name | `string` | required, replaces on change |  | User provided identifier for integration, unique to the user account. |
| `IntegrationStatus` | integration_status | `string` | computed |  | Status of creation for the Integration and its resources |
| `IntegrationType` | integration_type | `string` | required, replaces on change |  | The type of the Integration. |
| `ResourceConfig` | resource_config | `map` | required, replaces on change, write-only |  | OpenSearchResourceConfig for the given Integration |

Supports update: no

Discovery: supported
