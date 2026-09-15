# aws.assessmenttarget

**CloudFormation type:** `AWS::Inspector::AssessmentTarget`

Resource Type definition for AWS::Inspector::AssessmentTarget

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Inspector::AssessmentTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssessmentTargetName` | assessment_target_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ResourceGroupArn` | resource_group_arn | `string` | optional, computed, provider-chosen | aws.resourcegroup.Arn |  |

Supports update: yes

Discovery: supported
