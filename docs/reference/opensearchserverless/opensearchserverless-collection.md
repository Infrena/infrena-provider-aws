# aws.opensearchserverless.collection

**CloudFormation type:** `AWS::OpenSearchServerless::Collection`

Amazon OpenSearchServerless collection resource

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::OpenSearchServerless::Collection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the collection. |
| `CollectionEndpoint` | collection_endpoint | `string` | computed |  | The endpoint for the collection. |
| `CollectionGroupName` | collection_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the collection group to associate with the collection. |
| `DashboardEndpoint` | dashboard_endpoint | `string` | computed |  | The OpenSearch Dashboards endpoint for the collection. |
| `DeletionProtection` | deletion_protection | `string` | optional, computed, provider-chosen |  | The deletion protection state of the collection |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the collection |
| `EncryptionConfig` | encryption_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Encryption settings for the collection |
| `FipsEndpoints` | fips_endpoints | `map` | computed |  |  |
| `Id` |  | `string` | computed |  | The identifier of the collection |
| `KmsKeyArn` | kms_key_arn | `string` | computed |  | Key Management Service key used to encrypt the collection. |
| `Name` |  | `string` | required, replaces on change |  | The name of the collection. |
| `StandbyReplicas` | standby_replicas | `string` | optional, computed, provider-chosen, replaces on change |  | The possible standby replicas for the collection |
| `Tags` |  | `map` | replaces on change, write-only, tags map |  | List of tags to be added to the resource |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The possible types for the collection |
| `VectorOptions` | vector_options | `map` | optional, computed, provider-chosen |  | Vector search configuration options for the collection |

Supports update: yes

Discovery: supported
