# aws.grant

**CloudFormation type:** `AWS::LicenseManager::Grant`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/GrantArn` (AWS::LicenseManager::Grant)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedOperations` | allowed_operations | `list` | optional, computed, provider-chosen, write-only |  |  |
| `GrantArn` | grant_arn | `string` | computed |  | Arn of the grant. |
| `GrantName` | grant_name | `string` | optional, computed, provider-chosen |  | Name for the created Grant. |
| `HomeRegion` | home_region | `string` | optional, computed, provider-chosen |  | Home region for the created grant. |
| `LicenseArn` | license_arn | `string` | optional, computed, provider-chosen | aws.license.LicenseArn | License Arn for the grant. |
| `Principals` |  | `list` | optional, computed, provider-chosen, write-only |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Tags` |  | `map` | tags map |  | A list of tags to attach. |
| `Version` |  | `string` | computed |  | The version of the grant. |

Supports update: yes

Discovery: supported
