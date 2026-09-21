# aws.ipamresourcediscoveryassociation

**CloudFormation type:** `AWS::EC2::IPAMResourceDiscoveryAssociation`

Resource Schema of AWS::EC2::IPAMResourceDiscoveryAssociation Type

Region attribute: `region`

**Import ID:** `<region>/IpamResourceDiscoveryAssociationId` (AWS::EC2::IPAMResourceDiscoveryAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IpamArn` | ipam_arn | `string` | computed |  | Arn of the IPAM. |
| `IpamId` | ipam_id | `string` | required, replaces on change | aws.ipam.IpamId | The Id of the IPAM this Resource Discovery is associated to. |
| `IpamRegion` | ipam_region | `string` | computed |  | The home region of the IPAM. |
| `IpamResourceDiscoveryAssociationArn` | ipam_resource_discovery_association_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the resource discovery association is a part of. |
| `IpamResourceDiscoveryAssociationId` | ipam_resource_discovery_association_id | `string` | computed |  | Id of the IPAM Resource Discovery Association. |
| `IpamResourceDiscoveryId` | ipam_resource_discovery_id | `string` | required, replaces on change | aws.ipamresourcediscovery.IpamResourceDiscoveryId | The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association. |
| `IsDefault` | is_default | `boolean` | computed |  | If the Resource Discovery Association exists due as part of CreateIpam. |
| `OwnerId` | owner_id | `string` | computed |  | The AWS Account ID for the account where the shared IPAM exists. |
| `ResourceDiscoveryStatus` | resource_discovery_status | `string` | computed |  | The status of the resource discovery. |
| `State` |  | `string` | computed |  | The operational state of the Resource Discovery Association. Related to Create/Delete activities. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
