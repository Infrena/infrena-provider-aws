# aws.sqs.queue

**CloudFormation type:** `AWS::SQS::Queue`

The ``AWS::SQS::Queue`` resource creates an SQS standard or FIFO queue.

Region attribute: `region`

**Import ID:** `<region>/QueueUrl` (AWS::SQS::Queue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ContentBasedDeduplication` | content_based_deduplication | `boolean` | optional, computed, provider-chosen |  | For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the ``ContentBasedDeduplication`` attribute for the ``CreateQueue`` action in the *API Reference*. |
| `DeduplicationScope` | deduplication_scope | `string` | optional, computed, provider-chosen |  | For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level. Valid values are ``messageGroup`` and ``queue``. |
| `DelaySeconds` | delay_seconds | `integer` | optional, computed, provider-chosen |  | The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of ``0`` to ``900`` (15 minutes). The default value is ``0``. |
| `FifoQueue` | fifo_queue | `boolean` | optional, computed, provider-chosen, replaces on change |  | If set to true, creates a FIFO queue. If you don't specify this property, SQS creates a standard queue. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Developer Guide*. |
| `FifoThroughputLimit` | fifo_throughput_limit | `string` | optional, computed, provider-chosen |  | For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are ``perQueue`` and ``perMessageGroupId``. |
| `KmsDataKeyReusePeriodSeconds` | kms_data_key_reuse_period_seconds | `integer` | optional, computed, provider-chosen |  | The length of time in seconds for which SQS can reuse a data key to encrypt or decrypt messages before calling KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes). |
| `KmsMasterKeyId` | kms_master_key_id | `string` | optional, computed, provider-chosen |  | The ID of an AWS Key Management Service (KMS) for SQS, or a custom KMS. To use the AWS managed KMS for SQS, specify a (default) alias ARN, alias name (for example ``alias/aws/sqs``), key ARN, or key ID. For more information, see the following: |
| `MaximumMessageSize` | maximum_message_size | `integer` | optional, computed, provider-chosen |  | The limit of how many bytes that a message can contain before SQS rejects it. You can specify an integer from 1,024 bytes (1 KiB) to 1,048,576 bytes (1 MiB). Default: 1,048,576 bytes (1 MiB). |
| `MessageRetentionPeriod` | message_retention_period | `integer` | optional, computed, provider-chosen |  | The number of seconds that SQS retains a message. You can specify an integer value from ``60`` seconds (1 minute) to ``1,209,600`` seconds (14 days). The default value is ``345,600`` seconds (4 days). |
| `QueueName` | queue_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the ``.fifo`` suffix. For more information, see [Amazon SQS FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-fifo-queues.html) in the *Developer Guide*. |
| `QueueUrl` | queue_url | `string` | computed |  |  |
| `ReceiveMessageWaitTimeSeconds` | receive_message_wait_time_seconds | `integer` | optional, computed, provider-chosen |  | Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Developer Guide*. |
| `RedriveAllowPolicy` | redrive_allow_policy | `string` | optional, computed, provider-chosen |  | The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object. The parameters are as follows: |
| `RedrivePolicy` | redrive_policy | `string` | optional, computed, provider-chosen |  | The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object. The parameters are as follows: |
| `SqsManagedSseEnabled` | sqs_managed_sse_enabled | `boolean` | optional, computed, provider-chosen |  | Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)). When ``SqsManagedSseEnabled`` is not defined, ``SSE-SQS`` encryption is enabled by default. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags that you attach to this queue. For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *User Guide*. |
| `VisibilityTimeout` | visibility_timeout | `integer` | optional, computed, provider-chosen |  | The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. |

Supports update: yes

Discovery: supported
