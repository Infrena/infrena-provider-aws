# aws.apigateway.vpclink

**CloudFormation type:** `AWS::ApiGateway::VpcLink`

The ``AWS::ApiGateway::VpcLink`` resource creates an API Gateway VPC link for a REST API to access resources in an Amazon Virtual Private Cloud (VPC). For more information, see [vpclink:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateVpcLink.html) in the ``Amazon API Gateway REST API Reference``.

Region attribute: `region`

**Import ID:** `<region>/VpcLinkId` (AWS::ApiGateway::VpcLink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  | An array of arbitrary tags (key-value pairs) to associate with the VPC link. |
| `TargetArns` | target_arns | `list` | required, replaces on change |  |  |
| `VpcLinkId` | vpc_link_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
