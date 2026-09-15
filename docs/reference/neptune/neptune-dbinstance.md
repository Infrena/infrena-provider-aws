# aws.neptune.dbinstance

**CloudFormation type:** `AWS::Neptune::DBInstance`

The AWS::Neptune::DBInstance resource creates an Amazon Neptune DB instance.

Region attribute: `region`

**Import ID:** `<region>/DBInstanceIdentifier` (AWS::Neptune::DBInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowMajorVersionUpgrade` | allow_major_version_upgrade | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates that major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible. This parameter must be set to true when specifying a value for the EngineVersion parameter that is a different major version than the DB instance's current version. |
| `AutoMinorVersionUpgrade` | auto_minor_version_upgrade | `boolean` | optional, computed, provider-chosen |  | Indicates that minor version patches are applied automatically. |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the name of the Availability Zone the DB instance is located in. |
| `DBClusterIdentifier` | db_cluster_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | If the DB instance is a member of a DB cluster, contains the name of the DB cluster that the DB instance is a member of. |
| `DBInstanceClass` | db_instance_class | `string` | required |  | Contains the name of the compute and memory capacity class of the DB instance. |
| `DBInstanceIdentifier` | db_instance_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | Contains a user-supplied database identifier. This identifier is the unique key that identifies a DB instance. |
| `DBParameterGroupName` | db_parameter_group_name | `string` | optional, computed, provider-chosen |  | The name of an existing DB parameter group or a reference to an AWS::Neptune::DBParameterGroup resource created in the template. If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot. |
| `DBSnapshotIdentifier` | db_snapshot_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | This parameter is not supported. |
| `DBSubnetGroupName` | db_subnet_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | A DB subnet group to associate with the DB instance. If you update this value, the new subnet group must be a subnet group in a new virtual private cloud (VPC). |
| `Endpoint` |  | `string` | computed |  | The connection endpoint for the database. For example: `mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`. |
| `Port` |  | `string` | computed |  | The port number on which the database accepts connections. For example: `8182`. |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  | Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC). |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen |  | Indicates that public accessibility is enabled. This should be enabled in combination with IAM Auth enabled on the DBCluster |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this DB instance. |

Supports update: yes

Discovery: supported
