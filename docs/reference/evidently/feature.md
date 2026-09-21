# aws.feature

**CloudFormation type:** `AWS::Evidently::Feature`

Resource Type definition for AWS::Evidently::Feature.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Evidently::Feature)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DefaultVariation` | default_variation | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EntityOverrides` | entity_overrides | `list` | optional, computed, provider-chosen |  |  |
| `EvaluationStrategy` | evaluation_strategy | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Project` |  | `string` | required, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to this resource. |
| `Variations` |  | `list` | required |  |  |

Supports update: yes

Discovery: not supported
