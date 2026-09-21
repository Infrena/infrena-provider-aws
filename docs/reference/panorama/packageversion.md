# aws.packageversion

**CloudFormation type:** `AWS::Panorama::PackageVersion`

Registers a package version.

Region attribute: `region`

**Import ID:** `<region>/PackageId|PackageVersion|PatchVersion` (AWS::Panorama::PackageVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IsLatestPatch` | is_latest_patch | `boolean` | computed |  |  |
| `MarkLatest` | mark_latest | `boolean` | optional, computed, provider-chosen |  | Whether to mark the new version as the latest version. |
| `OwnerAccount` | owner_account | `string` | optional, computed, provider-chosen, replaces on change |  | An owner account. |
| `PackageArn` | package_arn | `string` | computed |  |  |
| `PackageId` | package_id | `string` | required, replaces on change | aws.package.PackageId | A package ID. |
| `PackageName` | package_name | `string` | computed |  |  |
| `PackageVersion` | package_version | `string` | required, replaces on change |  | A package version. |
| `PatchVersion` | patch_version | `string` | required, replaces on change |  | A patch version. |
| `RegisteredTime` | registered_time | `integer` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusDescription` | status_description | `string` | computed |  |  |
| `UpdatedLatestPatchVersion` | updated_latest_patch_version | `string` | optional, computed, provider-chosen, write-only |  | If the version was marked latest, the new version to maker as latest. |

Supports update: yes

Discovery: not supported
