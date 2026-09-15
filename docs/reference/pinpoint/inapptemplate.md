# aws.inapptemplate

**CloudFormation type:** `AWS::Pinpoint::InAppTemplate`

Resource Type definition for AWS::Pinpoint::InAppTemplate

Region attribute: `region`

**Import ID:** `<region>/TemplateName` (AWS::Pinpoint::InAppTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Content` |  | `list` | optional, computed, provider-chosen |  |  |
| `CustomConfig` | custom_config | `map` | optional, computed, provider-chosen |  |  |
| `Layout` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `TemplateDescription` | template_description | `string` | optional, computed, provider-chosen |  |  |
| `TemplateName` | template_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
