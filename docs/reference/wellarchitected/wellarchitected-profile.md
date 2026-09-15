# aws.wellarchitected.profile

**CloudFormation type:** `AWS::WellArchitected::Profile`

Definition of AWS::WellArchitected::Profile Resource Type

Region attribute: `region`

**Import ID:** `<region>/ProfileArn` (AWS::WellArchitected::Profile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The date and time the profile was created. |
| `Owner` |  | `string` | computed |  | The owner of the profile. |
| `ProfileArn` | profile_arn | `string` | computed |  | The profile ARN. |
| `ProfileDescription` | profile_description | `string` | required |  | The profile description. |
| `ProfileName` | profile_name | `string` | required, replaces on change |  | The name of the profile. |
| `ProfileQuestions` | profile_questions | `list` | required |  | The profile questions. |
| `ProfileVersion` | profile_version | `string` | computed |  | The profile version. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the profile. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the profile was last updated. |

Supports update: yes

Discovery: supported
