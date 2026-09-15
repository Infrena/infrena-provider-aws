# aws.dbproxy

**CloudFormation type:** `AWS::RDS::DBProxy`

Resource schema for AWS::RDS::DBProxy

Region attribute: `region`

**Import ID:** `<region>/DBProxyName` (AWS::RDS::DBProxy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Auth` |  | `list` | optional, computed, provider-chosen |  | The authorization mechanism that the proxy uses. |
| `DBProxyArn` | db_proxy_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the proxy. |
| `DBProxyName` | db_proxy_name | `string` | required, replaces on change |  | The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. |
| `DebugLogging` | debug_logging | `boolean` | optional, computed, provider-chosen |  | Whether the proxy includes detailed information about SQL statements in its logs. |
| `DefaultAuthScheme` | default_auth_scheme | `string` | optional, computed, provider-chosen |  | The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database. |
| `Endpoint` |  | `string` | computed |  | The endpoint that you can use to connect to the proxy. You include the endpoint value in the connection string for a database client application. |
| `EndpointNetworkType` | endpoint_network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports. |
| `EngineFamily` | engine_family | `string` | required, replaces on change |  | The kinds of databases that the proxy can connect to. |
| `IdleClientTimeout` | idle_client_timeout | `integer` | optional, computed, provider-chosen |  | The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it. |
| `RequireTLS` | require_tls | `boolean` | optional, computed, provider-chosen |  | A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy. |
| `TargetConnectionNetworkType` | target_connection_network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The network type that the proxy uses to connect to the target database. The network type determines the IP version that the proxy uses for connections to the database. |
| `VpcId` | vpc_id | `string` | computed |  | VPC ID to associate with the new DB proxy. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | VPC security group IDs to associate with the new proxy. |
| `VpcSubnetIds` | vpc_subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId | VPC subnet IDs to associate with the new proxy. |

Supports update: yes

Discovery: supported
