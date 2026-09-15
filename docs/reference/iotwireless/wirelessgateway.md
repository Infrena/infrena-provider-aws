# aws.wirelessgateway

**CloudFormation type:** `AWS::IoTWireless::WirelessGateway`

Create and manage wireless gateways, including LoRa gateways.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::WirelessGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn for Wireless Gateway. Returned upon successful create. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of Wireless Gateway. |
| `Id` |  | `string` | computed |  | Id for Wireless Gateway. Returned upon successful create. |
| `LastUplinkReceivedAt` | last_uplink_received_at | `string` | optional, computed, provider-chosen |  | The date and time when the most recent uplink was received. |
| `LoRaWAN` | lo_ra_wan | `map` | required |  | The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Gateway. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Name of Wireless Gateway. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the gateway. |
| `ThingArn` | thing_arn | `string` | optional, computed, provider-chosen |  | Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway. |
| `ThingName` | thing_name | `string` | optional, computed, provider-chosen |  | Thing Name. If there is a Thing created, this can be returned with a Get call. |

Supports update: yes

Discovery: supported
