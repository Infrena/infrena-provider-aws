# aws.timestream.database

**CloudFormation type:** `AWS::Timestream::Database`

The AWS::Timestream::Database resource creates a Timestream database.

Region attribute: `region`

**Import ID:** `<region>/DatabaseName` (AWS::Timestream::Database)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DatabaseName` | database_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
