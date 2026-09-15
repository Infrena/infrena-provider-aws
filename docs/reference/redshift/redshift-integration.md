# aws.redshift.integration

**CloudFormation type:** `AWS::Redshift::Integration`

Integration from a source AWS service to a Redshift cluster

Region attribute: `region`

**Import ID:** `<region>/IntegrationArn` (AWS::Redshift::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  | An optional set of non-secret key–value pairs that contains additional contextual information about the data. |
| `CreateTime` | create_time | `string` | computed |  | The time (UTC) when the integration was created. |
| `IntegrationArn` | integration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the integration. |
| `IntegrationName` | integration_name | `string` | optional, computed, provider-chosen |  | The name of the integration. |
| `KMSKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | An KMS key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, the default AWS owned KMS key is used. |
| `SourceArn` | source_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the database to use as the source for replication |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetArn` | target_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the Redshift data warehouse to use as the target for replication |

Supports update: yes

Discovery: supported
