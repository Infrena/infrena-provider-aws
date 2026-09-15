# aws.ipamresourcediscovery

**CloudFormation type:** `AWS::EC2::IPAMResourceDiscovery`

Resource Schema of AWS::EC2::IPAMResourceDiscovery Type

Region attribute: `region`

**Import ID:** `<region>/IpamResourceDiscoveryId` (AWS::EC2::IPAMResourceDiscovery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IpamResourceDiscoveryArn` | ipam_resource_discovery_arn | `string` | computed |  | Amazon Resource Name (Arn) for the Resource Discovery. |
| `IpamResourceDiscoveryId` | ipam_resource_discovery_id | `string` | computed |  | Id of the IPAM Pool. |
| `IpamResourceDiscoveryRegion` | ipam_resource_discovery_region | `string` | computed |  | The region the resource discovery is setup in. |
| `IsDefault` | is_default | `boolean` | computed |  | Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6. |
| `OperatingRegions` | operating_regions | `list` | optional, computed, provider-chosen |  | The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring |
| `OrganizationalUnitExclusions` | organizational_unit_exclusions | `list` | optional, computed, provider-chosen |  | A set of organizational unit (OU) exclusions for this resource. |
| `OwnerId` | owner_id | `string` | computed |  | Owner Account ID of the Resource Discovery |
| `State` |  | `string` | computed |  | The state of this Resource Discovery. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
