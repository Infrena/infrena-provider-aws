# aws.vpcendpointassociation

**CloudFormation type:** `AWS::NetworkFirewall::VpcEndpointAssociation`

Resource type definition for AWS::NetworkFirewall::VpcEndpointAssociation

Region attribute: `region`

**Import ID:** `<region>/VpcEndpointAssociationArn` (AWS::NetworkFirewall::VpcEndpointAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `EndpointId` | endpoint_id | `string` | computed |  | An endpoint Id. |
| `FirewallArn` | firewall_arn | `string` | required, replaces on change | aws.firewall.FirewallArn | A resource ARN. |
| `SubnetMapping` | subnet_mapping | `map` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcEndpointAssociationArn` | vpc_endpoint_association_arn | `string` | computed |  | A resource ARN. |
| `VpcEndpointAssociationId` | vpc_endpoint_association_id | `string` | computed |  |  |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
