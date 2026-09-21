# aws.smsvoice.resourcepolicy

**CloudFormation type:** `AWS::SMSVOICE::ResourcePolicy`

Resource Type definition for AWS::SMSVOICE::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::SMSVOICE::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | required |  | The JSON formatted resource-based policy to attach. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the AWS End User Messaging SMS and Voice resource to attach the resource-based policy to. |

Supports update: yes

Discovery: supported
