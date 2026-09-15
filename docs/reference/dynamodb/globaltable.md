# aws.globaltable

**CloudFormation type:** `AWS::DynamoDB::GlobalTable`

Version: None. Resource Type definition for AWS::DynamoDB::GlobalTable

Region attribute: `region`

**Import ID:** `<region>/TableName` (AWS::DynamoDB::GlobalTable)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AttributeDefinitions` | attribute_definitions | `list` | optional, computed, provider-chosen |  |  |
| `BillingMode` | billing_mode | `string` | optional, computed, provider-chosen |  |  |
| `GlobalSecondaryIndexes` | global_secondary_indexes | `list` | optional, computed, provider-chosen |  |  |
| `GlobalTableSourceArn` | global_table_source_arn | `string` | optional, computed, provider-chosen, write-only |  |  |
| `GlobalTableWitnesses` | global_table_witnesses | `list` | optional, computed, provider-chosen |  |  |
| `KeySchema` | key_schema | `list` | optional, computed, provider-chosen |  |  |
| `LocalSecondaryIndexes` | local_secondary_indexes | `list` | optional, computed, provider-chosen |  |  |
| `MultiRegionConsistency` | multi_region_consistency | `string` | optional, computed, provider-chosen |  |  |
| `ReadOnDemandThroughputSettings` | read_on_demand_throughput_settings | `map` | optional, computed, provider-chosen |  |  |
| `ReadProvisionedThroughputSettings` | read_provisioned_throughput_settings | `map` | optional, computed, provider-chosen |  |  |
| `Replicas` |  | `list` | required |  |  |
| `SSESpecification` | sse_specification | `map` | optional, computed, provider-chosen |  |  |
| `StreamArn` | stream_arn | `string` | computed |  |  |
| `StreamSpecification` | stream_specification | `map` | optional, computed, provider-chosen |  |  |
| `TableId` | table_id | `string` | computed |  |  |
| `TableName` | table_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `TimeToLiveSpecification` | time_to_live_specification | `map` | optional, computed, provider-chosen |  |  |
| `WarmThroughput` | warm_throughput | `map` | optional, computed, provider-chosen |  |  |
| `WriteOnDemandThroughputSettings` | write_on_demand_throughput_settings | `map` | optional, computed, provider-chosen |  |  |
| `WriteProvisionedThroughputSettings` | write_provisioned_throughput_settings | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
