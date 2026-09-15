# aws.internetgateway

**CloudFormation type:** `AWS::EC2::InternetGateway`

Allocates an internet gateway for use with a VPC. After creating the Internet gateway, you then attach it to a VPC.

Region attribute: `region`

**Import ID:** `<region>/InternetGatewayId` (AWS::EC2::InternetGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InternetGatewayId` | internet_gateway_id | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags to assign to the internet gateway. |

Supports update: yes

Discovery: supported
