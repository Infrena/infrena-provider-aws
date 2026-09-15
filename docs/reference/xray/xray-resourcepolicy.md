# aws.xray.resourcepolicy

**CloudFormation type:** `AWS::XRay::ResourcePolicy`

This schema provides construct and validation rules for AWS-XRay Resource Policy resource parameters.

Region attribute: `region`

**Import ID:** `<region>/PolicyName` (AWS::XRay::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BypassPolicyLockoutCheck` | bypass_policy_lockout_check | `boolean` | optional, computed, provider-chosen, write-only |  | A flag to indicate whether to bypass the resource policy lockout safety check |
| `PolicyDocument` | policy_document | `string` | required |  | The resource policy document, which can be up to 5kb in size. |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | The name of the resource policy. Must be unique within a specific AWS account. |

Supports update: yes

Discovery: supported
