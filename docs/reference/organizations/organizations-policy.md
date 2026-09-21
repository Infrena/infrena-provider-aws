# aws.organizations.policy

**CloudFormation type:** `AWS::Organizations::Policy`

Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Organizations::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the Policy |
| `AwsManaged` | aws_managed | `boolean` | computed |  | A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it. |
| `Content` |  | `string` | required |  | The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Human readable description of the policy |
| `Id` |  | `string` | computed |  | Id of the Policy |
| `Name` |  | `string` | required |  | Name of the Policy |
| `Tags` |  | `map` | tags map |  | A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null. |
| `TargetIds` | target_ids | `list` | optional, computed, provider-chosen |  | List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to |
| `Type` | type_value | `string` | required, replaces on change |  | The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, BEDROCK_POLICY, CHATBOT_POLICY, DECLARATIVE_POLICY_EC2, INSPECTOR_POLICY, NETWORK_SECURITY_DIRECTOR_POLICY, RESOURCE_CONTROL_POLICY, S3_POLICY, SECURITYHUB_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY, UPGRADE_ROLLOUT_POLICY |

Supports update: yes

Discovery: supported (parent resource required)
