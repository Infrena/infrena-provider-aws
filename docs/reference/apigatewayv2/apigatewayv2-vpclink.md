# aws.apigatewayv2.vpclink

**CloudFormation type:** `AWS::ApiGatewayV2::VpcLink`

The ``AWS::ApiGatewayV2::VpcLink`` resource creates a VPC link. This VPC link can be used with both REST and HTTP APIs. The VPC link status must transition from ``PENDING`` to ``AVAILABLE`` to successfully create a VPC link, which can take up to 10 minutes. To learn more, see [Working with VPC Links for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vpc-links.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/VpcLinkId` (AWS::ApiGatewayV2::VpcLink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Name` |  | `string` | required |  | The name of the VPC link. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | A list of security group IDs for the VPC link. |
| `SubnetIds` | subnet_ids | `list` | required, replaces on change | aws.subnet.SubnetId | A list of subnet IDs to include in the VPC link. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The collection of tags. Each tag element is associated with a given resource. |
| `VpcLinkId` | vpc_link_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
