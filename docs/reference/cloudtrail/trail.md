# aws.trail

**CloudFormation type:** `AWS::CloudTrail::Trail`

Creates a trail that specifies the settings for delivery of log data to an Amazon S3 bucket. A maximum of five trails can exist in a region, irrespective of the region in which they were created.

Region attribute: `region`

**Import ID:** `<region>/TrailName` (AWS::CloudTrail::Trail)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdvancedEventSelectors` | advanced_event_selectors | `list` | optional, computed, provider-chosen |  | The advanced event selectors that were used to select events for the data store. |
| `AggregationConfigurations` | aggregation_configurations | `list` | optional, computed, provider-chosen |  | Specifies the aggregation configuration to aggregate CloudTrail Events. A maximum of 1 aggregation configuration is allowed. |
| `Arn` |  | `string` | computed |  |  |
| `CloudWatchLogsLogGroupArn` | cloud_watch_logs_log_group_arn | `string` | optional, computed, provider-chosen | aws.loggroup.Arn | Specifies a log group name using an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered. Not required unless you specify CloudWatchLogsRoleArn. |
| `CloudWatchLogsRoleArn` | cloud_watch_logs_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group. |
| `EnableLogFileValidation` | enable_log_file_validation | `boolean` | optional, computed, provider-chosen |  | Specifies whether log file validation is enabled. The default is false. |
| `EventSelectors` | event_selectors | `list` | optional, computed, provider-chosen |  | Use event selectors to further specify the management and data event settings for your trail. By default, trails created without specific event selectors will be configured to log all read and write management events, and no data events. When an event occurs in your account, CloudTrail evaluates the event selector for all trails. For each trail, if the event matches any event selector, the trail processes and logs the event. If the event doesn't match any event selector, the trail doesn't log the event. You can configure up to five event selectors for a trail. |
| `IncludeGlobalServiceEvents` | include_global_service_events | `boolean` | optional, computed, provider-chosen |  | Specifies whether the trail is publishing events from global services such as IAM to the log files. |
| `InsightSelectors` | insight_selectors | `list` | optional, computed, provider-chosen |  | Lets you enable Insights event logging by specifying the Insights selectors that you want to enable on an existing trail. |
| `IsLogging` | is_logging | `boolean` | required |  | Whether the CloudTrail is currently logging AWS API calls. |
| `IsMultiRegionTrail` | is_multi_region_trail | `boolean` | optional, computed, provider-chosen |  | Specifies whether the trail applies only to the current region or to all regions. The default is false. If the trail exists only in the current region and this value is set to true, shadow trails (replications of the trail) will be created in the other regions. If the trail exists in all regions and this value is set to false, the trail will remain in the region where it was created, and its shadow trails in other regions will be deleted. As a best practice, consider using trails that log events in all regions. |
| `IsOrganizationTrail` | is_organization_trail | `boolean` | optional, computed, provider-chosen |  | Specifies whether the trail is created for all accounts in an organization in AWS Organizations, or only for the current AWS account. The default is false, and cannot be true unless the call is made on behalf of an AWS account that is the master account for an organization in AWS Organizations. |
| `KMSKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | Specifies the KMS key ID to use to encrypt the logs delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier. |
| `S3BucketName` | s3_bucket_name | `string` | required |  | Specifies the name of the Amazon S3 bucket designated for publishing log files. See Amazon S3 Bucket Naming Requirements. |
| `S3KeyPrefix` | s3_key_prefix | `string` | optional, computed, provider-chosen |  | Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery. For more information, see Finding Your CloudTrail Log Files. The maximum length is 200 characters. |
| `SnsTopicArn` | sns_topic_arn | `string` | computed |  |  |
| `SnsTopicName` | sns_topic_name | `string` | optional, computed, provider-chosen |  | Specifies the name of the Amazon SNS topic defined for notification of log file delivery. The maximum length is 256 characters. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TrailName` | trail_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
