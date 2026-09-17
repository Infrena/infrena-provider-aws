# aws.deliverystream

**CloudFormation type:** `AWS::KinesisFirehose::DeliveryStream`

Resource Type definition for AWS::KinesisFirehose::DeliveryStream

Region attribute: `region`

**Import ID:** `<region>/DeliveryStreamName` (AWS::KinesisFirehose::DeliveryStream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonOpenSearchServerlessDestinationConfiguration` | amazon_open_search_serverless_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `AmazonopensearchserviceDestinationConfiguration` | amazonopensearchservice_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `Arn` |  | `string` | computed |  |  |
| `DatabaseSourceConfiguration` | database_source_configuration | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `DeliveryStreamEncryptionConfigurationInput` | delivery_stream_encryption_configuration_input | `map` | optional, computed, provider-chosen |  |  |
| `DeliveryStreamName` | delivery_stream_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DeliveryStreamType` | delivery_stream_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DirectPutSourceConfiguration` | direct_put_source_configuration | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ElasticsearchDestinationConfiguration` | elasticsearch_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `ExtendedS3DestinationConfiguration` | extended_s3_destination_configuration | `map` | optional, computed, provider-chosen |  |  |
| `HttpEndpointDestinationConfiguration` | http_endpoint_destination_configuration | `map` | optional, computed, provider-chosen |  |  |
| `IcebergDestinationConfiguration` | iceberg_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `KinesisStreamSourceConfiguration` | kinesis_stream_source_configuration | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `MSKSourceConfiguration` | msk_source_configuration | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `RedshiftDestinationConfiguration` | redshift_destination_configuration | `map` | optional, computed, provider-chosen |  |  |
| `S3DestinationConfiguration` | s3_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `SnowflakeDestinationConfiguration` | snowflake_destination_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `SplunkDestinationConfiguration` | splunk_destination_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
