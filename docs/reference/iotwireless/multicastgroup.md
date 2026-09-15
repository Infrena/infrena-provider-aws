# aws.multicastgroup

**CloudFormation type:** `AWS::IoTWireless::MulticastGroup`

Create and manage Multicast groups.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::MulticastGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Multicast group arn. Returned after successful create. |
| `AssociateWirelessDevice` | associate_wireless_device | `string` | optional, computed, provider-chosen |  | Wireless device to associate. Only for update request. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Multicast group description |
| `DisassociateWirelessDevice` | disassociate_wireless_device | `string` | optional, computed, provider-chosen |  | Wireless device to disassociate. Only for update request. |
| `Id` |  | `string` | computed |  | Multicast group id. Returned after successful create. |
| `LoRaWAN` | lo_ra_wan | `map` | required |  | Multicast group LoRaWAN |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Name of Multicast group |
| `Status` |  | `string` | computed |  | Multicast group status. Returned after successful read. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the Multicast group. |

Supports update: yes

Discovery: supported
