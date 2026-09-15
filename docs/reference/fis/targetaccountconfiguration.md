# aws.targetaccountconfiguration

**CloudFormation type:** `AWS::FIS::TargetAccountConfiguration`

Resource schema for AWS::FIS::TargetAccountConfiguration

Region attribute: `region`

**Import ID:** `<region>/ExperimentTemplateId|AccountId` (AWS::FIS::TargetAccountConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | required, replaces on change |  | The AWS account ID of the target account. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the target account. |
| `ExperimentTemplateId` | experiment_template_id | `string` | required, replaces on change | aws.experimenttemplate.Id | The ID of the experiment template. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM role for the target account. |

Supports update: yes

Discovery: supported (parent resource required)
