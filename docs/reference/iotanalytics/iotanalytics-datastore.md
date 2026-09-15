# aws.iotanalytics.datastore

**CloudFormation type:** `AWS::IoTAnalytics::Datastore`

Resource Type definition for AWS::IoTAnalytics::Datastore

Region attribute: `region`

**Import ID:** `<region>/DatastoreName` (AWS::IoTAnalytics::Datastore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatastoreName` | datastore_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DatastorePartitions` | datastore_partitions | `map` | optional, computed, provider-chosen |  |  |
| `DatastoreStorage` | datastore_storage | `map` | optional, computed, provider-chosen |  |  |
| `FileFormatConfiguration` | file_format_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `RetentionPeriod` | retention_period | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
