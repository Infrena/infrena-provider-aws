# aws.dbproxyendpoint

**CloudFormation type:** `AWS::RDS::DBProxyEndpoint`

Resource schema for AWS::RDS::DBProxyEndpoint.

Region attribute: `region`

**Import ID:** `<region>/DBProxyEndpointName` (AWS::RDS::DBProxyEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DBProxyEndpointArn` | db_proxy_endpoint_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the DB proxy endpoint. |
| `DBProxyEndpointName` | db_proxy_endpoint_name | `string` | required, replaces on change |  | The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region. |
| `DBProxyName` | db_proxy_name | `string` | required, replaces on change |  | The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region. |
| `Endpoint` |  | `string` | computed |  | The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application. |
| `EndpointNetworkType` | endpoint_network_type | `string` | optional, computed, provider-chosen, replaces on change |  | The network type of the DB proxy endpoint. The network type determines the IP version that the proxy endpoint supports. |
| `IsDefault` | is_default | `boolean` | computed |  | A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only. |
| `Tags` |  | `map` | tags map |  | An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint. |
| `TargetRole` | target_role | `string` | optional, computed, provider-chosen |  | A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations. |
| `VpcId` | vpc_id | `string` | computed |  | VPC ID to associate with the new DB proxy endpoint. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | VPC security group IDs to associate with the new DB proxy endpoint. |
| `VpcSubnetIds` | vpc_subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId | VPC subnet IDs to associate with the new DB proxy endpoint. |

Supports update: yes

Discovery: supported
