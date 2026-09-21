# aws.b2bi.profile

**CloudFormation type:** `AWS::B2BI::Profile`

Definition of AWS::B2BI::Profile Resource Type

Region attribute: `region`

**Import ID:** `<region>/ProfileId` (AWS::B2BI::Profile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BusinessName` | business_name | `string` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Email` |  | `string` | optional, computed, provider-chosen |  |  |
| `LogGroupName` | log_group_name | `string` | computed |  |  |
| `Logging` |  | `string` | required, replaces on change |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Phone` |  | `string` | required |  |  |
| `ProfileArn` | profile_arn | `string` | computed |  |  |
| `ProfileId` | profile_id | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
