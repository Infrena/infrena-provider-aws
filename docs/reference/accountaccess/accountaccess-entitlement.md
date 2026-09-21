# aws.accountaccess.entitlement

**CloudFormation type:** `AWS::AccountAccess::Entitlement`

Resource Type definition for AWS::AccountAccess::Entitlement specifying an entitlement for account access

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn|EntitlementId` (AWS::AccountAccess::Entitlement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | required, replaces on change | aws.accountaccess.application.ApplicationArn | The ARN of the application |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the entitlement was created |
| `Entitlement` |  | `map` | required, replaces on change |  | The entitlement details |
| `EntitlementId` | entitlement_id | `string` | computed |  | The ID of the entitlement |

Supports update: no

Discovery: supported (parent resource required)
