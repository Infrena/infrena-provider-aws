# aws.streamconsumer

**CloudFormation type:** `AWS::Kinesis::StreamConsumer`

Resource Type definition for AWS::Kinesis::StreamConsumer

Region attribute: `region`

**Import ID:** `<region>/ConsumerARN` (AWS::Kinesis::StreamConsumer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConsumerARN` | consumer_arn | `string` | computed |  | The ARN returned by Kinesis Data Streams when you registered the consumer. If you don't know the ARN of the consumer that you want to deregister, you can use the ListStreamConsumers operation to get a list of the descriptions of all the consumers that are currently registered with a given data stream. The description of a consumer contains its ARN. |
| `ConsumerCreationTimestamp` | consumer_creation_timestamp | `string` | computed |  | Timestamp when the consumer was created. |
| `ConsumerName` | consumer_name | `string` | required, replaces on change |  | The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams. |
| `ConsumerStatus` | consumer_status | `string` | computed |  | A consumer can't read data while in the CREATING or DELETING states. Valid Values: CREATING \| DELETING \| ACTIVE |
| `StreamARN` | stream_arn | `string` | required, replaces on change |  | The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer. |

Supports update: no

Discovery: supported
