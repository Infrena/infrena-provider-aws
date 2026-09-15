# aws.config.connector

**CloudFormation type:** `AWS::Config::Connector`

Resource Type definition for AWS::Config::Connector

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Config::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the connector. |
| `ConnectorConfiguration` | connector_configuration | `map` | required, replaces on change |  | The configuration for the connector. Specify the third-party cloud provider configuration. |
| `CreatedTime` | created_time | `string` | computed |  | The time at which the connector was created. |
| `Name` |  | `string` | computed |  | The name of the connector. AWS Config automatically assigns the name when creating the Connector. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the connector. |

Supports update: yes

Discovery: supported
