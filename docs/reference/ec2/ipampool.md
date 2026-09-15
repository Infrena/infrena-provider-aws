# aws.ipampool

**CloudFormation type:** `AWS::EC2::IPAMPool`

Resource Schema of AWS::EC2::IPAMPool Type

Region attribute: `region`

**Import ID:** `<region>/IpamPoolId` (AWS::EC2::IPAMPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddressFamily` | address_family | `string` | required, replaces on change |  | The address family of the address space in this pool. Either IPv4 or IPv6. |
| `AllocationDefaultNetmaskLength` | allocation_default_netmask_length | `integer` | optional, computed, provider-chosen |  | The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified. |
| `AllocationMaxNetmaskLength` | allocation_max_netmask_length | `integer` | optional, computed, provider-chosen |  | The maximum allowed netmask length for allocations made from this pool. |
| `AllocationMinNetmaskLength` | allocation_min_netmask_length | `integer` | optional, computed, provider-chosen |  | The minimum allowed netmask length for allocations made from this pool. |
| `AllocationResourceTags` | allocation_resource_tags | `list` | optional, computed, provider-chosen |  | When specified, an allocation will not be allowed unless a resource has a matching set of tags. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM Pool. |
| `AutoImport` | auto_import | `boolean` | optional, computed, provider-chosen |  | Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically. |
| `AwsService` | aws_service | `string` | optional, computed, provider-chosen, replaces on change |  | Limits which service in Amazon Web Services that the pool can be used in. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IpamArn` | ipam_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM this pool is a part of. |
| `IpamPoolId` | ipam_pool_id | `string` | computed |  | Id of the IPAM Pool. |
| `IpamScopeArn` | ipam_scope_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the scope this pool is a part of. |
| `IpamScopeId` | ipam_scope_id | `string` | required, replaces on change | aws.ipamscope.IpamScopeId | The Id of the scope this pool is a part of. |
| `IpamScopeType` | ipam_scope_type | `string` | computed |  | Determines whether this scope contains publicly routable space or space for a private network |
| `Locale` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The region of this pool. If not set, this will default to "None" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match. |
| `PoolDepth` | pool_depth | `integer` | computed |  | The depth of this pool in the source pool hierarchy. |
| `ProvisionedCidrs` | provisioned_cidrs | `list` | optional, computed, provider-chosen |  | A list of cidrs representing the address space available for allocation in this pool. |
| `PublicIpSource` | public_ip_source | `string` | optional, computed, provider-chosen, replaces on change |  | The IP address source for pools in the public scope. Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`. |
| `PubliclyAdvertisable` | publicly_advertisable | `boolean` | optional, computed, provider-chosen, replaces on change |  | Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6. |
| `SourceIpamPoolId` | source_ipam_pool_id | `string` | optional, computed, provider-chosen, replaces on change | aws.ipampool.IpamPoolId | The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool. |
| `SourceResource` | source_resource | `map` | optional, computed, provider-chosen, replaces on change |  | The resource associated with this pool's space. Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations |
| `State` |  | `string` | computed |  | The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete" |
| `StateMessage` | state_message | `string` | computed |  | An explanation of how the pool arrived at it current state. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
