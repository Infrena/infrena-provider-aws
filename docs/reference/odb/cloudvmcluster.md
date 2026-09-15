# aws.cloudvmcluster

**CloudFormation type:** `AWS::ODB::CloudVmCluster`

The AWS::ODB::CloudVmCluster resource creates a Cloud VM Cluster

Region attribute: `region`

**Import ID:** `<region>/CloudVmClusterArn` (AWS::ODB::CloudVmCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudExadataInfrastructureId` | cloud_exadata_infrastructure_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cloudexadatainfrastructure.CloudExadataInfrastructureId | The unique identifier of the Exadata infrastructure that this VM cluster belongs to. |
| `CloudVmClusterArn` | cloud_vm_cluster_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the VM cluster. |
| `CloudVmClusterId` | cloud_vm_cluster_id | `string` | computed |  | The unique identifier of the VM cluster. |
| `ClusterName` | cluster_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Grid Infrastructure (GI) cluster. |
| `ComputeModel` | compute_model | `string` | computed |  | The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled. |
| `CpuCoreCount` | cpu_core_count | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of CPU cores enabled on the VM cluster. |
| `DataCollectionOptions` | data_collection_options | `map` | optional, computed, provider-chosen, replaces on change |  | Information about the data collection options enabled for a VM cluster. |
| `DataStorageSizeInTBs` | data_storage_size_in_t_bs | `float` | optional, computed, provider-chosen, replaces on change |  | The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster. |
| `DbNodeStorageSizeInGBs` | db_node_storage_size_in_g_bs | `integer` | optional, computed, provider-chosen, replaces on change |  | The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster. |
| `DbNodes` | db_nodes | `list` | optional, computed, provider-chosen |  | The DB nodes that are implicitly created and managed as part of this VM Cluster. |
| `DbServers` | db_servers | `list` | optional, computed, provider-chosen, replaces on change |  | The list of database servers for the VM cluster. |
| `DiskRedundancy` | disk_redundancy | `string` | computed |  | The type of redundancy configured for the VM cluster. NORMAL is 2-way redundancy. HIGH is 3-way redundancy. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen, replaces on change |  | The user-friendly name for the VM cluster. |
| `Domain` |  | `string` | computed |  | The domain of the VM cluster. |
| `GiVersion` | gi_version | `string` | optional, computed, provider-chosen, replaces on change |  | The software version of the Oracle Grid Infrastructure (GI) for the VM cluster. |
| `Hostname` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The host name for the VM cluster. |
| `IamRoles` | iam_roles | `list` | optional, computed, provider-chosen |  | The AWS Identity and Access Management (IAM) service roles associated with the VM cluster. |
| `IsLocalBackupEnabled` | is_local_backup_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether database backups to local Exadata storage is enabled for the VM cluster. |
| `IsSparseDiskgroupEnabled` | is_sparse_diskgroup_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the VM cluster is configured with a sparse disk group. |
| `LicenseModel` | license_model | `string` | optional, computed, provider-chosen, replaces on change |  | The Oracle license model applied to the VM cluster. |
| `ListenerPort` | listener_port | `integer` | computed |  | The port number configured for the listener on the VM cluster. |
| `MemorySizeInGBs` | memory_size_in_g_bs | `integer` | optional, computed, provider-chosen, replaces on change |  | The amount of memory, in gigabytes (GB), that's allocated for the VM cluster. |
| `NodeCount` | node_count | `integer` | computed |  | The number of nodes in the VM cluster. |
| `OciResourceAnchorName` | oci_resource_anchor_name | `string` | computed |  | The name of the OCI resource anchor for the VM cluster. |
| `OciUrl` | oci_url | `string` | computed |  | The HTTPS link to the VM cluster in OCI. |
| `Ocid` |  | `string` | computed |  | The OCID of the VM cluster. |
| `OdbNetworkId` | odb_network_id | `string` | optional, computed, provider-chosen, replaces on change | aws.odbnetwork.OdbNetworkId | The unique identifier of the ODB network for the VM cluster. |
| `ScanDnsName` | scan_dns_name | `string` | computed |  | The FQDN of the DNS record for the Single Client Access Name (SCAN) IP addresses that are associated with the VM cluster. |
| `ScanIpIds` | scan_ip_ids | `list` | computed |  | The OCID of the SCAN IP addresses that are associated with the VM cluster. |
| `ScanListenerPortTcp` | scan_listener_port_tcp | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | Property description not available. |
| `Shape` |  | `string` | computed |  | The hardware model name of the Exadata infrastructure that's running the VM cluster. |
| `SshPublicKeys` | ssh_public_keys | `list` | optional, computed, provider-chosen, replaces on change |  | The public key portion of one or more key pairs used for SSH access to the VM cluster. |
| `StorageSizeInGBs` | storage_size_in_g_bs | `integer` | computed |  | The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster. |
| `SystemVersion` | system_version | `string` | optional, computed, provider-chosen, replaces on change |  | The operating system version of the image chosen for the VM cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Vm Cluster. |
| `TimeZone` | time_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The time zone of the VM cluster. |
| `VipIds` | vip_ids | `list` | computed |  | The virtual IP (VIP) addresses that are associated with the VM cluster. Oracle's Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the VM cluster to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster. |

Supports update: yes

Discovery: supported
