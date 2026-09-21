# aws.softwarepackageversion

**CloudFormation type:** `AWS::IoT::SoftwarePackageVersion`

resource definition

Region attribute: `region`

**Import ID:** `<region>/PackageName|VersionName` (AWS::IoT::SoftwarePackageVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Artifact` |  | `map` | optional, computed, provider-chosen |  | The artifact location of the package version |
| `Attributes` |  | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ErrorReason` | error_reason | `string` | computed |  |  |
| `PackageName` | package_name | `string` | required, replaces on change |  |  |
| `PackageVersionArn` | package_version_arn | `string` | computed |  |  |
| `Recipe` |  | `string` | optional, computed, provider-chosen |  | The inline json job document associated with a software package version |
| `Sbom` |  | `map` | optional, computed, provider-chosen |  | The sbom zip archive location of the package version |
| `SbomValidationStatus` | sbom_validation_status | `string` | computed |  | The validation status of the Sbom file |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `VersionName` | version_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
