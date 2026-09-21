# aws.memorydb.subnetgroup

**CloudFormation type:** `AWS::MemoryDB::SubnetGroup`

The AWS::MemoryDB::SubnetGroup resource creates an Amazon MemoryDB Subnet Group.

Region attribute: `region`

**Import ID:** `<region>/SubnetGroupName` (AWS::MemoryDB::SubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARN` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the subnet group. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description of the subnet group. |
| `SubnetGroupName` | subnet_group_name | `string` | required, replaces on change |  | The name of the subnet group. This value must be unique as it also serves as the subnet group identifier. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | A list of VPC subnet IDs for the subnet group. |
| `SupportedNetworkTypes` | supported_network_types | `list` | computed |  | Supported network types would be a list of network types supported by subnet group and can be either [ipv4] or [ipv4, dual_stack] or [ipv6]. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this subnet group. |

Supports update: yes

Discovery: supported
