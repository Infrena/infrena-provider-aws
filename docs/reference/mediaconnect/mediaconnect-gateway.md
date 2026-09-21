# aws.mediaconnect.gateway

**CloudFormation type:** `AWS::MediaConnect::Gateway`

Resource schema for AWS::MediaConnect::Gateway

Region attribute: `region`

**Import ID:** `<region>/GatewayArn` (AWS::MediaConnect::Gateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EgressCidrBlocks` | egress_cidr_blocks | `list` | required, replaces on change |  | The range of IP addresses that contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16. |
| `GatewayArn` | gateway_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the gateway. |
| `GatewayState` | gateway_state | `string` | computed |  | The current status of the gateway. |
| `Name` |  | `string` | required, replaces on change |  | The name of the gateway. This name can not be modified after the gateway is created. |
| `Networks` |  | `list` | required, replaces on change |  | The list of networks in the gateway. |

Supports update: no

Discovery: supported
