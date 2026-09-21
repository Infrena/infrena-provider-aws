# aws.delivery

**CloudFormation type:** `AWS::Logs::Delivery`

This structure contains information about one delivery in your account.

Region attribute: `region`

**Import ID:** `<region>/DeliveryId` (AWS::Logs::Delivery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) that uniquely identify AWS resource. |
| `DeliveryDestinationArn` | delivery_destination_arn | `string` | required, replaces on change | aws.deliverydestination.Arn | Amazon Resource Name (ARN) that uniquely identify AWS resource. |
| `DeliveryDestinationType` | delivery_destination_type | `string` | computed |  | Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose. |
| `DeliveryId` | delivery_id | `string` | computed |  | The unique ID that identifies this delivery in your account. |
| `DeliverySourceName` | delivery_source_name | `string` | required, replaces on change |  | The name of the delivery source that is associated with this delivery. |
| `FieldDelimiter` | field_delimiter | `string` | optional, computed, provider-chosen |  | The field delimiter to use between record fields when the final output format of a delivery is in Plain , W3C , or Raw format. |
| `RecordFields` | record_fields | `list` | optional, computed, provider-chosen |  | The list of record fields to be delivered to the destination, in order. If the delivery's log source has mandatory fields, they must be included in this list. |
| `S3EnableHiveCompatiblePath` | s3_enable_hive_compatible_path | `boolean` | optional, computed, provider-chosen |  | This parameter causes the S3 objects that contain delivered logs to use a prefix structure that allows for integration with Apache Hive. |
| `S3SuffixPath` | s3_suffix_path | `string` | optional, computed, provider-chosen |  | This string allows re-configuring the S3 object prefix to contain either static or variable sections. The valid variables to use in the suffix path will vary by each log source. See ConfigurationTemplate$allowedSuffixPathFields for more info on what values are supported in the suffix path for each log source. |
| `Tags` |  | `map` | tags map |  | The tags that have been assigned to this delivery. |

Supports update: yes

Discovery: supported
