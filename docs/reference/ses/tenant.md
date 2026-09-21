# aws.tenant

**CloudFormation type:** `AWS::SES::Tenant`

Resource Type definition for AWS::SES::Tenant

Region attribute: `region`

**Import ID:** `<region>/TenantName` (AWS::SES::Tenant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) of the tenant. |
| `ResourceAssociations` | resource_associations | `list` | optional, computed, provider-chosen |  | The list of resources to associate with the tenant. |
| `Tags` |  | `map` | tags map |  | The tags (keys and values) associated with the tenant. |
| `TenantName` | tenant_name | `string` | required, replaces on change |  | The name of the tenant. |

Supports update: yes

Discovery: supported
