# aws.medialive.network

**CloudFormation type:** `AWS::MediaLive::Network`

Resource schema for AWS::MediaLive::Network.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaLive::Network)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Network. |
| `AssociatedClusterIds` | associated_cluster_ids | `list` | computed |  |  |
| `Id` |  | `string` | computed |  | The unique ID of the Network. |
| `IpPools` | ip_pools | `list` | required |  | The list of IP address cidr pools for the network |
| `Name` |  | `string` | required |  | The user-specified name of the Network to be created. |
| `Routes` |  | `list` | optional, computed, provider-chosen |  | The routes for the network |
| `State` |  | `string` | computed |  | The current state of the Network. |
| `Tags` |  | `map` | tags map |  | A collection of key-value pairs. |

Supports update: yes

Discovery: supported
