# aws.ivschat.loggingconfiguration

**CloudFormation type:** `AWS::IVSChat::LoggingConfiguration`

Resource type definition for AWS::IVSChat::LoggingConfiguration.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVSChat::LoggingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | LoggingConfiguration ARN is automatically generated on creation and assigned as the unique identifier. |
| `DestinationConfiguration` | destination_configuration | `map` | required |  | Destination configuration for IVS Chat logging. |
| `Id` |  | `string` | computed |  | The system-generated ID of the logging configuration. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the logging configuration. The value does not need to be unique. |
| `State` |  | `string` | computed |  | The state of the logging configuration. When the state is ACTIVE, the configuration is ready to log chat content. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
