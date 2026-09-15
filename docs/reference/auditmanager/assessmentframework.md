# aws.assessmentframework

**CloudFormation type:** `AWS::AuditManager::AssessmentFramework`

Creates a custom framework in AWS Audit Manager.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AuditManager::AssessmentFramework)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the framework. |
| `ComplianceType` | compliance_type | `string` | optional, computed, provider-chosen |  | The compliance type that the framework supports, such as CIS or HIPAA. |
| `ControlSets` | control_sets | `list` | required |  | The control sets that are associated with the framework. |
| `CreatedAt` | created_at | `string` | computed |  | The time when the framework was created. |
| `CreatedBy` | created_by | `string` | computed |  | The user or role that created the framework. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the framework. |
| `FrameworkId` | framework_id | `string` | computed |  | The unique identifier for the framework. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The time when the framework was most recently updated. |
| `LastUpdatedBy` | last_updated_by | `string` | computed |  | The user or role that most recently updated the framework. |
| `Name` |  | `string` | required |  | The name of the framework. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the framework. |
| `Type` | type_value | `string` | computed |  | The framework type, such as a standard framework or a custom framework. |

Supports update: yes

Discovery: supported
