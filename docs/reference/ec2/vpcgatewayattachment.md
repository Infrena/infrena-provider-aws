# aws.vpcgatewayattachment

**CloudFormation type:** `AWS::EC2::VPCGatewayAttachment`

Resource Type definition for AWS::EC2::VPCGatewayAttachment

Region attribute: `region`

**Import ID:** `<region>/AttachmentType|VpcId` (AWS::EC2::VPCGatewayAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentType` | attachment_type | `string` | computed |  | Used to identify if this resource is an Internet Gateway or Vpn Gateway Attachment |
| `InternetGatewayId` | internet_gateway_id | `string` | optional, computed, provider-chosen | aws.internetgateway.InternetGatewayId | The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC. |
| `VpnGatewayId` | vpn_gateway_id | `string` | optional, computed, provider-chosen | aws.vpngateway.VPNGatewayId | The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both. |

Supports update: yes

Discovery: supported
