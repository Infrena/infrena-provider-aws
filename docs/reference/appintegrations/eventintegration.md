# aws.eventintegration

**CloudFormation type:** `AWS::AppIntegrations::EventIntegration`

Resource Type definition for AWS::AppIntegrations::EventIntegration

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::AppIntegrations::EventIntegration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The event integration description. |
| `EventBridgeBus` | event_bridge_bus | `string` | required, replaces on change |  | The Amazon Eventbridge bus for the event integration. |
| `EventFilter` | event_filter | `map` | required, replaces on change |  | The EventFilter (source) associated with the event integration. |
| `EventIntegrationArn` | event_integration_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the event integration. |
| `Name` |  | `string` | required, replaces on change |  | The name of the event integration. |
| `Tags` |  | `map` | tags map |  | The tags (keys and values) associated with the event integration. |

Supports update: yes

Discovery: supported
