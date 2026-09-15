# aws.inferencescheduler

**CloudFormation type:** `AWS::LookoutEquipment::InferenceScheduler`

Resource schema for LookoutEquipment InferenceScheduler.

Region attribute: `region`

**Import ID:** `<region>/InferenceSchedulerName` (AWS::LookoutEquipment::InferenceScheduler)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataDelayOffsetInMinutes` | data_delay_offset_in_minutes | `integer` | optional, computed, provider-chosen |  | A period of time (in minutes) by which inference on the data is delayed after the data starts. |
| `DataInputConfiguration` | data_input_configuration | `map` | required |  | Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location. |
| `DataOutputConfiguration` | data_output_configuration | `map` | required |  | Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output. |
| `DataUploadFrequency` | data_upload_frequency | `string` | required |  | How often data is uploaded to the source S3 bucket for the input data. |
| `InferenceSchedulerArn` | inference_scheduler_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the inference scheduler being created. |
| `InferenceSchedulerName` | inference_scheduler_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the inference scheduler being created. |
| `ModelName` | model_name | `string` | required, replaces on change |  | The name of the previously trained ML model being used to create the inference scheduler. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference. |
| `ServerSideKmsKeyId` | server_side_kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment. |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Any tags associated with the inference scheduler. |

Supports update: yes

Discovery: supported
