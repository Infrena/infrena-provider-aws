# aws.rds.integration

**CloudFormation type:** `AWS::RDS::Integration`

A zero-ETL integration with Amazon Redshift.

Region attribute: `region`

**Import ID:** `<region>/IntegrationArn` (AWS::RDS::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEncryptionContext` | additional_encryption_context | `map` | optional, computed, provider-chosen, replaces on change |  | An optional set of non-secret key–value pairs that contains additional contextual information about the data. |
| `CreateTime` | create_time | `string` | computed |  |  |
| `DataFilter` | data_filter | `string` | optional, computed, provider-chosen |  | Data filters for the integration. These filters determine which tables from the source database are sent to the target Amazon Redshift data warehouse. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the integration. |
| `IntegrationArn` | integration_arn | `string` | computed |  |  |
| `IntegrationName` | integration_name | `string` | optional, computed, provider-chosen |  | The name of the integration. |
| `KMSKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Key Management System (AWS KMS) key identifier for the key to use to encrypt the integration. If you don't specify an encryption key, RDS uses a default AWS owned key. |
| `SourceArn` | source_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the database to use as the source for replication. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags. For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*. |
| `TargetArn` | target_arn | `string` | required, replaces on change |  | The ARN of the Redshift data warehouse to use as the target for replication. |

Supports update: yes

Discovery: supported
