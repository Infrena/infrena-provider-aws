# aws.cloudexadatainfrastructure

**CloudFormation type:** `AWS::ODB::CloudExadataInfrastructure`

The AWS::ODB::CloudExadataInfrastructure resource creates an Exadata Infrastructure

Region attribute: `region`

**Import ID:** `<region>/CloudExadataInfrastructureArn` (AWS::ODB::CloudExadataInfrastructure)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActivatedStorageCount` | activated_storage_count | `integer` | computed |  | The number of storage servers requested for the Exadata infrastructure. |
| `AdditionalStorageCount` | additional_storage_count | `integer` | computed |  | The number of storage servers requested for the Exadata infrastructure. |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Availability Zone (AZ) where the Exadata infrastructure is located. |
| `AvailabilityZoneId` | availability_zone_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AZ ID of the AZ where the Exadata infrastructure is located. |
| `AvailableStorageSizeInGBs` | available_storage_size_in_g_bs | `integer` | computed |  | The amount of available storage, in gigabytes (GB), for the Exadata infrastructure. |
| `CloudExadataInfrastructureArn` | cloud_exadata_infrastructure_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the Exadata infrastructure. |
| `CloudExadataInfrastructureId` | cloud_exadata_infrastructure_id | `string` | computed |  | The unique identifier for the Exadata infrastructure. |
| `ComputeCount` | compute_count | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of database servers for the Exadata infrastructure. |
| `ComputeModel` | compute_model | `string` | computed |  | The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled. |
| `CpuCount` | cpu_count | `integer` | computed |  | The total number of CPU cores that are allocated to the Exadata infrastructure. |
| `CustomerContactsToSendToOCI` | customer_contacts_to_send_to_oci | `list` | optional, computed, provider-chosen, replaces on change |  | The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure. |
| `DataStorageSizeInTBs` | data_storage_size_in_t_bs | `float` | computed |  | The size of the Exadata infrastructure's data disk group, in terabytes (TB). |
| `DatabaseServerType` | database_server_type | `string` | optional, computed, provider-chosen, replaces on change |  | The database server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation. |
| `DbNodeStorageSizeInGBs` | db_node_storage_size_in_g_bs | `integer` | computed |  | The size of the Exadata infrastructure's local node storage, in gigabytes (GB). |
| `DbServerIds` | db_server_ids | `list` | computed |  | The list of database server identifiers for the Exadata infrastructure. |
| `DbServerVersion` | db_server_version | `string` | computed |  | The software version of the database servers (dom0) in the Exadata infrastructure. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen, replaces on change |  | The user-friendly name for the Exadata infrastructure. |
| `MaintenanceWindow` | maintenance_window | `map` | optional, computed, provider-chosen |  | The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window. |
| `MaxCpuCount` | max_cpu_count | `integer` | computed |  | The total number of CPU cores available on the Exadata infrastructure. |
| `MaxDataStorageInTBs` | max_data_storage_in_t_bs | `float` | computed |  | The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure. |
| `MaxDbNodeStorageSizeInGBs` | max_db_node_storage_size_in_g_bs | `integer` | computed |  | The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure. |
| `MaxMemoryInGBs` | max_memory_in_g_bs | `integer` | computed |  | The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure. |
| `MemorySizeInGBs` | memory_size_in_g_bs | `integer` | computed |  | The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure. |
| `OciResourceAnchorName` | oci_resource_anchor_name | `string` | computed |  | The name of the OCI resource anchor for the Exadata infrastructure. |
| `OciUrl` | oci_url | `string` | computed |  | The HTTPS link to the Exadata infrastructure in OCI. |
| `Ocid` |  | `string` | computed |  | The OCID of the Exadata infrastructure. |
| `Shape` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The model name of the Exadata infrastructure. |
| `StorageCount` | storage_count | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of storage servers that are activated for the Exadata infrastructure. |
| `StorageServerType` | storage_server_type | `string` | optional, computed, provider-chosen, replaces on change |  | The storage server model type of the Exadata infrastructure. For the list of valid model names, use the ListDbSystemShapes operation. |
| `StorageServerVersion` | storage_server_version | `string` | computed |  | The software version of the storage servers on the Exadata infrastructure. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Exadata Infrastructure. |
| `TotalStorageSizeInGBs` | total_storage_size_in_g_bs | `integer` | computed |  | The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure. |

Supports update: yes

Discovery: supported
