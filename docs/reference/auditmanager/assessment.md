# aws.assessment

**CloudFormation type:** `AWS::AuditManager::Assessment`

An entity that defines the scope of audit evidence collected by AWS Audit Manager.

Region attribute: `region`

**Import ID:** `<region>/AssessmentId` (AWS::AuditManager::Assessment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the assessment. |
| `AssessmentId` | assessment_id | `string` | computed |  |  |
| `AssessmentReportsDestination` | assessment_reports_destination | `map` | optional, computed, provider-chosen |  | The destination in which evidence reports are stored for the specified assessment. |
| `AwsAccount` | aws_account | `map` | optional, computed, provider-chosen, replaces on change |  | The AWS account associated with the assessment. |
| `CreationTime` | creation_time | `float` | computed |  | The sequence of characters that identifies when the event occurred. |
| `Delegations` |  | `list` | optional, computed, provider-chosen |  | The list of delegations. |
| `Description` |  | `string` | optional, computed, provider-chosen, write-only |  | The description of the specified assessment. |
| `FrameworkId` | framework_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier for the specified framework. |
| `Name` |  | `string` | optional, computed, provider-chosen, write-only |  | The name of the related assessment. |
| `Roles` |  | `list` | optional, computed, provider-chosen |  | The list of roles for the specified assessment. |
| `Scope` |  | `map` | optional, computed, provider-chosen |  | The wrapper that contains the AWS accounts and AWS services in scope for the assessment. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the specified assessment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the assessment. |

Supports update: yes

Discovery: supported
