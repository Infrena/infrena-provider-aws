# aws.quicksight.theme

**CloudFormation type:** `AWS::QuickSight::Theme`

Definition of the AWS::QuickSight::Theme Resource Type.

Region attribute: `region`

**Import ID:** `<region>/ThemeId|AwsAccountId` (AWS::QuickSight::Theme)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the theme.</p> |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `BaseThemeId` | base_theme_id | `string` | required | aws.quicksight.theme.ThemeId |  |
| `Configuration` |  | `map` | required |  | <p>The theme configuration. This configuration contains all of the display properties for |
| `CreatedTime` | created_time | `string` | computed |  | <p>The date and time that the theme was created.</p> |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The date and time that the theme was last updated.</p> |
| `Name` |  | `string` | required |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `ThemeId` | theme_id | `string` | required, replaces on change | aws.quicksight.theme.ThemeId |  |
| `Type` | type_value | `string` | computed |  |  |
| `Version` |  | `map` | computed |  | <p>A version of a theme.</p> |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
