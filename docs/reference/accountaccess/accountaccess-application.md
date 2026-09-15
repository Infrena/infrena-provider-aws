# aws.accountaccess.application

**CloudFormation type:** `AWS::AccountAccess::Application`

Resource Type definition for AWS::AccountAccess::Application specifying an application for account access

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn` (AWS::AccountAccess::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | computed |  | The ARN of the application |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the application was created |
| `IdentitySource` | identity_source | `map` | required, replaces on change |  | The identity source for the application |
| `Status` |  | `string` | computed |  | The status of the application |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `TenantId` | tenant_id | `string` | computed |  | The tenant ID of the application |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the application was last updated |

Supports update: yes

Discovery: supported
