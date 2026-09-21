# aws.iam.instanceprofile

**CloudFormation type:** `AWS::IAM::InstanceProfile`

Creates a new instance profile. For information about instance profiles, see [Using instance profiles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2_instance-profiles.html).

Global type (no region attribute)

**Import ID:** `global/InstanceProfileName` (AWS::IAM::InstanceProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `InstanceProfileName` | instance_profile_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the instance profile to create. |
| `Path` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*. |
| `Roles` |  | `list` | required |  | The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions. |

Supports update: yes

Discovery: supported
