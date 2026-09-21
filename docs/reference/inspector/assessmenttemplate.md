# aws.assessmenttemplate

**CloudFormation type:** `AWS::Inspector::AssessmentTemplate`

Resource Type definition for AWS::Inspector::AssessmentTemplate

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Inspector::AssessmentTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssessmentTargetArn` | assessment_target_arn | `string` | required, replaces on change | aws.assessmenttarget.Arn |  |
| `AssessmentTemplateName` | assessment_template_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DurationInSeconds` | duration_in_seconds | `integer` | required, replaces on change |  |  |
| `RulesPackageArns` | rules_package_arns | `list` | required, replaces on change |  |  |
| `UserAttributesForFindings` | user_attributes_for_findings | `list` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: no

Discovery: supported
