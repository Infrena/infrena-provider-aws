# aws.case

**CloudFormation type:** `AWS::Cases::Case`

Creates a case in the specified Cases domain.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Cases::Case)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the case. |
| `CaseId` | case_id | `string` | computed |  | A unique identifier of the case. |
| `CustomerId` | customer_id | `string` | required, replaces on change |  | The full customer profile ARN for the case. |
| `DomainId` | domain_id | `string` | required, replaces on change | aws.cases.domain.DomainId | The unique identifier of the Cases domain. |
| `Tags` |  | `map` | tags map |  | A list of tags for the case. |
| `TemplateId` | template_id | `string` | required, replaces on change | aws.cases.template.TemplateId | A unique identifier of a template. |
| `Title` |  | `string` | required |  | The title of the case. |

Supports update: yes

Discovery: supported (parent resource required)
