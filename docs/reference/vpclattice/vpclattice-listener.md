# aws.vpclattice.listener

**CloudFormation type:** `AWS::VpcLattice::Listener`

Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::Listener)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DefaultAction` | default_action | `map` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Port` |  | `integer` | optional, computed, provider-chosen, replaces on change |  |  |
| `Protocol` |  | `string` | required, replaces on change |  |  |
| `ServiceArn` | service_arn | `string` | computed |  |  |
| `ServiceId` | service_id | `string` | computed |  |  |
| `ServiceIdentifier` | service_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
