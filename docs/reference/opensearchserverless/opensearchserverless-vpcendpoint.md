# aws.opensearchserverless.vpcendpoint

**CloudFormation type:** `AWS::OpenSearchServerless::VpcEndpoint`

Resource Type definition for AWS::OpenSearchServerless::VpcEndpoint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::OpenSearchServerless::VpcEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | The identifier of the VPC Endpoint |
| `Name` |  | `string` | required, replaces on change |  | The name of the VPC Endpoint |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | The ID of one or more security groups to associate with the endpoint network interface |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The ID of one or more subnets in which to create an endpoint network interface |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC in which the endpoint will be used. |

Supports update: yes

Discovery: supported
