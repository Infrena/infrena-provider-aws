# aws.limitsprofile

**CloudFormation type:** `AWS::QuickSight::LimitsProfile`

Definition of AWS::QuickSight::LimitsProfile Resource Type

Region attribute: `region`

**Import ID:** `<region>/AccountId|ProfileId` (AWS::QuickSight::LimitsProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | required, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ProfileId` | profile_id | `string` | computed |  |  |
| `ProfileName` | profile_name | `string` | required |  |  |
| `ResourceLimits` | resource_limits | `map` | required |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
