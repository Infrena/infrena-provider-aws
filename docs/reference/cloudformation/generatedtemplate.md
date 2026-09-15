# aws.generatedtemplate

**CloudFormation type:** `AWS::CloudFormation::GeneratedTemplate`

Creates a generated template from existing resources using the CloudFormation IaC Generator.

Region attribute: `region`

**Import ID:** `<region>/GeneratedTemplateId` (AWS::CloudFormation::GeneratedTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time the generated template was created. |
| `GeneratedTemplateId` | generated_template_id | `string` | computed |  | The Amazon Resource Name (ARN) of the generated template. |
| `GeneratedTemplateName` | generated_template_name | `string` | required |  | The name assigned to the generated template. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The time the generated template was last updated. |
| `Progress` |  | `map` | computed |  | A summary of the progress of the template generation. |
| `Status` |  | `string` | computed |  | The status of the template generation. |
| `TemplateConfiguration` | template_configuration | `map` | optional, computed, provider-chosen |  | The configuration details of the generated template. |
| `TotalWarnings` | total_warnings | `integer` | computed |  | The number of warnings generated for this template. |

Supports update: yes

Discovery: supported
