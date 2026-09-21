# aws.glue.integration

**CloudFormation type:** `AWS::Glue::Integration`

Resource Type definition for AWS::Glue::Integration

Region attribute: `region`

**Import ID:** `<region>/IntegrationArn|IntegrationName` (AWS::Glue::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  | An optional set of non-secret key value pairs that contains additional contextual information about the data. |
| `CreateTime` | create_time | `string` | computed |  | The time (UTC) when the integration was created. |
| `DataFilter` | data_filter | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IntegrationArn` | integration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the integration. |
| `IntegrationConfig` | integration_config | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration settings for the integration. |
| `IntegrationName` | integration_name | `string` | required, replaces on change |  | The name of the integration. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used. |
| `SourceArn` | source_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the database to use as the source for replication |
| `Status` |  | `string` | computed |  | The status of the integration. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetArn` | target_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the Glue data warehouse to use as the target for replication |

Supports update: yes

Discovery: supported
