# aws.connectorv2

**CloudFormation type:** `AWS::SecurityHub::ConnectorV2`

Resource schema for AWS::SecurityHub::ConnectorV2

Region attribute: `region`

**Import ID:** `<region>/ConnectorArn` (AWS::SecurityHub::ConnectorV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | computed |  | The ARN of the connector |
| `ConnectorId` | connector_id | `string` | computed |  | The ID of the connector |
| `ConnectorStatus` | connector_status | `string` | computed |  | The status of the connector |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the connector |
| `EnablementStatus` | enablement_status | `string` | computed |  | The enablement status of the connector |
| `EnablementStatusReason` | enablement_status_reason | `string` | computed |  | The reason for the enablement status of the connector |
| `Issues` |  | `list` | computed |  | The list of health issues associated with the connector |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of KMS key used for the connector |
| `LastCheckedAt` | last_checked_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp formatted in ISO8601 |
| `Message` |  | `string` | computed |  | The message of the connector status change |
| `Name` |  | `string` | required, replaces on change |  | The name of the connector |
| `Provider` | provider_value | `map` | required |  | The third-party provider configuration for the connector |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
