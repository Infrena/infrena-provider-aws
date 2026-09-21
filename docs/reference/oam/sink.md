# aws.sink

**CloudFormation type:** `AWS::Oam::Sink`

Resource Type definition for AWS::Oam::Sink

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Oam::Sink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon resource name (ARN) of the ObservabilityAccessManager Sink |
| `Name` |  | `string` | required, replaces on change |  | The name of the ObservabilityAccessManager Sink. |
| `Policy` |  | `map` | optional, computed, provider-chosen |  | The policy of this ObservabilityAccessManager Sink. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Tags to apply to the sink |

Supports update: yes

Discovery: supported
