# aws.servicenetworkresourceassociation

**CloudFormation type:** `AWS::VpcLattice::ServiceNetworkResourceAssociation`

VpcLattice ServiceNetworkResourceAssociation CFN resource

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ServiceNetworkResourceAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `PrivateDnsEnabled` | private_dns_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `ResourceConfigurationId` | resource_configuration_id | `string` | optional, computed, provider-chosen, replaces on change | aws.resourceconfiguration.Id |  |
| `ServiceNetworkId` | service_network_id | `string` | optional, computed, provider-chosen, replaces on change | aws.servicenetwork.Id |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
