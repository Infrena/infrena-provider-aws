# aws.docdb.dbsubnetgroup

**CloudFormation type:** `AWS::DocDB::DBSubnetGroup`

Resource Type definition for AWS::DocDB::DBSubnetGroup

Region attribute: `region`

**Import ID:** `<region>/DBSubnetGroupName` (AWS::DocDB::DBSubnetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBSubnetGroupDescription` | db_subnet_group_description | `string` | required |  | The description for the subnet group. |
| `DBSubnetGroupName` | db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the db subnet group. This value is stored as a lowercase string. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | One or more subnet IDs to be assigned to the db subnet group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags to be assigned to the db subnet group |

Supports update: yes

Discovery: supported
