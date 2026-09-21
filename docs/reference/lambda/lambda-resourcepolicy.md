# aws.lambda.resourcepolicy

**CloudFormation type:** `AWS::Lambda::ResourcePolicy`

Use the ``AWS::Lambda::ResourcePolicy`` resource to attach a resource-based policy to a LAM resource. A resource-based policy applies to a single LAM resource, for example, a function, function version, or function alias. To learn more about using resource-based policies with LAM, see [Working with resource-based policies in](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) in the *Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::Lambda::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | required |  | The policy document you want to add to your LAM resource. This is formatted as a JSON string. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the LAM resource you want to add the policy to. For a function, you can use a qualified or an unqualified ARN. The value must be a complete ARN, and the operation does not accept wildcard characters. |

Supports update: yes

Discovery: not supported
