# aws.quicksight.template

**CloudFormation type:** `AWS::QuickSight::Template`

Definition of the AWS::QuickSight::Template Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|TemplateId` (AWS::QuickSight::Template)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the template.</p> |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>Time when this was created.</p> |
| `Definition` |  | `map` | optional, computed, provider-chosen, write-only |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>Time when this was last updated.</p> |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `SourceEntity` | source_entity | `map` | optional, computed, provider-chosen, write-only |  | <p>The source entity of the template.</p> |
| `Tags` |  | `map` | tags map |  |  |
| `TemplateId` | template_id | `string` | required, replaces on change | aws.quicksight.template.TemplateId |  |
| `ValidationStrategy` | validation_strategy | `map` | optional, computed, provider-chosen, write-only |  | <p>The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to <code>LENIENT</code>, validation is skipped for specific errors.</p> |
| `Version` |  | `map` | computed |  | <p>A version of a template.</p> |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported (parent resource required)
