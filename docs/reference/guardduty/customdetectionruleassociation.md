# aws.customdetectionruleassociation

**CloudFormation type:** `AWS::GuardDuty::CustomDetectionRuleAssociation`

Resource Type definition for AWS::GuardDuty::CustomDetectionRuleAssociation. Associates a GuardDuty custom detection rule with the caller's account, enabling the rule in either LIVE or DRY_RUN mode.

Region attribute: `region`

**Import ID:** `<region>/RuleId|AssociationId` (AWS::GuardDuty::CustomDetectionRuleAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The AWS account ID the association applies to. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the association. |
| `AssociationId` | association_id | `string` | computed |  | The service-generated unique identifier of the association. |
| `CreatedAt` | created_at | `string` | computed |  | The time the association was created. |
| `Mode` |  | `string` | required |  | Whether the rule runs in LIVE mode (generates findings) or DRY_RUN mode (evaluates without generating findings). |
| `RuleId` | rule_id | `string` | required, replaces on change |  | The catalog identifier of the custom detection rule to associate. |
| `Tags` |  | `map` | tags map |  | The tags applied to the association. |
| `UpdatedAt` | updated_at | `string` | computed |  | The time the association was last updated. |

Supports update: yes

Discovery: supported
