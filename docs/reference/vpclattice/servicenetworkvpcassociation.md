# aws.servicenetworkvpcassociation

**CloudFormation type:** `AWS::VpcLattice::ServiceNetworkVpcAssociation`

Associates a VPC with a service network.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ServiceNetworkVpcAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DnsOptions` | dns_options | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `PrivateDnsEnabled` | private_dns_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id |  |
| `ServiceNetworkArn` | service_network_arn | `string` | computed |  |  |
| `ServiceNetworkId` | service_network_id | `string` | computed |  |  |
| `ServiceNetworkIdentifier` | service_network_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ServiceNetworkName` | service_network_name | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcId` | vpc_id | `string` | computed |  |  |
| `VpcIdentifier` | vpc_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |

Supports update: yes

Discovery: supported
