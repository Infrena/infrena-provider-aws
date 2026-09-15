# aws.storageprofile

**CloudFormation type:** `AWS::Deadline::StorageProfile`

Resource Type definition for AWS::Deadline::StorageProfile

Region attribute: `region`

**Import ID:** `<region>/FarmId|StorageProfileId` (AWS::Deadline::StorageProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DisplayName` | display_name | `string` | required |  |  |
| `FarmId` | farm_id | `string` | required, replaces on change | aws.farm.FarmId |  |
| `FileSystemLocations` | file_system_locations | `list` | optional, computed, provider-chosen |  |  |
| `OsFamily` | os_family | `string` | required |  |  |
| `StorageProfileId` | storage_profile_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
