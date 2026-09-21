# aws.dms.endpoint

**CloudFormation type:** `AWS::DMS::Endpoint`

Resource Type definition for AWS::DMS::Endpoint

Region attribute: `region`

**Import ID:** `<region>/EndpointArn` (AWS::DMS::Endpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | optional, computed, provider-chosen | aws.dms.certificate.CertificateArn | The Amazon Resource Name (ARN) for the certificate. |
| `DatabaseName` | database_name | `string` | optional, computed, provider-chosen |  | The name of the endpoint database. For a MySQL source or target endpoint, don't specify DatabaseName. To migrate to a specific database, use this setting and targetDbType. |
| `DocDbSettings` | doc_db_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a DocumentDB endpoint. |
| `DynamoDbSettings` | dynamo_db_settings | `map` | optional, computed, provider-chosen |  | Provides information, including the Amazon Resource Name (ARN) of the IAM role used to define an Amazon DynamoDB target endpoint. |
| `ElasticsearchSettings` | elasticsearch_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an OpenSearch endpoint. |
| `EndpointArn` | endpoint_arn | `string` | computed |  | The endpoint ARN. |
| `EndpointIdentifier` | endpoint_identifier | `string` | optional, computed, provider-chosen |  | The database endpoint identifier. Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens. |
| `EndpointType` | endpoint_type | `string` | required |  | The type of endpoint. Valid values are source and target. |
| `EngineName` | engine_name | `string` | required |  | The type of engine for the endpoint, depending on the EndpointType value. |
| `ExternalId` | external_id | `string` | computed |  | A value that can be used for cross-account validation. |
| `ExtraConnectionAttributes` | extra_connection_attributes | `string` | optional, computed, provider-chosen, write-only |  | Additional attributes associated with the connection |
| `GcpMySQLSettings` | gcp_my_sql_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a GCP MySQL endpoint. |
| `IbmDb2Settings` | ibm_db2_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an IBMDB2 endpoint. |
| `KafkaSettings` | kafka_settings | `map` | optional, computed, provider-chosen |  | Provides information that describes an Apache Kafka endpoint. |
| `KinesisSettings` | kinesis_settings | `map` | optional, computed, provider-chosen |  | Provides information that describes an Amazon Kinesis Data Stream endpoint. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.If you don't specify a value for the KmsKeyId parameter, AWS DMS uses your default encryption key. |
| `MicrosoftSqlServerSettings` | microsoft_sql_server_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a Microsoft SQL Server endpoint. |
| `MongoDbSettings` | mongo_db_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a MongoDB endpoint. |
| `MySqlSettings` | my_sql_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a MySQL endpoint. |
| `NeptuneSettings` | neptune_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an Amazon Neptune endpoint |
| `OracleSettings` | oracle_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an Oracle endpoint |
| `Password` |  | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password to be used to log in to the endpoint database. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The port used by the endpoint database. |
| `PostgreSqlSettings` | postgre_sql_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a PostgreSQL endpoint |
| `RedisSettings` | redis_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a Redis target endpoint. |
| `RedshiftSettings` | redshift_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an Amazon Redshift endpoint. |
| `ResourceIdentifier` | resource_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A display name for the resource identifier at the end of the EndpointArn response parameter that is returned in the created Endpoint object. |
| `S3Settings` | s3_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines an Amazon S3 endpoint. |
| `ServerName` | server_name | `string` | optional, computed, provider-chosen |  | The name of the server where the endpoint database resides. |
| `SslMode` | ssl_mode | `string` | optional, computed, provider-chosen |  | The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is none. |
| `SybaseSettings` | sybase_settings | `map` | optional, computed, provider-chosen |  | Provides information that defines a SAP ASE endpoint. |
| `Tags` |  | `map` | tags map |  | One or more tags to be assigned to the endpoint. |
| `Username` |  | `string` | optional, computed, provider-chosen |  | The user name to be used to log in to the endpoint database. |

Supports update: yes

Discovery: supported
