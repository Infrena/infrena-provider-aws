# aws.custompermissions

**CloudFormation type:** `AWS::QuickSight::CustomPermissions`

Definition of the AWS::QuickSight::CustomPermissions Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|CustomPermissionsName` (AWS::QuickSight::CustomPermissions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `Capabilities` |  | `map` | optional, computed, provider-chosen |  |  |
| `CustomPermissionsName` | custom_permissions_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
