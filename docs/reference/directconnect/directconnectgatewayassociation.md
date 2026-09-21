# aws.directconnectgatewayassociation

**CloudFormation type:** `AWS::DirectConnect::DirectConnectGatewayAssociation`

Resource Type definition for AWS::DirectConnect::DirectConnectGatewayAssociation

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::DirectConnect::DirectConnectGatewayAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptDirectConnectGatewayAssociationProposalRoleArn` | accept_direct_connect_gateway_association_proposal_role_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of the role to accept the Direct Connect Gateway association proposal. Needs directconnect:AcceptDirectConnectGatewayAssociationProposal permissions. |
| `AllowedPrefixesToDirectConnectGateway` | allowed_prefixes_to_direct_connect_gateway | `list` | optional, computed, provider-chosen |  | The Amazon VPC prefixes to advertise to the Direct Connect gateway. This parameter is required when you create an association to a transit gateway. |
| `AssociatedGatewayId` | associated_gateway_id | `string` | required, replaces on change |  | The ID or ARN of the virtual private gateway or transit gateway. |
| `AssociationId` | association_id | `string` | computed |  | The ID of the Direct Connect gateway association. |
| `DirectConnectGatewayId` | direct_connect_gateway_id | `string` | required, replaces on change | aws.directconnectgateway.DirectConnectGatewayId | The ID or ARN of the Direct Connect gateway. |

Supports update: yes

Discovery: supported (parent resource required)
