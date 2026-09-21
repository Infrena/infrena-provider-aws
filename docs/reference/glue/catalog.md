# aws.catalog

**CloudFormation type:** `AWS::Glue::Catalog`

Creates a catalog in the Glue Data Catalog.

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::Glue::Catalog)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowFullTableExternalDataAccess` | allow_full_table_external_data_access | `string` | optional, computed, provider-chosen |  | Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation. |
| `CatalogId` | catalog_id | `string` | computed |  | The ID of the catalog. |
| `CatalogProperties` | catalog_properties | `map` | optional, computed, provider-chosen |  | A structure that specifies data lake access properties and other custom properties. |
| `CreateDatabaseDefaultPermissions` | create_database_default_permissions | `list` | optional, computed, provider-chosen |  | An array of PrincipalPermissions objects for default database permissions. |
| `CreateTableDefaultPermissions` | create_table_default_permissions | `list` | optional, computed, provider-chosen |  | An array of PrincipalPermissions objects for default table permissions. |
| `CreateTime` | create_time | `integer` | computed |  | The time at which the catalog was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the catalog. |
| `FederatedCatalog` | federated_catalog | `map` | optional, computed, provider-chosen |  | A FederatedCatalog structure that references an entity outside the Glue Data Catalog. |
| `Name` |  | `string` | required, replaces on change |  | The name of the catalog to create. |
| `OverwriteChildResourcePermissionsWithDefault` | overwrite_child_resource_permissions_with_default | `string` | optional, computed, provider-chosen, write-only |  | Specifies whether to overwrite child resource permissions with the default permissions. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | A map of key-value pairs that define parameters and properties of the catalog. |
| `ResourceArn` | resource_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the catalog. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetRedshiftCatalog` | target_redshift_catalog | `map` | optional, computed, provider-chosen |  | A structure that describes a target catalog for resource linking. |
| `UpdateTime` | update_time | `integer` | computed |  | The time at which the catalog was last updated. |

Supports update: yes

Discovery: supported
