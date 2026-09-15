# aws.dbproxytargetgroup

**CloudFormation type:** `AWS::RDS::DBProxyTargetGroup`

Resource schema for AWS::RDS::DBProxyTargetGroup

Region attribute: `region`

**Import ID:** `<region>/TargetGroupArn` (AWS::RDS::DBProxyTargetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionPoolConfigurationInfo` | connection_pool_configuration_info | `map` | optional, computed, provider-chosen |  |  |
| `DBClusterIdentifiers` | db_cluster_identifiers | `list` | optional, computed, provider-chosen |  |  |
| `DBInstanceIdentifiers` | db_instance_identifiers | `list` | optional, computed, provider-chosen |  |  |
| `DBProxyName` | db_proxy_name | `string` | required, replaces on change |  | The identifier for the proxy. |
| `TargetGroupArn` | target_group_arn | `string` | computed |  | The Amazon Resource Name (ARN) representing the target group. |
| `TargetGroupName` | target_group_name | `string` | required, replaces on change |  | The identifier for the DBProxyTargetGroup |

Supports update: yes

Discovery: supported
