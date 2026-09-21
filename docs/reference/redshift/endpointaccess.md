# aws.endpointaccess

**CloudFormation type:** `AWS::Redshift::EndpointAccess`

Resource schema for a Redshift-managed VPC endpoint.

Region attribute: `region`

**Import ID:** `<region>/EndpointName` (AWS::Redshift::EndpointAccess)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Address` |  | `string` | computed |  | The DNS address of the endpoint. |
| `ClusterIdentifier` | cluster_identifier | `string` | required, replaces on change |  | A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account |
| `EndpointCreateTime` | endpoint_create_time | `string` | computed |  | The time (UTC) that the endpoint was created. |
| `EndpointName` | endpoint_name | `string` | required, replaces on change |  | The name of the endpoint. |
| `EndpointStatus` | endpoint_status | `string` | computed |  | The status of the endpoint. |
| `Port` |  | `integer` | computed |  | The port number on which the cluster accepts incoming connections. |
| `ResourceOwner` | resource_owner | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS account ID of the owner of the cluster. |
| `SubnetGroupName` | subnet_group_name | `string` | required, replaces on change |  | The subnet group name where Amazon Redshift chooses to deploy the endpoint. |
| `VpcEndpoint` | vpc_endpoint | `map` | computed |  | The connection endpoint for connecting to an Amazon Redshift cluster through the proxy. |
| `VpcSecurityGroupIds` | vpc_security_group_ids | `list` | required | aws.securitygroup.Id | A list of vpc security group ids to apply to the created endpoint access. |
| `VpcSecurityGroups` | vpc_security_groups | `list` | computed |  | A list of Virtual Private Cloud (VPC) security groups to be associated with the endpoint. |

Supports update: yes

Discovery: supported
