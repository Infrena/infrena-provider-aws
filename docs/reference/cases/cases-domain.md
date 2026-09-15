# aws.cases.domain

**CloudFormation type:** `AWS::Cases::Domain`

A domain, which is a container for all case data, such as cases, fields, templates and layouts. Each Amazon Connect instance can be associated with only one Cases domain.

Region attribute: `region`

**Import ID:** `<region>/DomainArn` (AWS::Cases::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTime` | created_time | `string` | computed |  | The time at which the domain was created. |
| `DomainArn` | domain_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the Cases domain. |
| `DomainId` | domain_id | `string` | computed |  | The unique identifier of the Cases domain. |
| `DomainStatus` | domain_status | `string` | computed |  | The current status of the Cases domain. Indicates whether the domain is Active, CreationInProgress, or CreationFailed. |
| `Name` |  | `string` | required, replaces on change |  | The name for your Cases domain. It must be unique for your AWS account. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags that you attach to this domain. |

Supports update: yes

Discovery: supported
