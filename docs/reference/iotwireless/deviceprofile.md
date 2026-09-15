# aws.deviceprofile

**CloudFormation type:** `AWS::IoTWireless::DeviceProfile`

Device Profile's resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::DeviceProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Service profile Arn. Returned after successful create. |
| `Id` |  | `string` | computed |  | Service profile Id. Returned after successful create. |
| `LoRaWAN` | lo_ra_wan | `map` | optional, computed, provider-chosen, replaces on change |  | LoRaWANDeviceProfile supports all LoRa specific attributes for service profile for CreateDeviceProfile operation |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of service profile |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the device profile. |

Supports update: yes

Discovery: supported
