# aws.organizations.resourcepolicy

**CloudFormation type:** `AWS::Organizations::ResourcePolicy`

You can use AWS::Organizations::ResourcePolicy to delegate policy management for AWS Organizations to specified member accounts to perform policy actions that are by default available only to the management account.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Organizations::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the resource policy. |
| `Content` |  | `string` | required |  | The policy document. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it. |
| `Id` |  | `string` | computed |  | The unique identifier (ID) associated with this resource policy. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags that you want to attach to the resource policy |

Supports update: yes

Discovery: supported
