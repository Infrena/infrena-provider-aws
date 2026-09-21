# aws.reviewtemplate

**CloudFormation type:** `AWS::WellArchitected::ReviewTemplate`

Creates a review template for the Well-Architected Tool.

Region attribute: `region`

**Import ID:** `<region>/TemplateArn` (AWS::WellArchitected::ReviewTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | required |  | The review template description. |
| `Lenses` |  | `list` | required, replaces on change |  | The lenses applied to the review template. |
| `Notes` |  | `string` | optional, computed, provider-chosen |  | The notes associated with the review template. |
| `Owner` |  | `string` | computed |  | The owner of the review template. |
| `Tags` |  | `map` | tags map |  | The tags assigned to the review template. |
| `TemplateArn` | template_arn | `string` | computed |  | The review template ARN. |
| `TemplateName` | template_name | `string` | required |  | The name of the review template. |
| `UpdateStatus` | update_status | `string` | computed |  | The latest status of the review template. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the review template was last updated. |

Supports update: yes

Discovery: supported
