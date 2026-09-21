# aws.dynamodb.table

**CloudFormation type:** `AWS::DynamoDB::Table`

The ``AWS::DynamoDB::Table`` resource creates a DDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *API Reference*.

Region attribute: `region`

**Import ID:** `<region>/TableName` (AWS::DynamoDB::Table)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AttributeDefinitions` | attribute_definitions | `list` | optional, computed, provider-chosen |  | A list of attributes that describe the key schema for the table and indexes. |
| `BillingMode` | billing_mode | `string` | optional, computed, provider-chosen |  | Specify how you are charged for read and write throughput and how you manage capacity. |
| `ContributorInsightsSpecification` | contributor_insights_specification | `map` | optional, computed, provider-chosen |  | Configures contributor insights settings for a table or one of its indexes. |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | Determines if a table is protected from deletion. When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Developer Guide*. |
| `GlobalSecondaryIndexes` | global_secondary_indexes | `list` | optional, computed, provider-chosen |  | Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes. |
| `ImportSourceSpecification` | import_source_specification | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies the properties of data being imported from the S3 bucket source to the table. |
| `KeySchema` | key_schema | `string` | required |  | Specifies the attributes that make up the primary key for the table. The attributes in the ``KeySchema`` property must also be defined in the ``AttributeDefinitions`` property. |
| `KinesisStreamSpecification` | kinesis_stream_specification | `map` | optional, computed, provider-chosen |  | The Kinesis Data Streams configuration for the specified table. |
| `LocalSecondaryIndexes` | local_secondary_indexes | `list` | optional, computed, provider-chosen |  | Local secondary indexes to be created on the table. You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes. |
| `OnDemandThroughput` | on_demand_throughput | `map` | optional, computed, provider-chosen |  | Sets the maximum number of read and write units for the specified on-demand table. If you use this property, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both. |
| `PointInTimeRecoverySpecification` | point_in_time_recovery_specification | `map` | optional, computed, provider-chosen |  | The settings used to enable point in time recovery. |
| `ProvisionedThroughput` | provisioned_throughput | `map` | optional, computed, provider-chosen |  | Throughput for the specified table, which consists of values for ``ReadCapacityUnits`` and ``WriteCapacityUnits``. For more information about the contents of a provisioned throughput structure, see [Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html). |
| `ResourcePolicy` | resource_policy | `map` | optional, computed, provider-chosen |  | Creates or updates a resource-based policy document that contains the permissions for DDB resources, such as a table, its indexes, and stream. Resource-based policies let you define access permissions by specifying who has access to each resource, and the actions they are allowed to perform on each resource. |
| `SSESpecification` | sse_specification | `map` | optional, computed, provider-chosen |  | Represents the settings used to enable server-side encryption. |
| `StreamArn` | stream_arn | `string` | computed |  |  |
| `StreamSpecification` | stream_specification | `map` | optional, computed, provider-chosen |  | Represents the DynamoDB Streams configuration for a table in DynamoDB. |
| `TableClass` | table_class | `string` | optional, computed, provider-chosen |  | The table class of the new table. Valid values are ``STANDARD`` and ``STANDARD_INFREQUENT_ACCESS``. |
| `TableName` | table_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the table. If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TimeToLiveSpecification` | time_to_live_specification | `map` | optional, computed, provider-chosen |  | Represents the settings used to enable or disable Time to Live (TTL) for the specified table. |
| `WarmThroughput` | warm_throughput | `map` | optional, computed, provider-chosen |  | Provides visibility into the number of read and write operations your table or secondary index can instantaneously support. The settings can be modified using the ``UpdateTable`` operation to meet the throughput requirements of an upcoming peak event. |

Supports update: yes

Discovery: supported
