# aws.mwaaserverless.workflow

**CloudFormation type:** `AWS::MWAAServerless::Workflow`

Resource Type definition for AWS::MWAAServerless::Workflow resource

Region attribute: `region`

**Import ID:** `<region>/WorkflowArn` (AWS::MWAAServerless::Workflow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Code` |  | `map` | optional, computed, provider-chosen |  | The location of code artifacts in Amazon S3 for the workflow. Modeled as a single-member container so it stays extensible to future artifact types (e.g. OCI images). |
| `CodeSnapshottedAt` | code_snapshotted_at | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DefinitionS3Location` | definition_s3_location | `map` | required |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `ScheduleConfiguration` | schedule_configuration | `map` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of key-value pairs to be applied as tags |
| `TriggerMode` | trigger_mode | `string` | optional, computed, provider-chosen |  |  |
| `WorkflowArn` | workflow_arn | `string` | computed |  |  |
| `WorkflowStatus` | workflow_status | `string` | computed |  |  |
| `WorkflowVersion` | workflow_version | `string` | computed |  |  |

Supports update: yes

Discovery: supported
