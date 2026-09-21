# aws.servicenetwork

**CloudFormation type:** `AWS::VpcLattice::ServiceNetwork`

A service network is a logical boundary for a collection of services. You can associate services and VPCs with a service network.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::ServiceNetwork)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AuthType` | auth_type | `string` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `SharingConfig` | sharing_config | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
