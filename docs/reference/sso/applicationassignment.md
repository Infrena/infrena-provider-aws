# aws.applicationassignment

**CloudFormation type:** `AWS::SSO::ApplicationAssignment`

Resource Type definition for SSO application access grant to a user or group.

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn|PrincipalType|PrincipalId` (AWS::SSO::ApplicationAssignment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | required, replaces on change | aws.sso.application.ApplicationArn | The ARN of the application. |
| `PrincipalId` | principal_id | `string` | required, replaces on change |  | An identifier for an object in IAM Identity Center, such as a user or group |
| `PrincipalType` | principal_type | `string` | required, replaces on change |  | The entity type for which the assignment will be created. |

Supports update: no

Discovery: supported
