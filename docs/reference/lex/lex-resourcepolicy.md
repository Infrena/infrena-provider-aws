# aws.lex.resourcepolicy

**CloudFormation type:** `AWS::Lex::ResourcePolicy`

Resource Type definition for a resource policy with specified policy statements that attaches to a Lex bot or bot alias.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Lex::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | The Physical ID of the resource policy. |
| `Policy` |  | `map` | required |  | A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to. |
| `RevisionId` | revision_id | `string` | computed |  | The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource. |

Supports update: yes

Discovery: supported (parent resource required)
