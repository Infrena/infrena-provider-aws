# aws.softwarepackage

**CloudFormation type:** `AWS::IoT::SoftwarePackage`

resource definition

Region attribute: `region`

**Import ID:** `<region>/PackageName` (AWS::IoT::SoftwarePackage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `PackageArn` | package_arn | `string` | computed |  |  |
| `PackageName` | package_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
