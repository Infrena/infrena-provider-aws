# aws.streamprocessor

**CloudFormation type:** `AWS::Rekognition::StreamProcessor`

The AWS::Rekognition::StreamProcessor type is used to create an Amazon Rekognition StreamProcessor that you can use to analyze streaming videos.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Rekognition::StreamProcessor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the stream processor |
| `BoundingBoxRegionsOfInterest` | bounding_box_regions_of_interest | `list` | optional, computed, provider-chosen, replaces on change |  | The BoundingBoxRegionsOfInterest specifies an array of bounding boxes of interest in the video frames to analyze, as part of connected home feature. If an object is partially in a region of interest, Rekognition will tag it as detected if the overlap of the object with the region-of-interest is greater than 20%. |
| `ConnectedHomeSettings` | connected_home_settings | `map` | optional, computed, provider-chosen, replaces on change |  | Connected home settings to use on a streaming video. Note that either ConnectedHomeSettings or FaceSearchSettings should be set. Not both |
| `DataSharingPreference` | data_sharing_preference | `map` | optional, computed, provider-chosen, replaces on change |  | Indicates whether Rekognition is allowed to store the video stream data for model-training. |
| `FaceSearchSettings` | face_search_settings | `map` | optional, computed, provider-chosen, replaces on change |  | Face search settings to use on a streaming video. Note that either FaceSearchSettings or ConnectedHomeSettings should be set. Not both |
| `KinesisDataStream` | kinesis_data_stream | `map` | optional, computed, provider-chosen, replaces on change |  | The Amazon Kinesis Data Stream stream to which the Amazon Rekognition stream processor streams the analysis results, as part of face search feature. |
| `KinesisVideoStream` | kinesis_video_stream | `map` | required, replaces on change |  | The Kinesis Video Stream that streams the source video. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The KMS key that is used by Rekognition to encrypt any intermediate customer metadata and store in the customer's S3 bucket. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor. |
| `NotificationChannel` | notification_channel | `map` | optional, computed, provider-chosen, replaces on change |  | The ARN of the SNS notification channel where events of interests are published, as part of connected home feature. |
| `PolygonRegionsOfInterest` | polygon_regions_of_interest | `list` | optional, computed, provider-chosen, replaces on change |  | The PolygonRegionsOfInterest specifies a set of polygon areas of interest in the video frames to analyze, as part of connected home feature. Each polygon is in turn, an ordered list of Point |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | ARN of the IAM role that allows access to the stream processor, and provides Rekognition read permissions for KVS stream and write permissions to S3 bucket and SNS topic. |
| `S3Destination` | s3_destination | `map` | optional, computed, provider-chosen, replaces on change |  | The S3 location in customer's account where inference output & artifacts are stored, as part of connected home feature. |
| `Status` |  | `string` | computed |  | Current status of the stream processor. |
| `StatusMessage` | status_message | `string` | computed |  | Detailed status message about the stream processor. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
