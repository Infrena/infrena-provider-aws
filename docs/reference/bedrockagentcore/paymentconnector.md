# aws.paymentconnector

**CloudFormation type:** `AWS::BedrockAgentCore::PaymentConnector`

Resource Type definition for AWS::BedrockAgentCore::PaymentConnector

Region attribute: `region`

**Import ID:** `<region>/PaymentConnectorArn` (AWS::BedrockAgentCore::PaymentConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthorizationUrl` | authorization_url | `string` | computed |  | The URL the user must open to complete OAuth consent. Only present when ConnectorStatus is PENDING_AUTHENTICATION. |
| `ConnectorCreatedAt` | connector_created_at | `string` | computed |  | The timestamp when the connector was created |
| `ConnectorLastUpdatedAt` | connector_last_updated_at | `string` | computed |  | The timestamp when the connector was last updated |
| `ConnectorName` | connector_name | `string` | required, replaces on change |  | The name of the payment connector |
| `ConnectorStatus` | connector_status | `string` | computed |  |  |
| `ConnectorType` | connector_type | `string` | required, replaces on change |  |  |
| `CredentialProviderConfigurations` | credential_provider_configurations | `list` | optional, computed, provider-chosen |  | The credential provider configurations for the connector. Required when ProvisionMode is MANUAL or not specified. Empty for QUICK_CREATE until provisioning completes. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the payment connector |
| `PaymentConnectorArn` | payment_connector_arn | `string` | computed |  | Synthetic ARN for the payment connector (used for engine resolution) |
| `PaymentConnectorId` | payment_connector_id | `string` | computed |  | The unique identifier for the payment connector |
| `PaymentManagerId` | payment_manager_id | `string` | required, replaces on change | aws.paymentmanager.PaymentManagerId | The identifier of the parent payment manager |
| `ProvisionMode` | provision_mode | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The provision mode for creating the connector. MANUAL requires CredentialProviderConfigurations; QUICK_CREATE orchestrates OAuth consent and credential provisioning. |

Supports update: yes

Discovery: supported (parent resource required)
