# aws.subscriber

**CloudFormation type:** `AWS::SecurityLake::Subscriber`

Resource Type definition for AWS::SecurityLake::Subscriber

Region attribute: `region`

**Import ID:** `<region>/SubscriberArn` (AWS::SecurityLake::Subscriber)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessTypes` | access_types | `list` | required |  | The Amazon S3 or AWS Lake Formation access type. |
| `DataLakeArn` | data_lake_arn | `string` | required, replaces on change | aws.datalake.Arn | The ARN for the data lake. |
| `ResourceShareArn` | resource_share_arn | `string` | computed |  |  |
| `ResourceShareName` | resource_share_name | `string` | computed |  |  |
| `S3BucketArn` | s3_bucket_arn | `string` | computed |  |  |
| `Sources` |  | `list` | required |  | The supported AWS services from which logs and events are collected. |
| `SubscriberArn` | subscriber_arn | `string` | computed |  |  |
| `SubscriberDescription` | subscriber_description | `string` | optional, computed, provider-chosen |  | The description for your subscriber account in Security Lake. |
| `SubscriberIdentity` | subscriber_identity | `map` | required |  | The AWS identity used to access your data. |
| `SubscriberName` | subscriber_name | `string` | required |  | The name of your Security Lake subscriber account. |
| `SubscriberRoleArn` | subscriber_role_arn | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string. |

Supports update: yes

Discovery: supported
