# aws.wirelessdevice

**CloudFormation type:** `AWS::IoTWireless::WirelessDevice`

Create and manage wireless gateways, including LoRa gateways.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::WirelessDevice)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Wireless device arn. Returned after successful create. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Wireless device description |
| `DestinationName` | destination_name | `string` | required |  | Wireless device destination name |
| `Id` |  | `string` | computed |  | Wireless device Id. Returned after successful create. |
| `LastUplinkReceivedAt` | last_uplink_received_at | `string` | optional, computed, provider-chosen |  | The date and time when the most recent uplink was received. |
| `LoRaWAN` | lo_ra_wan | `map` | optional, computed, provider-chosen |  | The combination of Package, Station and Model which represents the version of the LoRaWAN Wireless Device. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Wireless device name |
| `Positioning` |  | `string` | optional, computed, provider-chosen |  | FPort values for the GNSS, stream, and ClockSync functions of the positioning information. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the device. Currently not supported, will not create if tags are passed. |
| `ThingArn` | thing_arn | `string` | optional, computed, provider-chosen | aws.thing.Arn | Thing arn. Passed into update to associate Thing with Wireless device. |
| `ThingName` | thing_name | `string` | computed |  | Thing Arn. If there is a Thing created, this can be returned with a Get call. |
| `Type` | type_value | `string` | required |  | Wireless device type, currently only Sidewalk and LoRa |

Supports update: yes

Discovery: supported
