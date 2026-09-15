# aws.restoretestingselection

**CloudFormation type:** `AWS::Backup::RestoreTestingSelection`

Resource Type definition for AWS::Backup::RestoreTestingSelection

Region attribute: `region`

**Import ID:** `<region>/RestoreTestingPlanName|RestoreTestingSelectionName` (AWS::Backup::RestoreTestingSelection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IamRoleArn` | iam_role_arn | `string` | required | aws.role.Arn |  |
| `ProtectedResourceArns` | protected_resource_arns | `list` | optional, computed, provider-chosen |  |  |
| `ProtectedResourceConditions` | protected_resource_conditions | `map` | optional, computed, provider-chosen |  |  |
| `ProtectedResourceType` | protected_resource_type | `string` | required, replaces on change |  |  |
| `RestoreMetadataOverrides` | restore_metadata_overrides | `map` | optional, computed, provider-chosen |  |  |
| `RestoreTestingPlanName` | restore_testing_plan_name | `string` | required, replaces on change |  |  |
| `RestoreTestingSelectionName` | restore_testing_selection_name | `string` | required, replaces on change |  |  |
| `ValidationWindowHours` | validation_window_hours | `integer` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
