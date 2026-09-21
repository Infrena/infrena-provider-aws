# aws.deliverychannel

**CloudFormation type:** `AWS::Config::DeliveryChannel`

Resource type definition for AWS::Config::DeliveryChannel

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Config::DeliveryChannel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigSnapshotDeliveryProperties` | config_snapshot_delivery_properties | `map` | optional, computed, provider-chosen |  | The options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the delivery channel. By default, AWS Config assigns the name "default" when creating the delivery channel. To change the delivery channel name, you must use the DeleteDeliveryChannel action to delete your current delivery channel, and then you must use the PutDeliveryChannel command to create a delivery channel that has the desired name. |
| `S3BucketName` | s3_bucket_name | `string` | required |  | The name of the Amazon S3 bucket to which AWS Config delivers configuration snapshots and configuration history files. |
| `S3KeyPrefix` | s3_key_prefix | `string` | optional, computed, provider-chosen |  | The prefix for the specified Amazon S3 bucket. |
| `S3KmsKeyArn` | s3_kms_key_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS ) AWS KMS key (KMS key) used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket. |
| `SnsTopicARN` | sns_topic_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the Amazon SNS topic to which AWS Config sends notifications about configuration changes. |

Supports update: yes

Discovery: supported
