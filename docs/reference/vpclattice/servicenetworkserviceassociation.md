# aws.servicenetworkserviceassociation

**CloudFormation type:** `AWS::VpcLattice::ServiceNetworkServiceAssociation`

Associates a service with a service network.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ServiceNetworkServiceAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DnsEntry` | dns_entry | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `ServiceArn` | service_arn | `string` | computed |  |  |
| `ServiceId` | service_id | `string` | computed |  |  |
| `ServiceIdentifier` | service_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ServiceName` | service_name | `string` | computed |  |  |
| `ServiceNetworkArn` | service_network_arn | `string` | computed |  |  |
| `ServiceNetworkId` | service_network_id | `string` | computed |  |  |
| `ServiceNetworkIdentifier` | service_network_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ServiceNetworkName` | service_network_name | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
