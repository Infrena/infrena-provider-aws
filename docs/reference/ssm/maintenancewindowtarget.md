# aws.maintenancewindowtarget

**CloudFormation type:** `AWS::SSM::MaintenanceWindowTarget`

Resource type definition for AWS::SSM::MaintenanceWindowTarget

Region attribute: `region`

**Import ID:** `<region>/WindowId|WindowTargetId` (AWS::SSM::MaintenanceWindowTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the target. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name for the maintenance window target. |
| `OwnerInformation` | owner_information | `string` | optional, computed, provider-chosen |  | A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window. |
| `ResourceType` | resource_type | `string` | required |  | The type of target that is being registered with the maintenance window. |
| `Targets` |  | `list` | required |  | The targets to register with the maintenance window. |
| `WindowId` | window_id | `string` | required, replaces on change |  | The ID of the maintenance window to register the target with. |
| `WindowTargetId` | window_target_id | `string` | computed |  | The ID of the target. |

Supports update: yes

Discovery: supported (parent resource required)
