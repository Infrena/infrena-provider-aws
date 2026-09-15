# aws.docdbelastic.cluster

**CloudFormation type:** `AWS::DocDBElastic::Cluster`

The AWS::DocDBElastic::Cluster Amazon DocumentDB (with MongoDB compatibility) Elastic Scale resource describes a Cluster

Region attribute: `region`

**Import ID:** `<region>/ClusterArn` (AWS::DocDBElastic::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdminUserName` | admin_user_name | `string` | required, replaces on change |  |  |
| `AdminUserPassword` | admin_user_password | `string` | optional, computed, provider-chosen, sensitive, write-only |  |  |
| `AuthType` | auth_type | `string` | required, replaces on change |  |  |
| `BackupRetentionPeriod` | backup_retention_period | `integer` | optional, computed, provider-chosen |  |  |
| `ClusterArn` | cluster_arn | `string` | computed |  |  |
| `ClusterEndpoint` | cluster_endpoint | `string` | computed |  |  |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PreferredBackupWindow` | preferred_backup_window | `string` | optional, computed, provider-chosen |  |  |
| `PreferredMaintenanceWindow` | preferred_maintenance_window | `string` | optional, computed, provider-chosen |  |  |
| `ShardCapacity` | shard_capacity | `integer` | required |  |  |
| `ShardCount` | shard_count | `integer` | required |  |  |
| `ShardInstanceCount` | shard_instance_count | `integer` | optional, computed, provider-chosen |  |  |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id |  |

Supports update: yes

Discovery: supported
