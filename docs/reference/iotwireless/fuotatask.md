# aws.fuotatask

**CloudFormation type:** `AWS::IoTWireless::FuotaTask`

Create and manage FUOTA tasks.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoTWireless::FuotaTask)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | FUOTA task arn. Returned after successful create. |
| `AssociateMulticastGroup` | associate_multicast_group | `string` | optional, computed, provider-chosen |  | Multicast group to associate. Only for update request. |
| `AssociateWirelessDevice` | associate_wireless_device | `string` | optional, computed, provider-chosen |  | Wireless device to associate. Only for update request. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | FUOTA task description |
| `DisassociateMulticastGroup` | disassociate_multicast_group | `string` | optional, computed, provider-chosen |  | Multicast group to disassociate. Only for update request. |
| `DisassociateWirelessDevice` | disassociate_wireless_device | `string` | optional, computed, provider-chosen |  | Wireless device to disassociate. Only for update request. |
| `FirmwareUpdateImage` | firmware_update_image | `string` | required |  | FUOTA task firmware update image binary S3 link |
| `FirmwareUpdateRole` | firmware_update_role | `string` | required |  | FUOTA task firmware IAM role for reading S3 |
| `FuotaTaskStatus` | fuota_task_status | `string` | computed |  | FUOTA task status. Returned after successful read. |
| `Id` |  | `string` | computed |  | FUOTA task id. Returned after successful create. |
| `LoRaWAN` | lo_ra_wan | `map` | required |  | FUOTA task LoRaWAN |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Name of FUOTA task |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the FUOTA task. |

Supports update: yes

Discovery: supported
