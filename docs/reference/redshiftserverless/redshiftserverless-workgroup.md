# aws.redshiftserverless.workgroup

**CloudFormation type:** `AWS::RedshiftServerless::Workgroup`

Definition of AWS::RedshiftServerless::Workgroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/WorkgroupName` (AWS::RedshiftServerless::Workgroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BaseCapacity` | base_capacity | `integer` | optional, computed, provider-chosen |  | The base compute capacity of the workgroup in Redshift Processing Units (RPUs). |
| `ConfigParameters` | config_parameters | `list` | optional, computed, provider-chosen, write-only |  | A list of parameters to set for finer control over a database. Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl. |
| `EnhancedVpcRouting` | enhanced_vpc_routing | `boolean` | optional, computed, provider-chosen |  | The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC. |
| `MaxCapacity` | max_capacity | `integer` | optional, computed, provider-chosen |  | The max compute capacity of the workgroup in Redshift Processing Units (RPUs). |
| `NamespaceName` | namespace_name | `string` | optional, computed, provider-chosen, replaces on change |  | The namespace the workgroup is associated with. |
| `Port` |  | `integer` | optional, computed, provider-chosen |  | The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439. |
| `PricePerformanceTarget` | price_performance_target | `map` | optional, computed, provider-chosen |  | A property that represents the price performance target settings for the workgroup. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen |  | A value that specifies whether the workgroup can be accessible from a public network. |
| `RecoveryPointId` | recovery_point_id | `string` | optional, computed, provider-chosen, write-only |  | The identifier of the recovery point to restore the namespace from. When this resource is first created, the namespace is restored from this recovery point. On subsequent updates, a restore occurs only when RecoveryPointId changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, write-only | aws.securitygroup.Id | A list of security group IDs to associate with the workgroup. |
| `SnapshotArn` | snapshot_arn | `string` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Name (ARN) of the snapshot to restore the namespace from. Specify either SnapshotArn or SnapshotName, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotArn changes from its previous value. If the value is unchanged or removed, no restore takes place and existing data is preserved. |
| `SnapshotName` | snapshot_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the snapshot to restore the namespace from. Because snapshot names are unique only within an account, also specify SnapshotOwnerAccount when restoring from a snapshot owned by a different account. Specify either SnapshotName or SnapshotArn, but not both. When this resource is first created, the namespace is restored from this snapshot. On subsequent updates, a restore occurs only when SnapshotName or SnapshotOwnerAccount changes from its previous value. If both values are unchanged or SnapshotName is removed, no restore takes place and existing data is preserved. |
| `SnapshotOwnerAccount` | snapshot_owner_account | `string` | optional, computed, provider-chosen, write-only |  | The AWS account ID that owns the snapshot. Required when restoring from a snapshot shared by another account. Used in combination with SnapshotName. On updates, changing this value while SnapshotName is set triggers a restore from the newly referenced snapshot. If the value is unchanged, no restore takes place and existing data is preserved. |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, write-only | aws.subnet.SubnetId | A list of subnet IDs the workgroup is associated with. |
| `Tags` |  | `map` | tags map |  | The map of the key-value pairs used to tag the workgroup. |
| `TrackName` | track_name | `string` | optional, computed, provider-chosen |  |  |
| `Workgroup` |  | `map` | optional, computed, provider-chosen |  | Definition for workgroup resource |
| `WorkgroupName` | workgroup_name | `string` | required, replaces on change |  | The name of the workgroup. |

Supports update: yes

Discovery: supported
