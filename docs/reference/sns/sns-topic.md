# aws.sns.topic

**CloudFormation type:** `AWS::SNS::Topic`

The ``AWS::SNS::Topic`` resource creates a topic to which notifications can be published.

Region attribute: `region`

**Import ID:** `<region>/TopicArn` (AWS::SNS::Topic)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ArchivePolicy` | archive_policy | `map` | optional, computed, provider-chosen |  | The ``ArchivePolicy`` determines the number of days SNS retains messages in FIFO topics. You can set a retention period ranging from 1 to 365 days. This property is only applicable to FIFO topics; attempting to use it with standard topics will result in a creation failure. |
| `ContentBasedDeduplication` | content_based_deduplication | `boolean` | optional, computed, provider-chosen |  | ``ContentBasedDeduplication`` enables deduplication of messages based on their content for FIFO topics. By default, this property is set to false. If you create a FIFO topic with ``ContentBasedDeduplication`` set to false, you must provide a ``MessageDeduplicationId`` for each ``Publish`` action. When set to true, SNS automatically generates a ``MessageDeduplicationId`` using a SHA-256 hash of the message body (excluding message attributes). You can optionally override this generated value by specifying a ``MessageDeduplicationId`` in the ``Publish`` action. Note that this property only applies to FIFO topics; using it with standard topics will cause the creation to fail. |
| `DataProtectionPolicy` | data_protection_policy | `map` | optional, computed, provider-chosen |  | The body of the policy document you want to use for this topic. |
| `DeliveryStatusLogging` | delivery_status_logging | `list` | optional, computed, provider-chosen |  | The ``DeliveryStatusLogging`` configuration enables you to log the delivery status of messages sent from your Amazon SNS topic to subscribed endpoints with the following supported delivery protocols: |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The display name to use for an SNS topic with SMS subscriptions. The display name must be maximum 100 characters long, including hyphens (-), underscores (_), spaces, and tabs. |
| `FifoThroughputScope` | fifo_throughput_scope | `string` | optional, computed, provider-chosen |  | Specifies the throughput quota and deduplication behavior to apply for the FIFO topic. Valid values are ``Topic`` or ``MessageGroup``. |
| `FifoTopic` | fifo_topic | `boolean` | optional, computed, provider-chosen, replaces on change |  | Set to true to create a FIFO topic. |
| `KmsMasterKeyId` | kms_master_key_id | `string` | optional, computed, provider-chosen |  | The ID of an AWS managed customer master key (CMK) for SNS or a custom CMK. For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms). For more examples, see ``KeyId`` in the *API Reference*. |
| `SignatureVersion` | signature_version | `string` | optional, computed, provider-chosen |  | The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS. By default, ``SignatureVersion`` is set to ``1``. |
| `Subscription` |  | `list` | optional, computed, provider-chosen |  | The SNS subscriptions (endpoints) for this topic. |
| `Tags` |  | `map` | tags map |  | The list of tags to add to a new topic. |
| `TopicArn` | topic_arn | `string` | computed |  |  |
| `TopicName` | topic_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with ``.fifo``. |
| `TracingConfig` | tracing_config | `string` | optional, computed, provider-chosen |  | Tracing mode of an SNS topic. By default ``TracingConfig`` is set to ``PassThrough``, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to ``Active``, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. |

Supports update: yes

Discovery: supported
