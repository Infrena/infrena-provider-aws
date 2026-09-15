# aws.workerconfiguration

**CloudFormation type:** `AWS::KafkaConnect::WorkerConfiguration`

The configuration of the workers, which are the processes that run the connector logic.

Region attribute: `region`

**Import ID:** `<region>/WorkerConfigurationArn` (AWS::KafkaConnect::WorkerConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A summary description of the worker configuration. |
| `Name` |  | `string` | required, replaces on change |  | The name of the worker configuration. |
| `PropertiesFileContent` | properties_file_content | `string` | required, replaces on change |  | Base64 encoded contents of connect-distributed.properties file. |
| `Revision` |  | `integer` | computed |  | The description of a revision of the worker configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource |
| `WorkerConfigurationArn` | worker_configuration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the custom configuration. |

Supports update: yes

Discovery: supported
