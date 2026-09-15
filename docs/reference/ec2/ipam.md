# aws.ipam

**CloudFormation type:** `AWS::EC2::IPAM`

Resource Schema of AWS::EC2::IPAM Type

Region attribute: `region`

**Import ID:** `<region>/IpamId` (AWS::EC2::IPAM)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM. |
| `DefaultResourceDiscoveryAssociationId` | default_resource_discovery_association_id | `string` | computed |  | The Id of the default association to the default resource discovery, created with this IPAM. |
| `DefaultResourceDiscoveryId` | default_resource_discovery_id | `string` | computed |  | The Id of the default resource discovery, created with this IPAM. |
| `DefaultResourceDiscoveryOrganizationalUnitExclusions` | default_resource_discovery_organizational_unit_exclusions | `list` | optional, computed, provider-chosen |  | A set of organizational unit (OU) exclusions for the default resource discovery, created with this IPAM. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnablePrivateGua` | enable_private_gua | `boolean` | optional, computed, provider-chosen |  | Enable provisioning of GUA space in private pools. |
| `IpamId` | ipam_id | `string` | computed |  | Id of the IPAM. |
| `MeteredAccount` | metered_account | `string` | optional, computed, provider-chosen |  | A metered account is an account that is charged for active IP addresses managed in IPAM |
| `OperatingRegions` | operating_regions | `list` | optional, computed, provider-chosen |  | The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring |
| `PrivateDefaultScopeId` | private_default_scope_id | `string` | computed |  | The Id of the default scope for publicly routable IP space, created with this IPAM. |
| `PublicDefaultScopeId` | public_default_scope_id | `string` | computed |  | The Id of the default scope for publicly routable IP space, created with this IPAM. |
| `ResourceDiscoveryAssociationCount` | resource_discovery_association_count | `integer` | computed |  | The count of resource discoveries associated with this IPAM. |
| `ScopeCount` | scope_count | `integer` | computed |  | The number of scopes that currently exist in this IPAM. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `Tier` |  | `string` | optional, computed, provider-chosen |  | The tier of the IPAM. |

Supports update: yes

Discovery: supported
