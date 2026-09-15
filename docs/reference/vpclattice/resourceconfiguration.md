# aws.resourceconfiguration

**CloudFormation type:** `AWS::VpcLattice::ResourceConfiguration`

VpcLattice ResourceConfiguration CFN resource

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ResourceConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowAssociationToSharableServiceNetwork` | allow_association_to_sharable_service_network | `boolean` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CustomDomainName` | custom_domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainVerificationId` | domain_verification_id | `string` | optional, computed, provider-chosen, replaces on change | aws.domainverification.Id |  |
| `GroupDomain` | group_domain | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `PortRanges` | port_ranges | `list` | optional, computed, provider-chosen |  |  |
| `ProtocolType` | protocol_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ResourceConfigurationAuthType` | resource_configuration_auth_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ResourceConfigurationDefinition` | resource_configuration_definition | `map` | optional, computed, provider-chosen |  |  |
| `ResourceConfigurationGroupId` | resource_configuration_group_id | `string` | optional, computed, provider-chosen, write-only |  |  |
| `ResourceConfigurationType` | resource_configuration_type | `string` | required, replaces on change |  |  |
| `ResourceGatewayId` | resource_gateway_id | `string` | optional, computed, provider-chosen, replaces on change | aws.resourcegateway.Id |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
