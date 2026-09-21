# aws.responseplan

**CloudFormation type:** `AWS::SSMIncidents::ResponsePlan`

Resource type definition for AWS::SSMIncidents::ResponsePlan

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMIncidents::ResponsePlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | optional, computed, provider-chosen |  | The list of actions. |
| `Arn` |  | `string` | computed |  | The ARN of the response plan. |
| `ChatChannel` | chat_channel | `map` | optional, computed, provider-chosen |  | The chat channel configuration. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The display name of the response plan. |
| `Engagements` |  | `list` | optional, computed, provider-chosen |  | The list of engagements to use. |
| `IncidentTemplate` | incident_template | `map` | required |  | The incident template configuration. |
| `Integrations` |  | `list` | optional, computed, provider-chosen |  | The list of integrations. |
| `Name` |  | `string` | required, replaces on change |  | The name of the response plan. |
| `Tags` |  | `map` | tags map |  | The tags to apply to the response plan. |

Supports update: yes

Discovery: supported
