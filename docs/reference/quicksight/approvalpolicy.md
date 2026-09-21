# aws.approvalpolicy

**CloudFormation type:** `AWS::QuickSight::ApprovalPolicy`

Definition of the AWS::QuickSight::ApprovalPolicy Resource Type.

Region attribute: `region`

**Import ID:** `<region>/PolicyId` (AWS::QuickSight::ApprovalPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | List of governed actions a policy applies to. |
| `ApplicableTo` | applicable_to | `map` | required |  | Scoping: who the policy applies to. |
| `ApprovalGroups` | approval_groups | `list` | required |  | List of approval group ARNs (e.g. QuickSight group ARNs). |
| `AssetTypes` | asset_types | `list` | required |  | List of asset types a policy applies to. |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `PolicyArn` | policy_arn | `string` | computed |  |  |
| `PolicyId` | policy_id | `string` | required, replaces on change |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
