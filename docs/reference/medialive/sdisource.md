# aws.sdisource

**CloudFormation type:** `AWS::MediaLive::SdiSource`

Definition of AWS::MediaLive::SdiSource Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaLive::SdiSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The unique arn of the SdiSource. |
| `Id` |  | `string` | computed |  | The unique identifier of the SdiSource. |
| `Inputs` |  | `list` | computed |  | The list of inputs currently using this SDI source. |
| `Mode` |  | `string` | optional, computed, provider-chosen |  | The current state of the SdiSource. |
| `Name` |  | `string` | required |  | The name of the SdiSource. |
| `State` |  | `string` | computed |  | The current state of the SdiSource. |
| `Tags` |  | `map` | tags map |  | A collection of key-value pairs. |
| `Type` | type_value | `string` | required |  | The interface mode of the SdiSource. |

Supports update: yes

Discovery: supported
