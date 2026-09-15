# aws.refactorspaces.environment

**CloudFormation type:** `AWS::RefactorSpaces::Environment`

Definition of AWS::RefactorSpaces::Environment Resource Type

Region attribute: `region`

**Import ID:** `<region>/EnvironmentIdentifier` (AWS::RefactorSpaces::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `EnvironmentIdentifier` | environment_identifier | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `NetworkFabricType` | network_fabric_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair. |
| `TransitGatewayId` | transit_gateway_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
