# aws.cases.template

**CloudFormation type:** `AWS::Cases::Template`

A template in the Cases domain. This template is used to define the case object model (that is, to define what data can be captured on cases) in a Cases domain. A template must have a unique name within a domain, and it must reference existing field IDs and layout IDs.

Region attribute: `region`

**Import ID:** `<region>/TemplateArn` (AWS::Cases::Template)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTime` | created_time | `string` | computed |  | The time at which the template was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description explaining the purpose and use case for this template. Should indicate what types of cases this template is designed for and any specific workflow it supports. |
| `DomainId` | domain_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cases.domain.DomainId | The unique identifier of the Cases domain. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the template was created or last modified. |
| `LayoutConfiguration` | layout_configuration | `map` | optional, computed, provider-chosen |  | Specifies the default layout to use when displaying cases created from this template. The layout determines which fields are visible and their arrangement in the agent interface. |
| `Name` |  | `string` | required |  | A name for the template. It must be unique per domain. |
| `RequiredFields` | required_fields | `list` | optional, computed, provider-chosen |  | A list of fields that must contain a value for a case to be successfully created with this template. |
| `Rules` |  | `list` | optional, computed, provider-chosen |  | A list of case rules (also known as case field conditions) on a template. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The current status of the template. Active templates can be used to create new cases, while Inactive templates are disabled but preserved for existing cases. |
| `Tags` |  | `map` | tags map |  | The tags that you attach to this template. |
| `TemplateArn` | template_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the template. |
| `TemplateId` | template_id | `string` | computed |  | The unique identifier of a template. |

Supports update: yes

Discovery: supported (parent resource required)
