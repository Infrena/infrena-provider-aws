# aws.odbnetwork

**CloudFormation type:** `AWS::ODB::OdbNetwork`

The AWS::ODB::OdbNetwork resource creates an ODB Network

Region attribute: `region`

**Import ID:** `<region>/OdbNetworkArn` (AWS::ODB::OdbNetwork)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Availability Zone (AZ) where the ODB network is located. |
| `AvailabilityZoneId` | availability_zone_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AZ ID of the AZ where the ODB network is located. |
| `BackupSubnetCidr` | backup_subnet_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The CIDR range of the backup subnet in the ODB network. |
| `ClientSubnetCidr` | client_subnet_cidr | `string` | optional, computed, provider-chosen, replaces on change |  | The CIDR range of the client subnet in the ODB network. |
| `CrossRegionS3RestoreSources` | cross_region_s3_restore_sources | `list` | optional, computed, provider-chosen, write-only |  | The cross-Region Amazon S3 restore sources for the ODB network. |
| `CustomDomainName` | custom_domain_name | `string` | optional, computed, provider-chosen, replaces on change |  | The domain name to use for the resources in the ODB network. |
| `DefaultDnsPrefix` | default_dns_prefix | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com. |
| `DeleteAssociatedResources` | delete_associated_resources | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to delete associated OCI networking resources along with the ODB network. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The user-friendly name of the ODB network. |
| `Ec2PlacementGroupIds` | ec2_placement_group_ids | `list` | computed |  | The list of EC2 placement group IDs associated with your ODB network. |
| `KmsAccess` | kms_access | `string` | optional, computed, provider-chosen, write-only |  | The AWS Key Management Service (KMS) access configuration for the ODB network. |
| `KmsPolicyDocument` | kms_policy_document | `string` | optional, computed, provider-chosen, write-only |  | The AWS Key Management Service (KMS) policy document that defines permissions for key usage within the ODB network. |
| `ManagedServices` | managed_services | `map` | computed |  | The managed services configuration for the ODB network. |
| `OciNetworkAnchorId` | oci_network_anchor_id | `string` | computed |  | The unique identifier of the OCI network anchor for the ODB network. |
| `OciResourceAnchorName` | oci_resource_anchor_name | `string` | computed |  | The name of the OCI resource anchor that's associated with the ODB network. |
| `OciVcnUrl` | oci_vcn_url | `string` | computed |  | The URL for the VCN that's associated with the ODB network. |
| `OdbNetworkArn` | odb_network_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the ODB network. |
| `OdbNetworkId` | odb_network_id | `string` | computed |  | The unique identifier of the ODB network. |
| `S3Access` | s3_access | `string` | optional, computed, provider-chosen, write-only |  | Specifies the configuration for Amazon S3 access from the ODB network. |
| `S3PolicyDocument` | s3_policy_document | `string` | optional, computed, provider-chosen, write-only |  | Specifies the endpoint policy for Amazon S3 access from the ODB network. |
| `StsAccess` | sts_access | `string` | optional, computed, provider-chosen, write-only |  | The AWS Security Token Service (STS) access configuration for the ODB network. |
| `StsPolicyDocument` | sts_policy_document | `string` | optional, computed, provider-chosen, write-only |  | The AWS Security Token Service (STS) policy document that defines permissions for token service usage within the ODB network. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the Odb Network. |
| `ZeroEtlAccess` | zero_etl_access | `string` | optional, computed, provider-chosen, write-only |  | Specifies the configuration for Zero-ETL access from the ODB network. |

Supports update: yes

Discovery: supported
