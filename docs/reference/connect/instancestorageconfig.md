# aws.instancestorageconfig

**CloudFormation type:** `AWS::Connect::InstanceStorageConfig`

Resource Type definition for AWS::Connect::InstanceStorageConfig

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|AssociationId|ResourceType` (AWS::Connect::InstanceStorageConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | An associationID is automatically generated when a storage config is associated with an instance |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | Connect Instance ID with which the storage config will be associated |
| `KinesisFirehoseConfig` | kinesis_firehose_config | `map` | optional, computed, provider-chosen |  |  |
| `KinesisStreamConfig` | kinesis_stream_config | `map` | optional, computed, provider-chosen |  |  |
| `KinesisVideoStreamConfig` | kinesis_video_stream_config | `map` | optional, computed, provider-chosen |  |  |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | Specifies the type of storage resource available for the instance |
| `S3Config` | s3_config | `map` | optional, computed, provider-chosen |  |  |
| `StorageType` | storage_type | `string` | required |  | Specifies the storage type to be associated with the instance |

Supports update: yes

Discovery: supported
