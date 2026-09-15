# aws.computeoptimizer.automationrule

**CloudFormation type:** `AWS::ComputeOptimizer::AutomationRule`

Creates an AWS Compute Optimizer automation rule that automatically implements recommended actions based on your defined criteria and schedule. Automation rules are global resources that manage automated actions across all AWS Regions where Compute Optimizer Automation is available. Organization-level rules can only be created by the management account or delegated administrator.

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::ComputeOptimizer::AutomationRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The AWS account ID that owns the automation rule. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The timestamp when the automation rule was created. |
| `Criteria` |  | `map` | optional, computed, provider-chosen |  | Filter criteria that specify which recommended actions qualify for implementation. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the automation rule. |
| `LastUpdatedTimestamp` | last_updated_timestamp | `string` | computed |  | The timestamp when the automation rule was last updated. |
| `Name` |  | `string` | required |  | The name of the automation rule. |
| `OrganizationConfiguration` | organization_configuration | `map` | optional, computed, provider-chosen |  | Organization configuration for organization rules, including rule apply order and account scope. |
| `Priority` |  | `string` | optional, computed, provider-chosen |  | Rule priority within its group |
| `RecommendedActionTypes` | recommended_action_types | `list` | required |  | The types of recommended actions this rule will implement. |
| `RuleArn` | rule_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the automation rule. |
| `RuleId` | rule_id | `string` | computed |  | The unique identifier of the automation rule. |
| `RuleRevision` | rule_revision | `string` | computed |  | The revision number of the automation rule. |
| `RuleType` | rule_type | `string` | required |  | The type of automation rule. |
| `Schedule` |  | `map` | required |  | The schedule configuration for when the rule runs. |
| `Status` |  | `string` | required |  | The status of the automation rule. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with the automation rule. |

Supports update: yes

Discovery: supported
