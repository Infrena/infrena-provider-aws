# aws.ipamprefixlistresolvertarget

**CloudFormation type:** `AWS::EC2::IPAMPrefixListResolverTarget`

Resource Type definition for AWS::EC2::IPAMPrefixListResolverTarget

Region attribute: `region`

**Import ID:** `<region>/IpamPrefixListResolverTargetId` (AWS::EC2::IPAMPrefixListResolverTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DesiredVersion` | desired_version | `integer` | optional, computed, provider-chosen |  | The desired version of the Prefix List Resolver that this Target should synchronize with. |
| `IpamPrefixListResolverId` | ipam_prefix_list_resolver_id | `string` | required, replaces on change | aws.ipamprefixlistresolver.IpamPrefixListResolverId | The Id of the IPAM Prefix List Resolver associated with this Target. |
| `IpamPrefixListResolverTargetArn` | ipam_prefix_list_resolver_target_arn | `string` | computed |  | Id of the IPAM Prefix List Resolver Target. |
| `IpamPrefixListResolverTargetId` | ipam_prefix_list_resolver_target_id | `string` | computed |  | Id of the IPAM Prefix List Resolver Target. |
| `PrefixListId` | prefix_list_id | `string` | required, replaces on change | aws.prefixlist.PrefixListId | The Id of the Managed Prefix List. |
| `PrefixListRegion` | prefix_list_region | `string` | required, replaces on change |  | The region that the Managed Prefix List is located in. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TrackLatestVersion` | track_latest_version | `boolean` | required |  | Indicates whether this Target automatically tracks the latest version of the Prefix List Resolver. |

Supports update: yes

Discovery: supported
