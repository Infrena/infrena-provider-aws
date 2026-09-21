# aws.transitgateway

**CloudFormation type:** `AWS::EC2::TransitGateway`

Resource Type definition for AWS::EC2::TransitGateway

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::TransitGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonSideAsn` | amazon_side_asn | `integer` | optional, computed, provider-chosen, replaces on change |  |  |
| `AssociationDefaultRouteTableId` | association_default_route_table_id | `string` | optional, computed, provider-chosen |  |  |
| `AutoAcceptSharedAttachments` | auto_accept_shared_attachments | `string` | optional, computed, provider-chosen |  |  |
| `DefaultRouteTableAssociation` | default_route_table_association | `string` | optional, computed, provider-chosen |  |  |
| `DefaultRouteTablePropagation` | default_route_table_propagation | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DnsSupport` | dns_support | `string` | optional, computed, provider-chosen |  |  |
| `EncryptionSupport` | encryption_support | `string` | optional, computed, provider-chosen, write-only |  |  |
| `EncryptionSupportState` | encryption_support_state | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `MulticastSupport` | multicast_support | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PropagationDefaultRouteTableId` | propagation_default_route_table_id | `string` | optional, computed, provider-chosen |  |  |
| `SecurityGroupReferencingSupport` | security_group_referencing_support | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TransitGatewayArn` | transit_gateway_arn | `string` | computed |  |  |
| `TransitGatewayCidrBlocks` | transit_gateway_cidr_blocks | `list` | optional, computed, provider-chosen |  |  |
| `VpnEcmpSupport` | vpn_ecmp_support | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
