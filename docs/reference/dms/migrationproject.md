# aws.migrationproject

**CloudFormation type:** `AWS::DMS::MigrationProject`

Resource schema for AWS::DMS::MigrationProject

Region attribute: `region`

**Import ID:** `<region>/MigrationProjectArn` (AWS::DMS::MigrationProject)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The optional description of the migration project. |
| `InstanceProfileArn` | instance_profile_arn | `string` | optional, computed, provider-chosen | aws.dms.instanceprofile.InstanceProfileArn | The property describes an instance profile arn for the migration project. For read |
| `InstanceProfileIdentifier` | instance_profile_identifier | `string` | optional, computed, provider-chosen, write-only |  | The property describes an instance profile identifier for the migration project. For create |
| `InstanceProfileName` | instance_profile_name | `string` | optional, computed, provider-chosen |  | The property describes an instance profile name for the migration project. For read |
| `MigrationProjectArn` | migration_project_arn | `string` | computed |  | The property describes an ARN of the migration project. |
| `MigrationProjectCreationTime` | migration_project_creation_time | `string` | optional, computed, provider-chosen |  | The property describes a creating time of the migration project. |
| `MigrationProjectIdentifier` | migration_project_identifier | `string` | optional, computed, provider-chosen, write-only |  | The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn |
| `MigrationProjectName` | migration_project_name | `string` | optional, computed, provider-chosen |  | The property describes a name to identify the migration project. |
| `SchemaConversionApplicationAttributes` | schema_conversion_application_attributes | `map` | optional, computed, provider-chosen |  | The property describes schema conversion application attributes for the migration project. |
| `SourceDataProviderDescriptors` | source_data_provider_descriptors | `list` | optional, computed, provider-chosen |  | The property describes source data provider descriptors for the migration project. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetDataProviderDescriptors` | target_data_provider_descriptors | `list` | optional, computed, provider-chosen |  | The property describes target data provider descriptors for the migration project. |
| `TransformationRules` | transformation_rules | `string` | optional, computed, provider-chosen |  | The property describes transformation rules for the migration project. |

Supports update: yes

Discovery: supported
