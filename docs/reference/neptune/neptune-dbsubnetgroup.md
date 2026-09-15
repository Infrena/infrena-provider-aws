# aws.neptune.dbsubnetgroup

**CloudFormation type:** `AWS::Neptune::DBSubnetGroup`

The AWS::Neptune::DBSubnetGroup type creates an Amazon Neptune DB subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same AWS Region.

Region attribute: `region`

**Import ID:** `<region>/DBSubnetGroupName` (AWS::Neptune::DBSubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBSubnetGroupDescription` | db_subnet_group_description | `string` | required |  | The description for the DB subnet group. |
| `DBSubnetGroupName` | db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the DB subnet group. This value is stored as a lowercase string. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The Amazon EC2 subnet IDs for the DB subnet group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An optional array of key-value pairs to apply to this DB subnet group. |

Supports update: yes

Discovery: supported
