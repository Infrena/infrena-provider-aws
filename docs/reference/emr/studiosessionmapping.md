# aws.studiosessionmapping

**CloudFormation type:** `AWS::EMR::StudioSessionMapping`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/StudioId|IdentityType|IdentityName` (AWS::EMR::StudioSessionMapping)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IdentityName` | identity_name | `string` | required, replaces on change |  | The name of the user or group. For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified. |
| `IdentityType` | identity_type | `string` | required, replaces on change |  | Specifies whether the identity to map to the Studio is a user or a group. |
| `SessionPolicyArn` | session_policy_arn | `string` | required |  | The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group. Session policies refine Studio user permissions without the need to use multiple IAM user roles. |
| `StudioId` | studio_id | `string` | required, replaces on change | aws.studio.StudioId | The ID of the Amazon EMR Studio to which the user or group will be mapped. |

Supports update: yes

Discovery: supported
