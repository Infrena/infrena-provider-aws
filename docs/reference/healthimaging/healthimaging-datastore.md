# aws.healthimaging.datastore

**CloudFormation type:** `AWS::HealthImaging::Datastore`

Definition of AWS::HealthImaging::Datastore Resource Type

Region attribute: `region`

**Import ID:** `<region>/DatastoreId` (AWS::HealthImaging::Datastore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the data store was created. |
| `DatastoreArn` | datastore_arn | `string` | computed |  | The Datastore's ARN. |
| `DatastoreId` | datastore_id | `string` | computed |  |  |
| `DatastoreName` | datastore_name | `string` | optional, computed, provider-chosen, replaces on change |  | User friendly name for Datastore. |
| `DatastoreStatus` | datastore_status | `string` | computed |  | A string to denote the Datastore's state. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | ARN referencing a KMS key or KMS key alias. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A Map of key value pairs for Tags. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the data store was created. |

Supports update: no

Discovery: supported
