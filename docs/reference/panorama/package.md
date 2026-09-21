# aws.package

**CloudFormation type:** `AWS::Panorama::Package`

Creates a package and storage location in an Amazon S3 access point.

Region attribute: `region`

**Import ID:** `<region>/PackageId` (AWS::Panorama::Package)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `integer` | computed |  |  |
| `PackageId` | package_id | `string` | computed |  |  |
| `PackageName` | package_name | `string` | required, replaces on change |  | A name for the package. |
| `StorageLocation` | storage_location | `map` | optional, computed, provider-chosen |  | A storage location. |
| `Tags` |  | `map` | tags map |  | Tags for the package. |

Supports update: yes

Discovery: supported
