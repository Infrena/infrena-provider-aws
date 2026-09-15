# aws.managedpolicy

**CloudFormation type:** `AWS::IAM::ManagedPolicy`

Creates a new managed policy for your AWS-account.

Global type (no region attribute)

**Import ID:** `global/PolicyArn` (AWS::IAM::ManagedPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentCount` | attachment_count | `integer` | computed |  |  |
| `CreateDate` | create_date | `string` | computed |  |  |
| `DefaultVersionId` | default_version_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A friendly description of the policy. |
| `Groups` |  | `list` | optional, computed, provider-chosen |  | The name (friendly name, not ARN) of the group to attach the policy to. |
| `IsAttachable` | is_attachable | `boolean` | computed |  |  |
| `ManagedPolicyName` | name, managed_policy_name | `string` | optional, computed, provider-chosen, replaces on change |  | The friendly name of the policy. |
| `Path` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The path for the policy. |
| `PermissionsBoundaryUsageCount` | permissions_boundary_usage_count | `integer` | computed |  |  |
| `PolicyArn` | policy_arn | `string` | computed |  |  |
| `PolicyDocument` | policy, policy_document | `string` | required |  | The JSON policy document that you want to use as the content for the new policy. |
| `PolicyId` | policy_id | `string` | computed |  |  |
| `Roles` |  | `list` | optional, computed, provider-chosen |  | The name (friendly name, not ARN) of the role to attach the policy to. |
| `UpdateDate` | update_date | `string` | computed |  |  |
| `Users` |  | `list` | optional, computed, provider-chosen |  | The name (friendly name, not ARN) of the IAM user to attach the policy to. |

Supports update: yes

Discovery: supported
