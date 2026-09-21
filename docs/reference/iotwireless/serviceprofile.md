# aws.serviceprofile

**CloudFormation type:** `AWS::IoTWireless::ServiceProfile`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::ServiceProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Service profile Arn. Returned after successful create. |
| `Id` |  | `string` | computed |  | Service profile Id. Returned after successful create. |
| `LoRaWAN` | lo_ra_wan | `map` | optional, computed, provider-chosen, replaces on change |  | LoRaWAN supports all LoRa specific attributes for service profile for CreateServiceProfile operation |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of service profile |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the service profile. |

Supports update: yes

Discovery: supported
