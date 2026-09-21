# aws.wirelessdeviceimporttask

**CloudFormation type:** `AWS::IoTWireless::WirelessDeviceImportTask`

Wireless Device Import Tasks

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::WirelessDeviceImportTask)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn for Wireless Device Import Task, Returned upon successful start. |
| `CreationDate` | creation_date | `string` | computed |  | CreationDate for import task |
| `DestinationName` | destination_name | `string` | required |  | Destination Name for import task |
| `FailedImportedDevicesCount` | failed_imported_devices_count | `integer` | computed |  | Failed Imported Devices Count |
| `Id` |  | `string` | computed |  | Id for Wireless Device Import Task, Returned upon successful start. |
| `InitializedImportedDevicesCount` | initialized_imported_devices_count | `integer` | computed |  | Initialized Imported Devices Count |
| `OnboardedImportedDevicesCount` | onboarded_imported_devices_count | `integer` | computed |  | Onboarded Imported Devices Count |
| `PendingImportedDevicesCount` | pending_imported_devices_count | `integer` | computed |  | Pending Imported Devices Count |
| `Sidewalk` |  | `map` | required |  | sidewalk contain file for created device and role |
| `Status` |  | `string` | computed |  | Status for import task |
| `StatusReason` | status_reason | `string` | computed |  | StatusReason for import task |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
