# aws.routernetworkinterface

**CloudFormation type:** `AWS::MediaConnect::RouterNetworkInterface`

Represents a router network interface in AWS Elemental MediaConnect that is used to define a network boundary for router resources

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaConnect::RouterNetworkInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssociatedInputCount` | associated_input_count | `integer` | computed |  | The number of router inputs associated with the network interface. |
| `AssociatedOutputCount` | associated_output_count | `integer` | computed |  | The number of router outputs associated with the network interface. |
| `Configuration` |  | `string` | required |  | The configuration settings for a router network interface. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the router network interface was created. |
| `Id` |  | `string` | computed |  | The unique identifier of the router network interface. |
| `Name` |  | `string` | required |  | The name of the router network interface. |
| `NetworkInterfaceType` | network_interface_type | `string` | computed |  |  |
| `RegionName` | region_name | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Region for the router network interface. Defaults to the current region if not specified. |
| `State` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to tag and organize this router network interface. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the router network interface was last updated. |

Supports update: yes

Discovery: supported
