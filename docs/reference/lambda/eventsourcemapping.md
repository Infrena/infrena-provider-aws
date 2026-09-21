# aws.eventsourcemapping

**CloudFormation type:** `AWS::Lambda::EventSourceMapping`

The ``AWS::Lambda::EventSourceMapping`` resource creates a mapping between an event source and an LAMlong function. LAM reads items from the event source and triggers the function.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Lambda::EventSourceMapping)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonManagedKafkaEventSourceConfig` | amazon_managed_kafka_event_source_config | `map` | optional, computed, provider-chosen |  | Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source. |
| `BatchSize` | batch_size | `integer` | optional, computed, provider-chosen |  | The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB). |
| `BisectBatchOnFunctionError` | bisect_batch_on_function_error | `boolean` | optional, computed, provider-chosen |  | (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry. The default value is false. |
| `DestinationConfig` | destination_config | `map` | optional, computed, provider-chosen |  | A configuration object that specifies the destination of an event after Lambda processes it. For more information, see [Adding a destination](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-async-destinations). |
| `DocumentDBEventSourceConfig` | document_db_event_source_config | `map` | optional, computed, provider-chosen |  | Specific configuration settings for a DocumentDB event source. |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  | When true, the event source mapping is active. When false, Lambda pauses polling and invocation. |
| `EventSourceArn` | event_source_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the event source. |
| `EventSourceMappingArn` | event_source_mapping_arn | `string` | computed |  |  |
| `FilterCriteria` | filter_criteria | `map` | optional, computed, provider-chosen |  | An object that contains the filters for an event source. |
| `FunctionName` | function_name | `string` | required |  | The name or ARN of the Lambda function. |
| `FunctionResponseTypes` | function_response_types | `list` | optional, computed, provider-chosen |  | (Kinesis, DynamoDB Streams, and SQS) A list of current response type enums applied to the event source mapping. |
| `Id` |  | `string` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The ARN of the KMSlong (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics). |
| `LoggingConfig` | logging_config | `map` | optional, computed, provider-chosen |  | The function's Amazon CloudWatch Logs configuration settings. |
| `MaximumBatchingWindowInSeconds` | maximum_batching_window_in_seconds | `integer` | optional, computed, provider-chosen |  | The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function. |
| `MaximumRecordAgeInSeconds` | maximum_record_age_in_seconds | `integer` | optional, computed, provider-chosen |  | (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records. |
| `MaximumRetryAttempts` | maximum_retry_attempts | `integer` | optional, computed, provider-chosen |  | (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka) Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source. |
| `MetricsConfig` | metrics_config | `map` | optional, computed, provider-chosen |  | The metrics configuration for your event source. Use this configuration object to define which metrics you want your event source mapping to produce. |
| `ParallelizationFactor` | parallelization_factor | `integer` | optional, computed, provider-chosen |  | (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard. The default value is 1. |
| `ProvisionedPollerConfig` | provisioned_poller_config | `map` | optional, computed, provider-chosen |  | The [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode) configuration for the event source. Use Provisioned Mode to customize the minimum and maximum number of event pollers for your event source. |
| `Queues` |  | `list` | optional, computed, provider-chosen |  | (Amazon MQ) The name of the Amazon MQ broker destination queue to consume. |
| `ScalingConfig` | scaling_config | `map` | optional, computed, provider-chosen |  | (Amazon SQS only) The scaling configuration for the event source. To remove the configuration, pass an empty value. |
| `SelfManagedEventSource` | self_managed_event_source | `map` | optional, computed, provider-chosen, replaces on change |  | The self-managed Apache Kafka cluster for your event source. |
| `SelfManagedKafkaEventSourceConfig` | self_managed_kafka_event_source_config | `map` | optional, computed, provider-chosen |  | Specific configuration settings for a self-managed Apache Kafka event source. |
| `SourceAccessConfigurations` | source_access_configurations | `list` | optional, computed, provider-chosen |  | An array of the authentication protocol, VPC components, or virtual host to secure and define your event source. |
| `StartingPosition` | starting_position | `string` | optional, computed, provider-chosen, replaces on change |  | The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB. |
| `StartingPositionTimestamp` | starting_position_timestamp | `float` | optional, computed, provider-chosen, replaces on change |  | With ``StartingPosition`` set to ``AT_TIMESTAMP``, the time from which to start reading, in Unix time seconds. ``StartingPositionTimestamp`` cannot be in the future. |
| `Tags` |  | `map` | tags map |  | A list of tags to add to the event source mapping. |
| `Topics` |  | `list` | optional, computed, provider-chosen |  | The name of the Kafka topic. |
| `TumblingWindowInSeconds` | tumbling_window_in_seconds | `integer` | optional, computed, provider-chosen |  | (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds indicates no tumbling window. |

Supports update: yes

Discovery: supported
