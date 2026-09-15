# aws.cloudautonomousvmcluster

**CloudFormation type:** `AWS::ODB::CloudAutonomousVmCluster`

The AWS::ODB::CloudAutonomousVmCluster resource creates a Cloud Autonomous VM Cluster

Region attribute: `region`

**Import ID:** `<region>/CloudAutonomousVmClusterArn` (AWS::ODB::CloudAutonomousVmCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutonomousDataStoragePercentage` | autonomous_data_storage_percentage | `float` | computed |  | The percentage of data storage currently in use for Autonomous Databases in the Autonomous VM cluster. |
| `AutonomousDataStorageSizeInTBs` | autonomous_data_storage_size_in_t_bs | `float` | optional, computed, provider-chosen, replaces on change |  | The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB. |
| `AvailableAutonomousDataStorageSizeInTBs` | available_autonomous_data_storage_size_in_t_bs | `float` | computed |  | The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB. |
| `AvailableContainerDatabases` | available_container_databases | `integer` | computed |  | The number of Autonomous CDBs that you can create with the currently available storage. |
| `AvailableCpus` | available_cpus | `float` | computed |  | The number of CPU cores available for allocation to Autonomous Databases. |
| `CloudAutonomousVmClusterArn` | cloud_autonomous_vm_cluster_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the Autonomous VM cluster. |
| `CloudAutonomousVmClusterId` | cloud_autonomous_vm_cluster_id | `string` | computed |  | The unique identifier of the Autonomous VM cluster. |
| `CloudExadataInfrastructureId` | cloud_exadata_infrastructure_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cloudexadatainfrastructure.CloudExadataInfrastructureId | The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster. |
| `ComputeModel` | compute_model | `string` | computed |  | The compute model of the Autonomous VM cluster: ECPU or OCPU. |
| `CpuCoreCount` | cpu_core_count | `integer` | computed |  | The total number of CPU cores in the Autonomous VM cluster. |
| `CpuCoreCountPerNode` | cpu_core_count_per_node | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of CPU cores enabled per node in the Autonomous VM cluster. |
| `CpuPercentage` | cpu_percentage | `float` | computed |  | The percentage of total CPU cores currently in use in the Autonomous VM cluster. |
| `DataStorageSizeInGBs` | data_storage_size_in_g_bs | `float` | computed |  | The total data storage allocated to the Autonomous VM cluster, in GB. |
| `DataStorageSizeInTBs` | data_storage_size_in_t_bs | `float` | computed |  | The total data storage allocated to the Autonomous VM cluster, in TB. |
| `DbNodeStorageSizeInGBs` | db_node_storage_size_in_g_bs | `integer` | computed |  | The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB). |
| `DbServers` | db_servers | `list` | optional, computed, provider-chosen, replaces on change |  | The list of database servers associated with the Autonomous VM cluster. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The user-provided description of the Autonomous VM cluster. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen, replaces on change |  | The display name of the Autonomous VM cluster. |
| `Domain` |  | `string` | computed |  | The domain name for the Autonomous VM cluster. |
| `ExadataStorageInTBsLowestScaledValue` | exadata_storage_in_t_bs_lowest_scaled_value | `float` | computed |  | The minimum value to which you can scale down the Exadata storage, in TB. |
| `Hostname` |  | `string` | computed |  | The hostname for the Autonomous VM cluster. |
| `IamRoles` | iam_roles | `list` | optional, computed, provider-chosen |  | The AWS Identity and Access Management (IAM) service roles associated with the Autonomous VM cluster. |
| `IsMtlsEnabledVmCluster` | is_mtls_enabled_vm_cluster | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster. |
| `LicenseModel` | license_model | `string` | optional, computed, provider-chosen, replaces on change |  | The Oracle license model that applies to the Autonomous VM cluster. Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE. |
| `MaintenanceWindow` | maintenance_window | `map` | optional, computed, provider-chosen, replaces on change |  | The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window. |
| `MaxAcdsLowestScaledValue` | max_acds_lowest_scaled_value | `integer` | computed |  | The minimum value to which you can scale down the maximum number of Autonomous CDBs. |
| `MemoryPerOracleComputeUnitInGBs` | memory_per_oracle_compute_unit_in_g_bs | `integer` | optional, computed, provider-chosen, replaces on change |  | The amount of memory allocated per Oracle Compute Unit, in GB. |
| `MemorySizeInGBs` | memory_size_in_g_bs | `integer` | computed |  | The total amount of memory allocated to the Autonomous VM cluster, in gigabytes (GB). |
| `NodeCount` | node_count | `integer` | computed |  | The number of database server nodes in the Autonomous VM cluster. |
| `NonProvisionableAutonomousContainerDatabases` | non_provisionable_autonomous_container_databases | `integer` | computed |  | The number of Autonomous CDBs that can't be provisioned because of resource constraints. |
| `OciResourceAnchorName` | oci_resource_anchor_name | `string` | computed |  | The name of the OCI resource anchor associated with this Autonomous VM cluster. |
| `OciUrl` | oci_url | `string` | computed |  | The URL for accessing the OCI console page for this Autonomous VM cluster. |
| `Ocid` |  | `string` | computed |  | The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster. |
| `OdbNetworkId` | odb_network_id | `string` | optional, computed, provider-chosen, replaces on change | aws.odbnetwork.OdbNetworkId | The unique identifier of the ODB network associated with this Autonomous VM cluster. |
| `ProvisionableAutonomousContainerDatabases` | provisionable_autonomous_container_databases | `integer` | computed |  | The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster. |
| `ProvisionedAutonomousContainerDatabases` | provisioned_autonomous_container_databases | `integer` | computed |  | The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster. |
| `ProvisionedCpus` | provisioned_cpus | `float` | computed |  | The number of CPU cores currently provisioned in the Autonomous VM cluster. |
| `ReclaimableCpus` | reclaimable_cpus | `float` | computed |  | The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases. |
| `ReservedCpus` | reserved_cpus | `float` | computed |  | The number of CPU cores reserved for system operations and redundancy. |
| `ScanListenerPortNonTls` | scan_listener_port_non_tls | `integer` | optional, computed, provider-chosen, replaces on change |  | The SCAN listener port for non-TLS (TCP) protocol. The default is 1521. |
| `ScanListenerPortTls` | scan_listener_port_tls | `integer` | optional, computed, provider-chosen, replaces on change |  | The SCAN listener port for TLS (TCP) protocol. The default is 2484. |
| `Shape` |  | `string` | computed |  | The shape of the Exadata infrastructure for the Autonomous VM cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the Autonomous VM cluster. |
| `TimeZone` | time_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The time zone of the Autonomous VM cluster. |
| `TotalContainerDatabases` | total_container_databases | `integer` | optional, computed, provider-chosen, replaces on change |  | The total number of Autonomous Container Databases that can be created with the allocated local storage. |

Supports update: yes

Discovery: supported
