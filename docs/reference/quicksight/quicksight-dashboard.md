# aws.quicksight.dashboard

**CloudFormation type:** `AWS::QuickSight::Dashboard`

Definition of the AWS::QuickSight::Dashboard Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|DashboardId` (AWS::QuickSight::Dashboard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the resource.</p> |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that this dashboard was created.</p> |
| `DashboardId` | dashboard_id | `string` | required, replaces on change | aws.quicksight.dashboard.DashboardId |  |
| `DashboardPublishOptions` | dashboard_publish_options | `map` | optional, computed, provider-chosen, write-only |  | <p>Dashboard publish options.</p> |
| `Definition` |  | `map` | optional, computed, provider-chosen, write-only |  |  |
| `FolderArns` | folder_arns | `list` | optional, computed, provider-chosen, write-only | aws.folder.Arn |  |
| `LastPublishedTime` | last_published_time | `string` | computed |  | <p>The last time that this dashboard was published.</p> |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The last time that this dashboard was updated.</p> |
| `LinkEntities` | link_entities | `list` | optional, computed, provider-chosen |  |  |
| `LinkSharingConfiguration` | link_sharing_configuration | `map` | optional, computed, provider-chosen, write-only |  |  |
| `Name` |  | `string` | required |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen, write-only |  | <p>A list of Amazon QuickSight parameters and the list's override values.</p> |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `SourceEntity` | source_entity | `map` | optional, computed, provider-chosen, write-only |  | <p>Dashboard source entity.</p> |
| `Tags` |  | `map` | tags map |  |  |
| `ThemeArn` | theme_arn | `string` | optional, computed, provider-chosen, write-only | aws.quicksight.theme.Arn |  |
| `ValidationStrategy` | validation_strategy | `map` | optional, computed, provider-chosen, write-only |  | <p>The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to <code>LENIENT</code>, validation is skipped for specific errors.</p> |
| `Version` |  | `map` | computed |  | <p>Dashboard version.</p> |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported (parent resource required)
