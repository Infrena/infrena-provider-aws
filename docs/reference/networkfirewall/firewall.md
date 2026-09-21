# aws.firewall

**CloudFormation type:** `AWS::NetworkFirewall::Firewall`

Resource type definition for AWS::NetworkFirewall::Firewall

Region attribute: `region`

**Import ID:** `<region>/FirewallArn` (AWS::NetworkFirewall::Firewall)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZoneChangeProtection` | availability_zone_change_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `AvailabilityZoneMappings` | availability_zone_mappings | `list` | optional, computed, provider-chosen |  |  |
| `DeleteProtection` | delete_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnabledAnalysisTypes` | enabled_analysis_types | `list` | optional, computed, provider-chosen |  | The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both. |
| `EndpointIds` | endpoint_ids | `list` | computed |  |  |
| `FirewallArn` | firewall_arn | `string` | computed |  | A resource ARN. |
| `FirewallId` | firewall_id | `string` | computed |  |  |
| `FirewallName` | firewall_name | `string` | required, replaces on change |  |  |
| `FirewallPolicyArn` | firewall_policy_arn | `string` | required | aws.firewallpolicy.FirewallPolicyArn | A resource ARN. |
| `FirewallPolicyChangeProtection` | firewall_policy_change_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `SubnetChangeProtection` | subnet_change_protection | `boolean` | optional, computed, provider-chosen |  |  |
| `SubnetMappings` | subnet_mappings | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TransitGatewayAttachmentId` | transit_gateway_attachment_id | `string` | computed |  |  |
| `TransitGatewayId` | transit_gateway_id | `string` | optional, computed, provider-chosen | aws.transitgateway.Id |  |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
