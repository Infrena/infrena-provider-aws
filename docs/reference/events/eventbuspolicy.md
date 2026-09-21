# aws.eventbuspolicy

**CloudFormation type:** `AWS::Events::EventBusPolicy`

Resource Type definition for AWS::Events::EventBusPolicy

Region attribute: `region`

**Import ID:** `<region>/EventBusName|StatementId` (AWS::Events::EventBusPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | optional, computed, provider-chosen, write-only |  | The action that you are enabling the other account to perform. |
| `Condition` |  | `map` | optional, computed, provider-chosen, write-only |  | This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization. |
| `EventBusName` | event_bus_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the event bus associated with the rule. If you omit this, the default event bus is used. |
| `Principal` |  | `string` | optional, computed, provider-chosen, write-only |  | The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify "*" to permit any account to put events to your default event bus. |
| `Statement` |  | `map` | optional, computed, provider-chosen |  | A JSON string that describes the permission policy statement. You can include a Policy parameter in the request instead of using the StatementId, Action, Principal, or Condition parameters. |
| `StatementId` | statement_id | `string` | required, replaces on change |  | An identifier string for the external account that you are granting permissions to |

Supports update: yes

Discovery: supported
