# aws.caserule

**CloudFormation type:** `AWS::Cases::CaseRule`

A case rule. In the Amazon Connect admin website, case rules are known as case field conditions. Case rules are used to define the situations under which fields should have certain effects (such as required).

Region attribute: `region`

**Import ID:** `<region>/CaseRuleArn` (AWS::Cases::CaseRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CaseRuleArn` | case_rule_arn | `string` | computed |  | The Amazon Resource Name (ARN) of a case rule. |
| `CaseRuleId` | case_rule_id | `string` | computed |  | The unique identifier of a case rule. |
| `CreatedTime` | created_time | `string` | computed |  | The time at which the case rule was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description explaining the purpose and behavior of this case rule. Helps administrators understand when and why this rule applies to case fields. |
| `DomainId` | domain_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cases.domain.DomainId | The unique identifier of the Cases domain. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the case rule was created or last modified. |
| `Name` |  | `string` | required |  | A descriptive name for the case rule. Must be unique within the domain and should clearly indicate the rule's purpose (e.g., 'Priority Field Required for Urgent Cases'). |
| `Rule` |  | `string` | required |  | Defines the rule behavior and conditions. Specifies the rule type and the conditions under which it applies. In the Amazon Connect admin website, this corresponds to case field conditions. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags that you attach to this case rule. |

Supports update: yes

Discovery: supported (parent resource required)
