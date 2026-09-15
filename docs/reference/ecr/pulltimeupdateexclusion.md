# aws.pulltimeupdateexclusion

**CloudFormation type:** `AWS::ECR::PullTimeUpdateExclusion`

The ARN of the IAM principal to remove from the pull time update exclusion list.

Region attribute: `region`

**Import ID:** `<region>/PrincipalArn` (AWS::ECR::PullTimeUpdateExclusion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PrincipalArn` | principal_arn | `string` | required, replaces on change |  | Principal arn that should not update image pull times. |

Supports update: no

Discovery: supported
