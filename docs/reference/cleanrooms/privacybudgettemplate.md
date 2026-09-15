# aws.privacybudgettemplate

**CloudFormation type:** `AWS::CleanRooms::PrivacyBudgetTemplate`

Represents a privacy budget within a collaboration

Region attribute: `region`

**Import ID:** `<region>/PrivacyBudgetTemplateIdentifier|MembershipIdentifier` (AWS::CleanRooms::PrivacyBudgetTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AutoRefresh` | auto_refresh | `string` | required, replaces on change |  |  |
| `CollaborationArn` | collaboration_arn | `string` | computed |  |  |
| `CollaborationIdentifier` | collaboration_identifier | `string` | computed |  |  |
| `MembershipArn` | membership_arn | `string` | computed |  |  |
| `MembershipIdentifier` | membership_identifier | `string` | required, replaces on change |  |  |
| `Parameters` |  | `map` | required |  |  |
| `PrivacyBudgetTemplateIdentifier` | privacy_budget_template_identifier | `string` | computed |  |  |
| `PrivacyBudgetType` | privacy_budget_type | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this cleanrooms privacy budget template. |

Supports update: yes

Discovery: supported (parent resource required)
