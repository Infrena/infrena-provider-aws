# aws.role

**CloudFormation type:** `AWS::IAM::Role`

Creates a new role for your AWS-account.

Global type (no region attribute)

**Import ID:** `global/RoleName` (AWS::IAM::Role)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssumeRolePolicyDocument` | assume_role_policy, assume_role_policy_document | `string` | required |  | The trust policy that is associated with this role. Trust policies define which entities can assume the role. You can associate only one trust policy with a role. For an example of a policy that can be used to assume a role, see [Template Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#aws-resource-iam-role--examples). For more information about the elements that you can use in an IAM policy, see [Policy Elements Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in the *User Guide*. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the role that you provide. |
| `ManagedPolicyArns` | managed_policy_arns | `list` | optional, computed, provider-chosen | aws.managedpolicy.PolicyArn | A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. |
| `MaxSessionDuration` | max_session_duration | `integer` | optional, computed, provider-chosen |  | The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default value of one hour is applied. This setting can have a value from 1 hour to 12 hours. |
| `Path` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The path to the role. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*. |
| `PermissionsBoundary` | permissions_boundary | `string` | optional, computed, provider-chosen |  | The ARN of the policy used to set the permissions boundary for the role. |
| `Policies` |  | `list` | optional, computed, provider-chosen |  | Adds or updates an inline policy document that is embedded in the specified IAM role. |
| `RoleId` | role_id | `string` | computed |  |  |
| `RoleName` | name, role_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the IAM role, up to 64 characters in length. For valid values, see the ``RoleName`` parameter for the [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the *User Guide*. |
| `Tags` |  | `map` | tags map |  | A list of tags that are attached to the role. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the *IAM User Guide*. |

Supports update: yes

Discovery: supported
