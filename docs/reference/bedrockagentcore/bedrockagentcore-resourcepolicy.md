# aws.bedrockagentcore.resourcepolicy

**CloudFormation type:** `AWS::BedrockAgentCore::ResourcePolicy`

Resource Type definition for AWS::BedrockAgentCore::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::BedrockAgentCore::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Policy` |  | `string` | required |  | The resource policy to create or update. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the resource for which to create or update the resource policy. |

Supports update: yes

Discovery: supported (parent resource required)
