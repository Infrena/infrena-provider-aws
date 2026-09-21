# aws.rds.dbsubnetgroup

**CloudFormation type:** `AWS::RDS::DBSubnetGroup`

The ``AWS::RDS::DBSubnetGroup`` resource creates a database subnet group. Subnet groups must contain at least two subnets in two different Availability Zones in the same region. 

Region attribute: `region`

**Import ID:** `<region>/DBSubnetGroupName` (AWS::RDS::DBSubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBSubnetGroupArn` | db_subnet_group_arn | `string` | computed |  |  |
| `DBSubnetGroupDescription` | description, db_subnet_group_description | `string` | required |  | The description for the DB subnet group. |
| `DBSubnetGroupName` | name, db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the DB subnet group. This value is stored as a lowercase string. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The EC2 Subnet IDs for the DB subnet group. |
| `Tags` |  | `map` | tags map |  | Tags to assign to the DB subnet group. |

Supports update: yes

Discovery: supported
