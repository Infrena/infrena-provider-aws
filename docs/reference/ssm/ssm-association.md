# aws.ssm.association

**CloudFormation type:** `AWS::SSM::Association`

The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.

Region attribute: `region`

**Import ID:** `<region>/AssociationId` (AWS::SSM::Association)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplyOnlyAtCronInterval` | apply_only_at_cron_interval | `boolean` | optional, computed, provider-chosen |  |  |
| `AssociationDispatchAssumeRole` | association_dispatch_assume_role | `string` | optional, computed, provider-chosen |  | A role used by association to take actions on your behalf. |
| `AssociationId` | association_id | `string` | computed |  | Unique identifier of the association. |
| `AssociationName` | association_name | `string` | optional, computed, provider-chosen |  | The name of the association. |
| `AutomationTargetParameterName` | automation_target_parameter_name | `string` | optional, computed, provider-chosen |  |  |
| `CalendarNames` | calendar_names | `list` | optional, computed, provider-chosen |  |  |
| `ComplianceSeverity` | compliance_severity | `string` | optional, computed, provider-chosen |  |  |
| `DocumentVersion` | document_version | `string` | optional, computed, provider-chosen |  | The version of the SSM document to associate with the target. |
| `InstanceId` | instance_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the instance that the SSM document is associated with. |
| `MaxConcurrency` | max_concurrency | `string` | optional, computed, provider-chosen |  |  |
| `MaxErrors` | max_errors | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  | The name of the SSM document. |
| `OutputLocation` | output_location | `map` | optional, computed, provider-chosen |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | Parameter values that the SSM document uses at runtime. |
| `ScheduleExpression` | schedule_expression | `string` | optional, computed, provider-chosen |  | A Cron or Rate expression that specifies when the association is applied to the target. |
| `ScheduleOffset` | schedule_offset | `integer` | optional, computed, provider-chosen |  |  |
| `SyncCompliance` | sync_compliance | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A key-value pair to associate with a resource. |
| `Targets` |  | `list` | optional, computed, provider-chosen |  | The targets that the SSM document sends commands to. |
| `WaitForSuccessTimeoutSeconds` | wait_for_success_timeout_seconds | `integer` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
