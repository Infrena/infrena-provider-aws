# aws.securityhub.connector

**CloudFormation type:** `AWS::SecurityHub::Connector`

Creates a connector to a third-party cloud provider in Security Hub CSPM. A connector establishes a connection between Security Hub CSPM and a third-party cloud provider, enabling Security Hub CSPM to ingest security findings and resource data from the connected environment.

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::SecurityHub::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | computed |  |  |
| `ConnectorId` | connector_id | `string` | computed |  |  |
| `ConnectorStatus` | connector_status | `string` | computed |  | The status of the connector |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `CreatedBy` | created_by | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the connector. |
| `EnablementStatus` | enablement_status | `string` | computed |  | The enablement status of the connector |
| `Issues` |  | `list` | computed |  |  |
| `LastCheckedAt` | last_checked_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `Message` |  | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  | The name of the connector. |
| `Provider` | provider_value | `map` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
