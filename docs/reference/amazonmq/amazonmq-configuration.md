# aws.amazonmq.configuration

**CloudFormation type:** `AWS::AmazonMQ::Configuration`

Resource Type definition for AWS::AmazonMQ::Configuration

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AmazonMQ::Configuration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Amazon MQ configuration. |
| `AuthenticationStrategy` | authentication_strategy | `string` | optional, computed, provider-chosen, replaces on change |  | The authentication strategy associated with the configuration. The default is SIMPLE. |
| `Data` |  | `string` | optional, computed, provider-chosen, write-only |  | The base64-encoded XML configuration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the configuration. |
| `EngineType` | engine_type | `string` | required, replaces on change |  | The type of broker engine. Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations. |
| `EngineVersion` | engine_version | `string` | optional, computed, provider-chosen, replaces on change |  | The version of the broker engine. |
| `Id` |  | `string` | computed |  | The ID of the Amazon MQ configuration. |
| `Name` |  | `string` | required, replaces on change |  | The name of the configuration. |
| `Revision` |  | `string` | computed |  | The revision number of the configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Create tags when creating the configuration. |

Supports update: yes

Discovery: supported
