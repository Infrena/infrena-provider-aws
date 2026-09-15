# aws.workspacespool

**CloudFormation type:** `AWS::WorkSpaces::WorkspacesPool`

Resource Type definition for AWS::WorkSpaces::WorkspacesPool

Region attribute: `region`

**Import ID:** `<region>/PoolId` (AWS::WorkSpaces::WorkspacesPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationSettings` | application_settings | `map` | optional, computed, provider-chosen |  |  |
| `BundleId` | bundle_id | `string` | required |  |  |
| `Capacity` |  | `map` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DirectoryId` | directory_id | `string` | required |  |  |
| `PoolArn` | pool_arn | `string` | computed |  |  |
| `PoolId` | pool_id | `string` | computed |  |  |
| `PoolName` | pool_name | `string` | required, replaces on change |  |  |
| `RunningMode` | running_mode | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TimeoutSettings` | timeout_settings | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
