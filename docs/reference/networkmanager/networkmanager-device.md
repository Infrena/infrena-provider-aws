# aws.networkmanager.device

**CloudFormation type:** `AWS::NetworkManager::Device`

The AWS::NetworkManager::Device type describes a device.

Region attribute: `region`

**Import ID:** `<region>/GlobalNetworkId|DeviceId` (AWS::NetworkManager::Device)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AWSLocation` | aws_location | `map` | optional, computed, provider-chosen |  | The Amazon Web Services location of the device, if applicable. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the device was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the device. |
| `DeviceArn` | device_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the device. |
| `DeviceId` | device_id | `string` | computed |  | The ID of the device. |
| `GlobalNetworkId` | global_network_id | `string` | required, replaces on change | aws.globalnetwork.Id | The ID of the global network. |
| `Location` |  | `map` | optional, computed, provider-chosen |  | The site location. |
| `Model` |  | `string` | optional, computed, provider-chosen |  | The device model |
| `SerialNumber` | serial_number | `string` | optional, computed, provider-chosen |  | The device serial number. |
| `SiteId` | site_id | `string` | optional, computed, provider-chosen | aws.networkmanager.site.SiteId | The site ID. |
| `State` |  | `string` | computed |  | The state of the device. |
| `Tags` |  | `map` | tags map |  | The tags for the device. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | The device type. |
| `Vendor` |  | `string` | optional, computed, provider-chosen |  | The device vendor. |

Supports update: yes

Discovery: supported (parent resource required)
