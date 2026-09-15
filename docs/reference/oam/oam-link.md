# aws.oam.link

**CloudFormation type:** `AWS::Oam::Link`

Definition of AWS::Oam::Link Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Oam::Link)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Label` |  | `string` | computed |  |  |
| `LabelTemplate` | label_template | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `LinkConfiguration` | link_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ResourceTypes` | resource_types | `list` | required |  |  |
| `SinkIdentifier` | sink_identifier | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | Tags to apply to the link |

Supports update: yes

Discovery: supported
