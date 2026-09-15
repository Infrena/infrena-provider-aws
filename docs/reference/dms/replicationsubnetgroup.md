# aws.replicationsubnetgroup

**CloudFormation type:** `AWS::DMS::ReplicationSubnetGroup`

Resource Type definition for AWS::DMS::ReplicationSubnetGroup

Region attribute: `region`

**Import ID:** `<region>/ReplicationSubnetGroupIdentifier` (AWS::DMS::ReplicationSubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ReplicationSubnetGroupDescription` | replication_subnet_group_description | `string` | required |  | The description for the subnet group. |
| `ReplicationSubnetGroupIdentifier` | replication_subnet_group_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the replication subnet group. This value is stored as a lowercase string. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | One or more subnet IDs to be assigned to the replication subnet group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags to be assigned to the replication subnet group |

Supports update: yes

Discovery: supported
