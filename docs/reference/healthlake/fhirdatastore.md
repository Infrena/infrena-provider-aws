# aws.fhirdatastore

**CloudFormation type:** `AWS::HealthLake::FHIRDatastore`

HealthLake FHIR Datastore

Region attribute: `region`

**Import ID:** `<region>/DatastoreId` (AWS::HealthLake::FHIRDatastore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `map` | computed |  | The time that a Data Store was created. |
| `DatastoreArn` | datastore_arn | `string` | computed |  | The Amazon Resource Name used in the creation of the Data Store. |
| `DatastoreEndpoint` | datastore_endpoint | `string` | computed |  | The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL. |
| `DatastoreId` | datastore_id | `string` | computed |  | The AWS-generated ID number for the Data Store. |
| `DatastoreName` | datastore_name | `string` | optional, computed, provider-chosen, replaces on change |  | The user-generated name for the Data Store. |
| `DatastoreStatus` | datastore_status | `string` | computed |  | The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'. |
| `DatastoreTypeVersion` | datastore_type_version | `string` | required, replaces on change |  | The FHIR version. Only R4 version data is supported. |
| `IdentityProviderConfiguration` | identity_provider_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The identity provider configuration for the datastore |
| `PreloadDataConfig` | preload_data_config | `map` | optional, computed, provider-chosen, replaces on change |  | The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported. |
| `SseConfiguration` | sse_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The server-side encryption key configuration for a customer provided encryption key. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
