# aws.approvalteam

**CloudFormation type:** `AWS::MPA::ApprovalTeam`

Resource Type definition for AWS::MPA::ApprovalTeam.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MPA::ApprovalTeam)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApprovalStrategy` | approval_strategy | `map` | required |  |  |
| `Approvers` |  | `list` | required |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `Description` |  | `string` | required |  |  |
| `LastUpdateTime` | last_update_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `NumberOfApprovers` | number_of_approvers | `integer` | computed |  |  |
| `Policies` |  | `list` | required, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusCode` | status_code | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UpdateSessionArn` | update_session_arn | `string` | computed |  |  |
| `VersionId` | version_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
