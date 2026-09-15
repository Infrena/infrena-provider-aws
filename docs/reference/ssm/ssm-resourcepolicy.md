# aws.ssm.resourcepolicy

**CloudFormation type:** `AWS::SSM::ResourcePolicy`

Resource Type definition for AWS::SSM::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/PolicyId|ResourceArn` (AWS::SSM::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Policy` |  | `string` | required |  | Actual policy statement. |
| `PolicyHash` | policy_hash | `string` | computed |  | A snapshot identifier for the policy over time. |
| `PolicyId` | policy_id | `string` | computed |  | An unique identifier within the policies of a resource. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | Arn of OpsItemGroup etc. |

Supports update: yes

Discovery: supported
