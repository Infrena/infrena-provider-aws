# aws.iam.group

**CloudFormation type:** `AWS::IAM::Group`

Creates a new group.

Global type (no region attribute)

**Import ID:** `global/GroupName` (AWS::IAM::Group)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `GroupName` | group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the group to create. Do not include the path in this value. |
| `ManagedPolicyArns` | managed_policy_arns | `list` | optional, computed, provider-chosen | aws.managedpolicy.PolicyArn | The Amazon Resource Name (ARN) of the IAM policy you want to attach. |
| `Path` |  | `string` | optional, computed, provider-chosen |  | The path to the group. For more information about paths, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*. |
| `Policies` |  | `list` | optional, computed, provider-chosen |  | Adds or updates an inline policy document that is embedded in the specified IAM group. To view AWS::IAM::Group snippets, see [Declaring an Group Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-iam-group). |

Supports update: yes

Discovery: supported
