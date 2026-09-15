# aws.ipamprefixlistresolver

**CloudFormation type:** `AWS::EC2::IPAMPrefixListResolver`

Resource Type definition for AWS::EC2::IPAMPrefixListResolver

Region attribute: `region`

**Import ID:** `<region>/IpamPrefixListResolverId` (AWS::EC2::IPAMPrefixListResolver)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddressFamily` | address_family | `string` | required, replaces on change |  | The address family of the address space in this Prefix List Resolver. Either IPv4 or IPv6. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `IpamArn` | ipam_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM this Prefix List Resolver is a part of. |
| `IpamId` | ipam_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.ipam.IpamId | The Id of the IPAM this Prefix List Resolver is a part of. |
| `IpamPrefixListResolverArn` | ipam_prefix_list_resolver_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM Prefix List Resolver |
| `IpamPrefixListResolverId` | ipam_prefix_list_resolver_id | `string` | computed |  | Id of the IPAM Prefix List Resolver. |
| `Rules` |  | `list` | optional, computed, provider-chosen |  | Rules define the business logic for selecting CIDRs from IPAM. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
