# aws.appflow.connector

**CloudFormation type:** `AWS::AppFlow::Connector`

Resource schema for AWS::AppFlow::Connector

Region attribute: `region`

**Import ID:** `<region>/ConnectorLabel` (AWS::AppFlow::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | computed |  | The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account. |
| `ConnectorLabel` | connector_label | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the connector. The name is unique for each ConnectorRegistration in your AWS account. |
| `ConnectorProvisioningConfig` | connector_provisioning_config | `map` | required |  | Contains information about the configuration of the connector being registered. |
| `ConnectorProvisioningType` | connector_provisioning_type | `string` | required |  | The provisioning type of the connector. Currently the only supported value is LAMBDA. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description about the connector that's being registered. |

Supports update: yes

Discovery: supported
