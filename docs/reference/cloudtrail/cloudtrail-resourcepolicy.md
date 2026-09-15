# aws.cloudtrail.resourcepolicy

**CloudFormation type:** `AWS::CloudTrail::ResourcePolicy`

Resource Type definition for AWS::CloudTrail::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::CloudTrail::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The ARN of the AWS CloudTrail resource to which the policy applies. |
| `ResourcePolicy` | resource_policy | `string` | required |  | A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. |

Supports update: yes

Discovery: not supported
