# aws.ram.permission

**CloudFormation type:** `AWS::RAM::Permission`

Resource type definition for AWS::RAM::Permission

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RAM::Permission)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `IsResourceTypeDefault` | is_resource_type_default | `boolean` | computed |  | Set to true to use this as the default permission. |
| `Name` |  | `string` | required, replaces on change |  | The name of the permission. |
| `PermissionType` | permission_type | `string` | computed |  |  |
| `PolicyTemplate` | policy_template | `map` | required, replaces on change |  | Policy template for the permission. |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | The resource type this permission can be used with. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Version` |  | `string` | computed |  | Version of the permission. |

Supports update: yes

Discovery: supported
