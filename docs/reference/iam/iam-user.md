# aws.iam.user

**CloudFormation type:** `AWS::IAM::User`

Creates a new IAM user for your AWS-account.

Global type (no region attribute)

**Import ID:** `global/UserName` (AWS::IAM::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Groups` |  | `list` | optional, computed, provider-chosen |  | A list of group names to which you want to add the user. |
| `LoginProfile` | login_profile | `map` | optional, computed, provider-chosen |  | Creates a password for the specified user, giving the user the ability to access AWS services through the console. For more information about managing passwords, see [Managing Passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *User Guide*. |
| `ManagedPolicyArns` | managed_policy_arns | `list` | optional, computed, provider-chosen | aws.managedpolicy.PolicyArn | A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the user. |
| `Path` |  | `string` | optional, computed, provider-chosen |  | The path for the user name. For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*. |
| `PermissionsBoundary` | permissions_boundary | `string` | optional, computed, provider-chosen |  | The ARN of the managed policy that is used to set the permissions boundary for the user. |
| `Policies` |  | `list` | optional, computed, provider-chosen |  | Adds or updates an inline policy document that is embedded in the specified IAM user. To view AWS::IAM::User snippets, see [Declaring an User Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-user). |
| `Tags` |  | `map` | tags map |  | A list of tags that you want to attach to the new user. Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*. |
| `UserName` | user_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the user to create. Do not include the path in this value. |

Supports update: yes

Discovery: supported
