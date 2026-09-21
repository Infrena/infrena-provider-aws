# aws.bedrock.resourcepolicy

**CloudFormation type:** `AWS::Bedrock::ResourcePolicy`

Definition of AWS::Bedrock::ResourcePolicy Resource Type

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::Bedrock::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | required |  | The IAM policy document defining access permissions for the guardrail and guardrail profile resources |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The ARN of the Bedrock Guardrail or Guardrail Profile resource |

Supports update: yes

Discovery: not supported
