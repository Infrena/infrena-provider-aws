# aws.inspectorv2.connector

**CloudFormation type:** `AWS::InspectorV2::Connector`

Creates and manages a multi-cloud connector for Amazon Inspector, enabling vulnerability scanning of Azure resources.

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::InspectorV2::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the connector. |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the connector was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Optional description of the connector. |
| `EnablementStatus` | enablement_status | `string` | computed |  | The enablement status of the connector. |
| `EnablementStatusReason` | enablement_status_reason | `string` | computed |  | Reason for the current enablement status, if applicable. |
| `Health` |  | `map` | computed |  | Health status of the connector. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | Timestamp when the connector was last updated. |
| `Name` |  | `string` | required, replaces on change |  | Display name for the connector. |
| `Provider` | provider_value | `string` | required, replaces on change |  | The cloud provider for this connector. |
| `ProviderConfiguration` | provider_configuration | `map` | required |  | Provider-specific configuration including regions and scope. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the connector. |

Supports update: yes

Discovery: supported
