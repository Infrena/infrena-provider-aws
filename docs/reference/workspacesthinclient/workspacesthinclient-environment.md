# aws.workspacesthinclient.environment

**CloudFormation type:** `AWS::WorkSpacesThinClient::Environment`

Resource type definition for AWS::WorkSpacesThinClient::Environment.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::WorkSpacesThinClient::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActivationCode` | activation_code | `string` | computed |  | Activation code for devices associated with environment. |
| `Arn` |  | `string` | computed |  | The environment ARN. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp in unix epoch format when environment was created. |
| `DesiredSoftwareSetId` | desired_software_set_id | `string` | optional, computed, provider-chosen |  | The ID of the software set to apply. |
| `DesktopArn` | desktop_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the desktop to stream from Amazon WorkSpaces, WorkSpaces Web, or AppStream 2.0. |
| `DesktopEndpoint` | desktop_endpoint | `string` | optional, computed, provider-chosen |  | The URL for the identity provider login (only for environments that use AppStream 2.0). |
| `DesktopType` | desktop_type | `string` | computed |  | The type of VDI. |
| `DeviceCreationTags` | device_creation_tags | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to the newly created devices for this environment. |
| `Id` |  | `string` | computed |  | Unique identifier of the environment. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the AWS Key Management Service key used to encrypt the environment. |
| `MaintenanceWindow` | maintenance_window | `map` | optional, computed, provider-chosen |  | A specification for a time window to apply software updates. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the environment. |
| `PendingSoftwareSetId` | pending_software_set_id | `string` | computed |  | The ID of the software set that is pending to be installed. |
| `PendingSoftwareSetVersion` | pending_software_set_version | `string` | computed |  | The version of the software set that is pending to be installed. |
| `RegisteredDevicesCount` | registered_devices_count | `integer` | computed |  | Number of devices registered to the environment. |
| `SoftwareSetComplianceStatus` | software_set_compliance_status | `string` | computed |  | Describes if the software currently installed on all devices in the environment is a supported version. |
| `SoftwareSetUpdateMode` | software_set_update_mode | `string` | optional, computed, provider-chosen |  | An option to define which software updates to apply. |
| `SoftwareSetUpdateSchedule` | software_set_update_schedule | `string` | optional, computed, provider-chosen |  | An option to define if software updates should be applied within a maintenance window. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp in unix epoch format when environment was last updated. |

Supports update: yes

Discovery: supported
