# aws.datamigration

**CloudFormation type:** `AWS::DMS::DataMigration`

Resource schema for AWS::DMS::DataMigration.

Region attribute: `region`

**Import ID:** `<region>/DataMigrationArn` (AWS::DMS::DataMigration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataMigrationArn` | data_migration_arn | `string` | computed |  | The property describes an ARN of the data migration. |
| `DataMigrationCreateTime` | data_migration_create_time | `string` | computed |  | The property describes the create time of the data migration. |
| `DataMigrationIdentifier` | data_migration_identifier | `string` | optional, computed, provider-chosen, write-only |  | The property describes an ARN of the data migration. |
| `DataMigrationName` | data_migration_name | `string` | optional, computed, provider-chosen |  | The property describes a name to identify the data migration. |
| `DataMigrationSettings` | data_migration_settings | `map` | optional, computed, provider-chosen |  | The property describes the settings for the data migration. |
| `DataMigrationType` | data_migration_type | `string` | required |  | The property describes the type of migration. |
| `MigrationProjectIdentifier` | migration_project_identifier | `string` | required |  | The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn |
| `ServiceAccessRoleArn` | service_access_role_arn | `string` | required | aws.role.Arn | The property describes Amazon Resource Name (ARN) of the service access role. |
| `SourceDataSettings` | source_data_settings | `list` | optional, computed, provider-chosen |  | The property describes the settings for the data migration. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
