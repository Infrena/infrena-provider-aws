# aws.networkanalyzerconfiguration

**CloudFormation type:** `AWS::IoTWireless::NetworkAnalyzerConfiguration`

Create and manage NetworkAnalyzerConfiguration resource.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTWireless::NetworkAnalyzerConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn for network analyzer configuration, Returned upon successful create. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the new resource |
| `Name` |  | `string` | required, replaces on change |  | Name of the network analyzer configuration |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TraceContent` | trace_content | `map` | optional, computed, provider-chosen |  | Trace content for your wireless gateway and wireless device resources |
| `WirelessDevices` | wireless_devices | `list` | optional, computed, provider-chosen |  | List of wireless gateway resources that have been added to the network analyzer configuration |
| `WirelessGateways` | wireless_gateways | `list` | optional, computed, provider-chosen |  | List of wireless gateway resources that have been added to the network analyzer configuration |

Supports update: yes

Discovery: supported
