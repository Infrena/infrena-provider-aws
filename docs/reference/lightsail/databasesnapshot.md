# aws.databasesnapshot

**CloudFormation type:** `AWS::Lightsail::DatabaseSnapshot`

Resource Type definition for AWS::Lightsail::DatabaseSnapshot

Region attribute: `region`

**Import ID:** `<region>/RelationalDatabaseSnapshotName` (AWS::Lightsail::DatabaseSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the database snapshot. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the database snapshot was created. |
| `Engine` |  | `string` | computed |  | The software of the database snapshot (for example, MySQL). |
| `EngineVersion` | engine_version | `string` | computed |  | The database engine version for the database snapshot (for example, 5.7.23). |
| `FromRelationalDatabaseArn` | from_relational_database_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the database from which the database snapshot was created. |
| `FromRelationalDatabaseBlueprintId` | from_relational_database_blueprint_id | `string` | computed |  | The blueprint ID of the database from which the database snapshot was created. A blueprint describes the major engine version of a database. |
| `FromRelationalDatabaseBundleId` | from_relational_database_bundle_id | `string` | computed |  | The bundle ID of the database from which the database snapshot was created. |
| `FromRelationalDatabaseName` | from_relational_database_name | `string` | computed |  | The name of the source database from which the database snapshot was created. |
| `Location` |  | `map` | computed |  | The Region name and Availability Zone where the database snapshot is located. |
| `Name` |  | `string` | computed |  | The name of the database snapshot. |
| `RelationalDatabaseName` | relational_database_name | `string` | required, replaces on change |  | The name of the database on which to base your new snapshot. |
| `RelationalDatabaseSnapshotName` | relational_database_snapshot_name | `string` | required, replaces on change |  | The name for your new database snapshot. |
| `ResourceType` | resource_type | `string` | computed |  | The Lightsail resource type. |
| `SizeInGb` | size_in_gb | `integer` | computed |  | The size of the disk in GB (for example, 32) for the database snapshot. |
| `State` |  | `string` | computed |  | The state of the database snapshot. |
| `SupportCode` | support_code | `string` | computed |  | The support code for the database snapshot. Include this code in your email to support when you have questions about a database snapshot in Lightsail. This code enables our support team to look up your Lightsail information more easily. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
