# aws.ipamscope

**CloudFormation type:** `AWS::EC2::IPAMScope`

Resource Schema of AWS::EC2::IPAMScope Type

Region attribute: `region`

**Import ID:** `<region>/IpamScopeId` (AWS::EC2::IPAMScope)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM scope. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExternalAuthorityConfiguration` | external_authority_configuration | `map` | optional, computed, provider-chosen |  | External service configuration to connect your AWS IPAM scope. |
| `IpamArn` | ipam_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the IPAM this scope is a part of. |
| `IpamId` | ipam_id | `string` | required, replaces on change | aws.ipam.IpamId | The Id of the IPAM this scope is a part of. |
| `IpamScopeId` | ipam_scope_id | `string` | computed |  | Id of the IPAM scope. |
| `IpamScopeType` | ipam_scope_type | `string` | computed |  | Determines whether this scope contains publicly routable space or space for a private network |
| `IsDefault` | is_default | `boolean` | computed |  | Is this one of the default scopes created with the IPAM. |
| `PoolCount` | pool_count | `integer` | computed |  | The number of pools that currently exist in this scope. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
